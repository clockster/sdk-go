// The examples the documentation is written around. They are compiled with the tests, so a
// snippet that stops working stops the build rather than the reader.

package clockster_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/clockster/sdk-go"
)

func ExampleNew() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	me, err := client.Me(ctx)
	if err != nil {
		log.Fatal(err)
	}

	locations, err := client.Locations.Upsert(ctx, &clockster.LocationsUpsertBody{
		Items: []clockster.LocationsUpsertItem{{ExternalID: "HQ", Title: "Head office"}},
	})
	if err != nil {
		log.Fatal(err)
	}

	if _, err := client.Users.Upsert(ctx, &clockster.UsersUpsertBody{
		Users: []clockster.UsersUpsertUser{{
			ExternalID: clockster.Set("HR-1"),
			FirstName:  "Aisulu",
			Role:       "employee",
			LocationID: locations.Data[0].ID,
		}},
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println(me.Data.Title)
}

func ExampleNew_options() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"),
		// A demo stand answers the same API; production is the default.
		clockster.WithBaseURL("https://demo.clockster.com"),
		clockster.WithTimeout(60*time.Second),
		clockster.WithUserAgent("acme-hr/1.4"),
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = client
}

func ExampleUsers_ListAll() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	for user, err := range client.Users.ListAll(context.Background(), &clockster.UsersListParams{
		PerPage: clockster.Set(int64(100)),
		Include: []string{"location"},
	}) {
		if err != nil {
			log.Fatal(err)
		}

		// A relation is absent unless `include` asked for it, and null where it is empty; both are
		// a nil pointer here.
		if user.Location != nil {
			fmt.Println(user.ID, user.Location.Title)
		}
	}
}

func ExampleUsers_Upsert() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Users.Upsert(context.Background(), &clockster.UsersUpsertBody{
		Users: []clockster.UsersUpsertUser{{
			ExternalID: clockster.Set("HR-1"),
			FirstName:  "Aisulu",
			Role:       "employee",
			LocationID: 3,
			// Written as null, which clears the stored position.
			PositionID: clockster.Null[int64](),
			// DepartmentID is not set: not written at all, and the stored department stays.
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleError() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Users.Upsert(context.Background(), &clockster.UsersUpsertBody{
		Users: []clockster.UsersUpsertUser{{FirstName: "Aisulu"}},
	})

	var refused *clockster.Error

	if errors.As(err, &refused) {
		// Code names the reason and does not change; Message is prose and may.
		fmt.Println(refused.Code, refused.Errors, refused.RequestID)
	}

	if errors.Is(err, clockster.ErrRateLimit) {
		time.Sleep(time.Duration(refused.RetryAfter) * time.Second)
	}
}

func ExampleFiles_Upload() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("agreement.pdf")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	stored, err := client.Files.Upload(context.Background(), &clockster.FilesUploadForm{
		File:     file,
		Filename: "agreement.pdf",
		Name:     clockster.Set("agreement"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stored.Data.ID)
}

func ExampleWithIdempotencyKey() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// A rota has nothing of your own to match a second attempt against, so a retry is safe only
	// with a key: the same one answers the first result rather than writing the rota twice.
	_, err = client.Schedules.Create(context.Background(), &clockster.SchedulesCreateBody{
		Schedules: []clockster.SchedulesCreateSchedule{
			clockster.WorkSchedule{
				Type:     "work",
				Dates:    []string{"2026-08-17"},
				Users:    []int64{7},
				Timezone: "Asia/Almaty",
				Start:    clockster.Set("2026-08-17T09:00:00+05:00"),
				End:      clockster.Set("2026-08-17T18:00:00+05:00"),
			},
		},
	}, clockster.WithIdempotencyKey("2f8a1c-attempt-1"))
	if err != nil {
		log.Fatal(err)
	}
}
