package webhooks_test

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/clockster/sdk-go/webhooks"
)

func ExampleVerifyRequest() {
	http.HandleFunc("/clockster", func(w http.ResponseWriter, r *http.Request) {
		event, err := webhooks.VerifyRequest(r, os.Getenv("CLOCKSTER_WEBHOOK_SECRET"))
		if err != nil {
			http.Error(w, "refused", http.StatusBadRequest)

			return
		}

		// Answer quickly and do the work afterwards: a delivery that times out is sent again.
		w.WriteHeader(http.StatusOK)

		go handle(event)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(event *webhooks.Event) {
	// The same event may arrive twice, so deduplicate on the id — null on a trial delivery, which
	// stands for no recorded event.
	if event.ID == nil {
		return
	}

	var user struct {
		ID int64 `json:"id"`
	}

	if err := json.Unmarshal(event.Data, &user); err != nil {
		return
	}

	log.Println(event.Event, user.ID)
}
