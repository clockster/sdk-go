// Code generated from openapi/company-v3.json; DO NOT EDIT.

// The operations of the Company API, as they are called.
//
// Every method answers the parsed body of the response, and a refusal is an error rather than a
// status to read — so there is no branch between the call and the rows.

package clockster

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"
)

// Client is the Company API, one token to one company.
//
// Build one with [New]. Every operation hangs off a field below, named after the section of
// the documentation it belongs to: `clockster.Users.List(ctx, …)`. A Client is safe for
// concurrent use.
type Client struct {
	*transport

	// Attendance is `clockster.Attendance`.
	Attendance *Attendance

	// Departments is `clockster.Departments`.
	Departments *Departments

	// Documents is `clockster.Documents`.
	Documents *Documents

	// Files is `clockster.Files`.
	Files *Files

	// Locations is `clockster.Locations`.
	Locations *Locations

	// Payroll is `clockster.Payroll`.
	Payroll *Payroll

	// Positions is `clockster.Positions`.
	Positions *Positions

	// Schedules is `clockster.Schedules`.
	Schedules *Schedules

	// Tasks is `clockster.Tasks`.
	Tasks *Tasks

	// Timesheets is `clockster.Timesheets`.
	Timesheets *Timesheets

	// UserFilters is `clockster.UserFilters`.
	UserFilters *UserFilters

	// UserRequests is `clockster.UserRequests`.
	UserRequests *UserRequests

	// Users is `clockster.Users`.
	Users *Users

	// Webhooks is `clockster.Webhooks`.
	Webhooks *Webhooks
}

func newClient(t *transport) *Client {
	return &Client{
		transport:    t,
		Attendance:   newAttendance(t),
		Departments:  newDepartments(t),
		Documents:    newDocuments(t),
		Files:        newFiles(t),
		Locations:    newLocations(t),
		Payroll:      newPayroll(t),
		Positions:    newPositions(t),
		Schedules:    newSchedules(t),
		Tasks:        newTasks(t),
		Timesheets:   newTimesheets(t),
		UserFilters:  newUserFilters(t),
		UserRequests: newUserRequests(t),
		Users:        newUsers(t),
		Webhooks:     newWebhooks(t),
	}
}

// Attendance holds the operations of `clockster.Attendance`.
type Attendance struct {
	*transport
}

func newAttendance(t *transport) *Attendance {
	return &Attendance{
		transport: t,
	}
}

// Departments holds the operations of `clockster.Departments`.
type Departments struct {
	*transport
}

func newDepartments(t *transport) *Departments {
	return &Departments{
		transport: t,
	}
}

// Documents holds the operations of `clockster.Documents`.
type Documents struct {
	*transport
}

func newDocuments(t *transport) *Documents {
	return &Documents{
		transport: t,
	}
}

// Files holds the operations of `clockster.Files`.
type Files struct {
	*transport
}

func newFiles(t *transport) *Files {
	return &Files{
		transport: t,
	}
}

// Locations holds the operations of `clockster.Locations`.
type Locations struct {
	*transport
}

func newLocations(t *transport) *Locations {
	return &Locations{
		transport: t,
	}
}

// Payroll holds the operations of `clockster.Payroll`.
type Payroll struct {
	*transport

	// Payslips is `clockster.Payroll.Payslips`.
	Payslips *PayrollPayslips
}

func newPayroll(t *transport) *Payroll {
	return &Payroll{
		transport: t,
		Payslips:  newPayrollPayslips(t),
	}
}

// PayrollPayslips holds the operations of `clockster.Payroll.Payslips`.
type PayrollPayslips struct {
	*transport
}

func newPayrollPayslips(t *transport) *PayrollPayslips {
	return &PayrollPayslips{
		transport: t,
	}
}

// Positions holds the operations of `clockster.Positions`.
type Positions struct {
	*transport
}

func newPositions(t *transport) *Positions {
	return &Positions{
		transport: t,
	}
}

// Schedules holds the operations of `clockster.Schedules`.
type Schedules struct {
	*transport
}

func newSchedules(t *transport) *Schedules {
	return &Schedules{
		transport: t,
	}
}

// Tasks holds the operations of `clockster.Tasks`.
type Tasks struct {
	*transport
}

func newTasks(t *transport) *Tasks {
	return &Tasks{
		transport: t,
	}
}

// Timesheets holds the operations of `clockster.Timesheets`.
type Timesheets struct {
	*transport
}

func newTimesheets(t *transport) *Timesheets {
	return &Timesheets{
		transport: t,
	}
}

// UserFilters holds the operations of `clockster.UserFilters`.
type UserFilters struct {
	*transport
}

func newUserFilters(t *transport) *UserFilters {
	return &UserFilters{
		transport: t,
	}
}

// UserRequests holds the operations of `clockster.UserRequests`.
type UserRequests struct {
	*transport
}

func newUserRequests(t *transport) *UserRequests {
	return &UserRequests{
		transport: t,
	}
}

// Users holds the operations of `clockster.Users`.
type Users struct {
	*transport
}

func newUsers(t *transport) *Users {
	return &Users{
		transport: t,
	}
}

// Webhooks holds the operations of `clockster.Webhooks`.
type Webhooks struct {
	*transport

	// Deliveries is `clockster.Webhooks.Deliveries`.
	Deliveries *WebhooksDeliveries

	// Events is `clockster.Webhooks.Events`.
	Events *WebhooksEvents
}

func newWebhooks(t *transport) *Webhooks {
	return &Webhooks{
		transport:  t,
		Deliveries: newWebhooksDeliveries(t),
		Events:     newWebhooksEvents(t),
	}
}

// WebhooksDeliveries holds the operations of `clockster.Webhooks.Deliveries`.
type WebhooksDeliveries struct {
	*transport
}

func newWebhooksDeliveries(t *transport) *WebhooksDeliveries {
	return &WebhooksDeliveries{
		transport: t,
	}
}

// WebhooksEvents holds the operations of `clockster.Webhooks.Events`.
type WebhooksEvents struct {
	*transport
}

func newWebhooksEvents(t *transport) *WebhooksEvents {
	return &WebhooksEvents{
		transport: t,
	}
}

// Me is GET /company/v3/me.
//
// Whose token this is.
//
// Confirms which company a key belongs to.
//
// Answers the id and the name, and nothing else: everything a key opens is reachable from
// the endpoints themselves, and a company attribute this surface does not act on would
// only read as one it does.
func (c *Client) Me(ctx context.Context, opts ...RequestOption) (*MeResponse, error) {
	var out MeResponse

	if err := c.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/me",
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/attendance.
//
// List clock-ins.
//
// Marks inside a window of days, oldest first.
//
// **The window is required and capped at 90 days.** This is the largest table in the
// product; an unbounded read of it has no plan that finishes. `date_from` and `date_to`
// are plain dates matched against the clock the mark was stamped with, not against an
// instant — a day means the same thing here as it does to the person who worked it.
//
// **`datetime` is the moment as it was recorded**, offset included — the stored wall clock
// read in the stored zone, never shifted anywhere else. It is one field because it is one
// fact: the wall clock is the value without its offset, and the zone is the offset. On the
// rare row whose zone cannot be read it comes back with no offset at all, which is how you
// see that the zone was not known.
//
// `status` is `in`, `out` or `break`, not the integer the column keeps.
//
// Late arrivals are the one thing to plan for: a device that was offline uploads what it
// recorded earlier, so a mark can appear inside a window you have already read. Re-read
// the last few days with overlap rather than paging strictly forward and never looking
// back.
func (n *Attendance) List(ctx context.Context, params *AttendanceListParams, opts ...RequestOption) (*AttendanceListResponse, error) {
	query := url.Values{}

	if params != nil {
		query.Set("date_from", scalar(params.DateFrom))
		query.Set("date_to", scalar(params.DateTo))
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryList(query, "users", params.Users)
		queryList(query, "locations", params.Locations)
		queryList(query, "statuses", params.Statuses)
		queryList(query, "sources", params.Sources)
		queryList(query, "include", params.Include)
	}

	var out AttendanceListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/attendance",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Attendance.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Attendance) ListAll(ctx context.Context, params *AttendanceListParams, opts ...RequestOption) iter.Seq2[AttendanceListRow, error] {
	return func(yield func(AttendanceListRow, error) bool) {
		walked := AttendanceListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none AttendanceListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Record is POST /company/v3/attendance.
//
// Record attendance.
//
// Marks recorded by something you run — a turnstile of your own, a kiosk, your app.
//
// **`datetime` carries its own offset and is the only place time is stated** — send
// `2026-08-10T09:03:00+05:00` and `09:03:00` goes on file with `+05:00` beside it. There
// is no separate timezone field, and the offset is not optional: the wall clock and the
// zone are read off that one value, so they cannot disagree. A mark may not be dated in the
// future, nor more than 24 hours back: older than that is history being rewritten, and
// lateness already computed against the day would move under it.
//
// `status` is `in`, `out` or `break`, and the answer echoes it back in those words, with
// the moment as the instant it was recorded at — the same shapes the listing answers with.
//
// **A mark already on file for the same person, moment and direction is not written
// again**, whether it repeats across two requests or inside one. There is no key you own
// on this resource, so no promise of idempotency is made in general — but a retried upload
// will not double somebody's day. The answer says `created` or `unchanged` per item, with
// the id either way.
//
// The shift a mark belongs to is worked out afterwards, so `shift_id` is yours to send
// only if you already know it.
func (n *Attendance) Record(ctx context.Context, body *AttendanceRecordBody, opts ...RequestOption) (*AttendanceRecordResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out AttendanceRecordResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/attendance",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Delete is DELETE /company/v3/departments/{id}.
//
// Delete a department.
//
// **Refused while anybody is in it** — 409, `department_in_use`.
//
// Once it goes, so do its managers, and there is no way to reconstruct who was in it.
//
// Move the people first, then delete. Somebody dismissed does not count as being in it.
func (n *Departments) Delete(ctx context.Context, id int64, opts ...RequestOption) (*DepartmentsDeleteResponse, error) {
	var out DepartmentsDeleteResponse

	if err := n.do(ctx, request{
		method: http.MethodDelete,
		path:   fmt.Sprintf("/company/v3/departments/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/departments/{id}.
//
// Read one department.
//
// The same keys and the same `include` vocabulary as the listing — `include=managers` adds the
// managers. Somebody else's id is a `404`.
func (n *Departments) Get(ctx context.Context, id int64, params *DepartmentsGetParams, opts ...RequestOption) (*DepartmentsGetResponse, error) {
	query := url.Values{}

	if params != nil {
		queryList(query, "include", params.Include)
	}

	var out DepartmentsGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/departments/%d", id),
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/departments.
//
// List departments.
//
// `include=managers` adds the managers. Without it the key is absent, never null standing in for
// "not asked for".
func (n *Departments) List(ctx context.Context, params *DepartmentsListParams, opts ...RequestOption) (*DepartmentsListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryOpt(query, "search", params.Search)
		queryOpt(query, "updated_since", params.UpdatedSince)
		queryList(query, "include", params.Include)
	}

	var out DepartmentsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/departments",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Departments.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Departments) ListAll(ctx context.Context, params *DepartmentsListParams, opts ...RequestOption) iter.Seq2[DepartmentsListRow, error] {
	return func(yield func(DepartmentsListRow, error) bool) {
		walked := DepartmentsListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none DepartmentsListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Upsert is POST /company/v3/departments/upsert.
//
// Create or update departments.
//
// Up to 100 entries in one call, matched on `external_id`.
//
// **Not matched on the name.** Renaming an entry on your side updates ours, where matching
// on `title` would have created a second and orphaned the first.
func (n *Departments) Upsert(ctx context.Context, body *DepartmentsUpsertBody, opts ...RequestOption) (*DepartmentsUpsertResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out DepartmentsUpsertResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/departments/upsert",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Delete is DELETE /company/v3/documents/{id}.
//
// Delete a document.
//
// Deletes the document, its file records and the stored objects behind them.
//
// **A document everyone has signed cannot be deleted** — `409`, code `document_signed`.
//
// Another company's id is a `404`.
func (n *Documents) Delete(ctx context.Context, id int64, opts ...RequestOption) (*DocumentsDeleteResponse, error) {
	var out DocumentsDeleteResponse

	if err := n.do(ctx, request{
		method: http.MethodDelete,
		path:   fmt.Sprintf("/company/v3/documents/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/documents/{id}.
//
// Read one document.
//
// The same keys the listing answers with, and the same `include` vocabulary. Another company's id
// is a `404`.
func (n *Documents) Get(ctx context.Context, id int64, params *DocumentsGetParams, opts ...RequestOption) (*DocumentsGetResponse, error) {
	query := url.Values{}

	if params != nil {
		queryList(query, "include", params.Include)
	}

	var out DocumentsGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/documents/%d", id),
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/documents.
//
// List documents.
//
// The company's paperwork, oldest first, paged on a cursor.
//
// **Both kinds of document.** `party` is `employee` for a document about one of your
// people and `counterparty` for one a counterparty signs. Filter with `party=employee` or
// `party=counterparty`; leave it out for both.
//
// **`signature.state` is derived, because there is no column for it.** `none` when nobody
// has to sign, then `rejected`, `revoked`, `pending` in that order of precedence, and
// `signed` only when every signer has. Refusal outranks an outstanding signature, so a
// document one person refused reads `rejected` even while others are still pending.
// `completed_at` is set only for `signed`, and is when the last signer signed.
//
// **Dates are three shapes in one payload.** `start_date`, `end_date` and
// `expiration_date` are plain `YYYY-MM-DD` — they are date columns and carry no time or
// zone. `created_at` is an instant with an offset.
//
// `expires_after` and `expires_before` are a plain range of dates.
//
// `effective_from` and `effective_to` select documents valid during a window: a document
// starts on or before your `effective_to` and either has no end date or ends on or after
// your `effective_from`. Both are required together. **A document with no `start_date`
// never matches** — a row with no interval cannot overlap one.
//
// `locations`, `departments`, `positions` and `user_filters` all reach the document
// through the person it is about, so **a document with no `user_id` matches none of
// them**. The web app files company-level documents that way.
//
// **`updated_since` reads only what changed**, as an instant rather than a date so a
// caller polling every few minutes can say which minute. Two things it will not show you,
// and both are ours rather than yours: a document changed by a path inside the product
// that writes the row directly keeps its old timestamp, and so does one the back office
// modifies. Signature progress is the common case of the second — a signature completed
// through the web app does not move `updated_at`. Re-read with overlap if you depend on
// catching those.
//
// Documents the product files for its own machinery are never listed.
func (n *Documents) List(ctx context.Context, params *DocumentsListParams, opts ...RequestOption) (*DocumentsListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryOpt(query, "party", params.Party)
		queryList(query, "external_ids", params.ExternalIDs)
		queryList(query, "users", params.Users)
		queryList(query, "locations", params.Locations)
		queryList(query, "departments", params.Departments)
		queryList(query, "positions", params.Positions)
		queryList(query, "user_filters", params.UserFilters)
		queryList(query, "types", params.Types)
		queryList(query, "employment_types", params.EmploymentTypes)
		queryOpt(query, "search", params.Search)
		queryOpt(query, "updated_since", params.UpdatedSince)
		queryOpt(query, "expires_after", params.ExpiresAfter)
		queryOpt(query, "expires_before", params.ExpiresBefore)
		queryOpt(query, "effective_from", params.EffectiveFrom)
		queryOpt(query, "effective_to", params.EffectiveTo)
		queryList(query, "include", params.Include)
	}

	var out DocumentsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/documents",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Documents.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Documents) ListAll(ctx context.Context, params *DocumentsListParams, opts ...RequestOption) iter.Seq2[DocumentsListRow, error] {
	return func(yield func(DocumentsListRow, error) bool) {
		walked := DocumentsListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none DocumentsListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Upsert is POST /company/v3/documents/upsert.
//
// File documents.
//
// Up to 100 documents in one call, matched on `external_id`, which is required here.
//
// **The bytes go up separately.** `POST /company/v3/files` takes one file and answers an
// id; send that as `file_id`. A file may be claimed once: one already hanging off a
// document is refused rather than moved.
//
// **An uploaded file waits to be claimed and is not cleaned up.** A batch refused at the
// hundredth document leaves all hundred files staged and still valid, so a retry should
// send the same `file_id`s rather than uploading them again.
//
// `parent_external_id` links a supplementary agreement to what it amends, by your key
// rather than ours. It resolves against documents already filed — not against another item
// of the same batch — and a key matching nothing clears the link rather than failing.
//
// `user_id` is required and may name somebody who has left: termination paperwork is filed
// after a dismissal, which is exactly when it is needed.
//
// **No signers.** Signing runs conversions and outbound calls that do not belong under a
// versioned contract, so documents filed here are always `party: employee`. Signature
// state is readable; creating a signing request is not offered yet.
//
// `author_id` is set to the subject: this token authenticates a company, not a person.
func (n *Documents) Upsert(ctx context.Context, body *DocumentsUpsertBody, opts ...RequestOption) (*DocumentsUpsertResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out DocumentsUpsertResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/documents/upsert",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Upload is POST /company/v3/files.
//
// Upload a file.
//
// One file, `multipart/form-data`, field name `file`. The only route on this surface that
// is not JSON.
//
// It answers an `id`. That id is what `POST /company/v3/documents/upsert` takes as
// `file_id`; until a document claims it the file belongs to nothing.
//
// Up to 10 MB. `pdf`, `doc`, `docx`, `xls`, `xlsx`, `jpg`, `jpeg`, `png` — the content is
// checked, not just the extension. `name` is optional and defaults to the uploaded
// filename without its extension.
//
// `url` is signed and short-lived: read it, do not store it. Read the document again to
// get a fresh one.
func (n *Files) Upload(ctx context.Context, form *FilesUploadForm, opts ...RequestOption) (*FilesUploadResponse, error) {
	if form == nil || form.File == nil {
		return nil, errNoFile
	}

	fields := url.Values{}

	queryOpt(fields, "name", form.Name)

	filename := form.Filename

	if filename == "" {
		filename = "upload"
	}

	var out FilesUploadResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/files",
		form:   &formBody{file: form.File, filename: filename, fields: fields},
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Delete is DELETE /company/v3/locations/{id}.
//
// Delete a location.
//
// **Refused while anybody works there** — 409, `location_in_use` — counting both the
// location on a person's record and the several they may also be assigned to.
//
// **Refused while anything sits beneath it** — 409, `location_has_children`. Deleting a
// parent detaches its whole subtree and destroys the rows describing the ancestry, and
// this API cannot express a hierarchy at all, so you would not be able to see what you
// had taken apart.
//
// Once it goes, so do its managers, its device assignments and any auto-scheduler
// configured for it; devices, schedules, tasks and approval routes keep working with the
// location set to null. None of that is recoverable and none of it is logged.
//
// Move the people first, then delete.
func (n *Locations) Delete(ctx context.Context, id int64, opts ...RequestOption) (*LocationsDeleteResponse, error) {
	var out LocationsDeleteResponse

	if err := n.do(ctx, request{
		method: http.MethodDelete,
		path:   fmt.Sprintf("/company/v3/locations/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/locations/{id}.
//
// Read one location.
//
// The same keys the listing answers with. Somebody else's id is a `404`, where asking the listing
// for it answers `200` with an empty array and leaves you counting.
func (n *Locations) Get(ctx context.Context, id int64, params *LocationsGetParams, opts ...RequestOption) (*LocationsGetResponse, error) {
	query := url.Values{}

	if params != nil {
		queryList(query, "include", params.Include)
	}

	var out LocationsGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/locations/%d", id),
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/locations.
//
// List locations.
//
// Ordered by `id` and paged on a cursor: no page number, no total, and nothing repeated
// or skipped while the list is written to. A cursor issued for another ordering is
// refused rather than silently restarting the list.
//
// Coordinates are numbers, and a latitude of exactly 0 is a coordinate rather than a
// missing one.
//
// `include=managers` adds the employees who manage the location, as on departments and
// user filters.
func (n *Locations) List(ctx context.Context, params *LocationsListParams, opts ...RequestOption) (*LocationsListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryOpt(query, "search", params.Search)
		queryList(query, "include", params.Include)
		queryList(query, "codes", params.Codes)
		queryOpt(query, "updated_since", params.UpdatedSince)
	}

	var out LocationsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/locations",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Locations.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Locations) ListAll(ctx context.Context, params *LocationsListParams, opts ...RequestOption) iter.Seq2[LocationsListRow, error] {
	return func(yield func(LocationsListRow, error) bool) {
		walked := LocationsListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none LocationsListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Upsert is POST /company/v3/locations/upsert.
//
// Create or update locations.
//
// Up to 100 entries in one call, matched on `external_id`.
//
// **Not matched on the name.** Renaming an entry on your side updates ours, where matching
// on `title` would have created a second and orphaned the first.
//
// `code` and `external_id` are different things and both are kept: `code` is a label you
// fill in and we never validate, `external_id` is what the match runs on. Coordinates and
// radius are set here too.
func (n *Locations) Upsert(ctx context.Context, body *LocationsUpsertBody, opts ...RequestOption) (*LocationsUpsertResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out LocationsUpsertResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/locations/upsert",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/payroll/payslips.
//
// List payslips.
//
// Reading only. Nothing on this surface creates or changes a payslip.
//
// **Amounts carry their currency.** A payslip that carries none — one built from a salary
// filed before the field was required — answers the company's own currency rather than
// null, so you do not have to invent that fallback yourself.
//
// **The line items are here**: `additions`, `deductions` and `allowances`, each with its
// `title`, `value` and `pre_tax`, which is what decides whether an amount lands before tax
// is worked out.
//
// **`loan_repaid` is what was taken back against an advance.** An advance is not a
// deduction: it becomes a loan and is repaid on scheduled days, so it never appears in
// `deductions` and a caller subtracting the line items from the total will be out by
// exactly this amount. The figure folds together loan repayments and one-off loan
// adjustments, because that is how the calculation records them.
//
// **The parts do not add up to `take_home`, and are not meant to.** Taxes and the gross
// figure are not published here, so what you get is what was added, taken off and repaid —
// not a derivation of the total.
//
// **`updated_since` matters more here than anywhere**: a payslip is recalculated and moves
// from `draft` to `approved` to `paid`, so without it a caller re-reads every month
// forever. `months` is `YYYY-MM`, and takes a list, so a quarter is one request.
func (n *PayrollPayslips) List(ctx context.Context, params *PayrollPayslipsListParams, opts ...RequestOption) (*PayrollPayslipsListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryList(query, "users", params.Users)
		queryList(query, "statuses", params.Statuses)
		queryList(query, "months", params.Months)
		queryOpt(query, "updated_since", params.UpdatedSince)
	}

	var out PayrollPayslipsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/payroll/payslips",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Payroll.Payslips.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *PayrollPayslips) ListAll(ctx context.Context, params *PayrollPayslipsListParams, opts ...RequestOption) iter.Seq2[PayrollPayslipsListRow, error] {
	return func(yield func(PayrollPayslipsListRow, error) bool) {
		walked := PayrollPayslipsListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none PayrollPayslipsListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Delete is DELETE /company/v3/positions/{id}.
//
// Delete a position.
//
// **Refused while anybody holds it** — 409, `position_in_use`.
//
// Beyond the people: deleting a position destroys its auto-scheduler staffing
// configuration outright, so a rota that says "two bakers on nights" stops saying it.
//
// Move the people first, then delete.
func (n *Positions) Delete(ctx context.Context, id int64, opts ...RequestOption) (*PositionsDeleteResponse, error) {
	var out PositionsDeleteResponse

	if err := n.do(ctx, request{
		method: http.MethodDelete,
		path:   fmt.Sprintf("/company/v3/positions/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/positions/{id}.
//
// Read one position.
//
// The same keys the listing answers with. A position carries no managers, so asking to include
// them is refused rather than answered with an empty list.
func (n *Positions) Get(ctx context.Context, id int64, params *PositionsGetParams, opts ...RequestOption) (*PositionsGetResponse, error) {
	query := url.Values{}

	if params != nil {
		queryList(query, "include", params.Include)
	}

	var out PositionsGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/positions/%d", id),
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/positions.
//
// List positions.
//
// A position carries no managers, so asking to include them is refused rather than answered with
// an empty list.
func (n *Positions) List(ctx context.Context, params *PositionsListParams, opts ...RequestOption) (*PositionsListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryOpt(query, "search", params.Search)
		queryOpt(query, "updated_since", params.UpdatedSince)
		queryList(query, "include", params.Include)
	}

	var out PositionsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/positions",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Positions.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Positions) ListAll(ctx context.Context, params *PositionsListParams, opts ...RequestOption) iter.Seq2[PositionsListRow, error] {
	return func(yield func(PositionsListRow, error) bool) {
		walked := PositionsListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none PositionsListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Upsert is POST /company/v3/positions/upsert.
//
// Create or update positions.
//
// Up to 100 entries in one call, matched on `external_id`.
//
// **Not matched on the name.** Renaming an entry on your side updates ours, where matching
// on `title` would have created a second and orphaned the first.
func (n *Positions) Upsert(ctx context.Context, body *PositionsUpsertBody, opts ...RequestOption) (*PositionsUpsertResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out PositionsUpsertResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/positions/upsert",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Create is POST /company/v3/schedules.
//
// Create schedules.
//
// Up to 25 schedules in one call, each a kind of day, the days it falls on, and the people
// it is for. `type` is per item, so a rota and the absences inside it go together.
//
// There are no repeat patterns: send the dates. If you want every Monday, say which
// Mondays.
//
// **Send the rota as one call, not as twenty.** Everyone named is notified, once per call —
// so the same twenty schedules sent one at a time buzz in somebody's pocket twenty times.
//
// **This is the write to send an `Idempotency-Key` with.** A schedule carries no key of
// yours, so a retry after a timeout files the rota a second time unless the header tells us
// it is the same attempt.
//
// **Schedules carry no key you own**, so a resend duplicates rather than converges. A
// `422` means none of the batch landed and is safe to fix and send again; a timeout is
// not — read back before retrying. The answer lists what was created, in the order sent.
//
// **`type` decides what else is required.** `work` needs `timezone` and either
// `start`/`end` or `shifts`. `free` — a day with hours to make up rather than hours to
// keep — needs `timezone`, `start`, `end`, and takes `time_planned`. `leave` needs only
// `leave_type`.
//
// `start` and `end` are clock times, `HH:MM:SS`, read in `timezone` — not instants, whatever
// a generated client calls the field. `timezone` is a fixed offset, `+05:00` or `Z`.
//
// **The answer is not an echo of the request, so read it.** A day with two or more `shifts`
// takes its `start`, `end` and `time_planned` from them and comes back with `is_split`
// true. A day with exactly one shift is not a split day: the hours move onto the day itself
// and `shifts` comes back empty. `time_planned` for a worked day is always computed —
// the hours less the break — never taken from what you sent.
//
// **Every span is seconds**, `break_time` and `grace_start`/`grace_end` included. Grace is
// stored to the minute, so send a multiple of 60; the maximum is 3600.
//
// **Grace is not the same as a boundary.** It is how far past the start a person may arrive
// and still be credited from the shift boundary. How far outside the shift a punch is
// collected at all is a company setting and is not on this endpoint.
//
// There is no `title` — one is generated and it means nothing to you. One schedule takes up
// to 366 dates, 200 people and 8 shifts.
//
// This write has nothing of your own to match a second attempt against, so a retry of it is
// safe only with WithIdempotencyKey.
func (n *Schedules) Create(ctx context.Context, body *SchedulesCreateBody, opts ...RequestOption) (*SchedulesCreateResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out SchedulesCreateResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/schedules",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Delete is DELETE /company/v3/schedules/{id}.
//
// Delete a schedule.
//
// Everyone who was on it is notified, the same way a change to it would notify them.
//
// **A default schedule is refused** — 409, `schedule_is_default`. It is what the company
// falls back to.
//
// **An open shift is refused** — 409, `schedule_is_open`. Deleting one also removes its
// siblings and recomputes their days, which this API can neither create nor show you. Use
// the web application.
func (n *Schedules) Delete(ctx context.Context, id int64, opts ...RequestOption) (*SchedulesDeleteResponse, error) {
	var out SchedulesDeleteResponse

	if err := n.do(ctx, request{
		method: http.MethodDelete,
		path:   fmt.Sprintf("/company/v3/schedules/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/schedules/{id}.
//
// Read one schedule.
//
// Exactly what creating it answered with — the same keys, the shifts and the people
// included, since those are the schedule rather than an optional extra.
//
// Worth reading back after a create: a day with two or more shifts takes its `start`,
// `end` and `time_planned` from them, and a day with exactly one has the shift folded into
// it and comes back with `shifts` empty.
//
// There is no listing of schedules. `GET /company/v3/timesheets` answers what a person is
// scheduled for on a day, which is the question a rota is usually asked.
func (n *Schedules) Get(ctx context.Context, id int64, opts ...RequestOption) (*SchedulesGetResponse, error) {
	var out SchedulesGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/schedules/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/tasks/{id}.
//
// Read one task.
//
// The same keys the listing answers with, and the same `include` vocabulary. Another company's id
// is a `404`.
func (n *Tasks) Get(ctx context.Context, id int64, params *TasksGetParams, opts ...RequestOption) (*TasksGetResponse, error) {
	query := url.Values{}

	if params != nil {
		queryList(query, "include", params.Include)
	}

	var out TasksGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/tasks/%d", id),
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/tasks.
//
// List tasks.
//
// Work as we hold it: what was issued, what became of it, and how it measured up.
//
// `kpi_fact` against `kpi_plan`, plus `time_worked`, are the point of reading a task back
// — what was asked for, what was achieved, how long it took. `status` says where it got
// to.
//
// **Twenty-four fields, not the forty the table has.** Eight of the rest configure how the
// mobile application behaves while the job is done — whether it demands a photo, records a
// location, keeps the steps in order. That is a task template's business, and neither
// useful nor settable here.
//
// `include=items` adds the steps, `include=managers` the people who approve or are
// notified. Without them the keys are absent, never null standing in for "not asked for".
//
// **`updated_since` is what an export should page on**, as an instant: a task moves
// through its statuses inside a working day, so a caller polling for completions needs to
// say which minute.
//
// `statuses` takes `created`, `started`, `paused`, `completed`, `incompleted` and
// `pastdue`. Three more exist in the database and none is offered: `finished` and
// `unfinished` are deprecated spellings, and `pending` is reached only through approval.
// A value that cannot be explained is worse than one that is absent.
//
// **Oldest first.** A first call lands on the earliest task this company ever issued,
// which for a long-standing one is years back. That order is what lets a full export
// finish in one walk, and it is not what you want for "what happened lately": ask with
// `updated_since`, or narrow with `due_from` and `due_to`.
func (n *Tasks) List(ctx context.Context, params *TasksListParams, opts ...RequestOption) (*TasksListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryList(query, "external_ids", params.ExternalIDs)
		queryList(query, "users", params.Users)
		queryList(query, "categories", params.Categories)
		queryList(query, "statuses", params.Statuses)
		queryOpt(query, "active", params.Active)
		queryOpt(query, "search", params.Search)
		queryOpt(query, "due_from", params.DueFrom)
		queryOpt(query, "due_to", params.DueTo)
		queryOpt(query, "updated_since", params.UpdatedSince)
		queryList(query, "include", params.Include)
	}

	var out TasksListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/tasks",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Tasks.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Tasks) ListAll(ctx context.Context, params *TasksListParams, opts ...RequestOption) iter.Seq2[TasksListRow, error] {
	return func(yield func(TasksListRow, error) bool) {
		walked := TasksListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none TasksListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Upsert is POST /company/v3/tasks/upsert.
//
// Issue tasks.
//
// Up to 100 pieces of work in one call, matched on `external_id`, which is required here
// — a task has no natural key of its own, since the same round is issued every week under
// the same title.
//
// **`status` is not accepted, and neither are the timestamps around it.** The product
// moves a task through its lifecycle with events — completing notifies, approval routes,
// reopening makes it pastdue again — and a status written straight onto the row fires none
// of that. Issue the work here; read where it got to with `GET /company/v3/tasks`.
//
// **Eight fields are not accepted either** — `req_photo`, `req_sequence`, `gallery`,
// `get_location`, `get_timing`, `is_keep_status`, `req_approve`, `req_notify`. They
// configure the mobile application, not the job.
//
// **Where the work sits is taken from whoever it is for.** Omit `location_id`,
// `department_id` and `position_id` and they come from the assignee — your system knows
// the person, not our org chart. Send them to override.
//
// **`items` is an exception to the omitted-field rule**: sending it replaces the steps outright,
// because a step carries no key
// to match an incoming one against — and replacing them discards the completion the person
// doing the work recorded. Omit the key to leave them alone. `managers` likewise states
// who approves now rather than adding to them.
//
// `kpi_plan` has no "unset" — the column is NOT NULL with a default of 0, so an omitted
// plan is a plan of zero.
func (n *Tasks) Upsert(ctx context.Context, body *TasksUpsertBody, opts ...RequestOption) (*TasksUpsertResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out TasksUpsertResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/tasks/upsert",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/timesheets.
//
// Timesheets.
//
// One row per person per calendar day of the window: what was planned, and — when asked
// for — what happened and how the two differ.
//
// **`planned` alone is the timesheet grid** — who was meant to work, when, and what kind
// of day it was. `include=actual` adds what was recorded, `include=variance` adds the
// difference. Either one is what makes the request expensive, because both require
// matching the clock-ins.
//
// **The four variance numbers are not additive.** Arriving three minutes late produces
// `time_late` 180 and `time_underworked` 180 — the same minutes, counted once as lateness
// and once as unfilled plan. Summing them double-counts.
//
// **`planned: null` means no schedule at all for that day.** It is rarer than it sounds: a
// company created with default settings carries a work and a leave schedule covering four
// years, so ordinary days off arrive as `type: leave` rather than as an absent plan. A day
// that was scheduled and not worked is the other case — a plan, an empty `actual`, and
// `time_underworked` equal to the whole planned time.
//
// **There is no `per_page`.** How many rows fifty people produce depends on the window and
// on who is scheduled, which the caller cannot predict and we can: the page is sized to a
// row budget instead, and `meta.users_per_page` reports what that came to. Paging walks
// people, so one person's whole period always arrives on a single page and a monthly total
// never has to be assembled across two.
//
// Holidays are not marked: take them from your own calendar. Drafts are never returned.
func (n *Timesheets) List(ctx context.Context, params *TimesheetsListParams, opts ...RequestOption) (*TimesheetsListResponse, error) {
	query := url.Values{}

	if params != nil {
		query.Set("date_from", scalar(params.DateFrom))
		query.Set("date_to", scalar(params.DateTo))
		queryOpt(query, "cursor", params.Cursor)
		queryList(query, "users", params.Users)
		queryList(query, "locations", params.Locations)
		queryList(query, "departments", params.Departments)
		queryList(query, "positions", params.Positions)
		queryOpt(query, "employment", params.Employment)
		queryList(query, "include", params.Include)
	}

	var out TimesheetsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/timesheets",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Timesheets.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Timesheets) ListAll(ctx context.Context, params *TimesheetsListParams, opts ...RequestOption) iter.Seq2[TimesheetsListRow, error] {
	return func(yield func(TimesheetsListRow, error) bool) {
		walked := TimesheetsListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none TimesheetsListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Delete is DELETE /company/v3/user-filters/{id}.
//
// Delete a user filter.
//
// Members do not block it — a filter is a label, and a caller who keeps their filters in
// step re-creates it on the next sync.
//
// **Refused while an approval route points at it** — 409, `user_filter_in_use`. The route
// would survive with nobody to approve through it and quietly stop routing. Change the
// route first.
func (n *UserFilters) Delete(ctx context.Context, id int64, opts ...RequestOption) (*UserFiltersDeleteResponse, error) {
	var out UserFiltersDeleteResponse

	if err := n.do(ctx, request{
		method: http.MethodDelete,
		path:   fmt.Sprintf("/company/v3/user-filters/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/user-filters/{id}.
//
// Read one user filter.
//
// The same keys and the same `include` vocabulary as the listing. Somebody else's id is a `404`.
func (n *UserFilters) Get(ctx context.Context, id int64, params *UserFiltersGetParams, opts ...RequestOption) (*UserFiltersGetResponse, error) {
	query := url.Values{}

	if params != nil {
		queryList(query, "include", params.Include)
	}

	var out UserFiltersGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/user-filters/%d", id),
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/user-filters.
//
// List user filters.
//
// `include=managers` adds the managers, as on departments.
func (n *UserFilters) List(ctx context.Context, params *UserFiltersListParams, opts ...RequestOption) (*UserFiltersListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryOpt(query, "search", params.Search)
		queryOpt(query, "updated_since", params.UpdatedSince)
		queryList(query, "include", params.Include)
	}

	var out UserFiltersListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/user-filters",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.UserFilters.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *UserFilters) ListAll(ctx context.Context, params *UserFiltersListParams, opts ...RequestOption) iter.Seq2[UserFiltersListRow, error] {
	return func(yield func(UserFiltersListRow, error) bool) {
		walked := UserFiltersListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none UserFiltersListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Upsert is POST /company/v3/user-filters/upsert.
//
// Create or update user filters.
//
// Up to 100 entries in one call, matched on `external_id`.
//
// **Not matched on the name.** Renaming an entry on your side updates ours, where matching
// on `title` would have created a second and orphaned the first.
func (n *UserFilters) Upsert(ctx context.Context, body *UserFiltersUpsertBody, opts ...RequestOption) (*UserFiltersUpsertResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out UserFiltersUpsertResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/user-filters/upsert",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/user-requests/{id}.
//
// Read one request.
//
// Answers with `content`, unlike the listing: one row is one shape to make sense of, not a
// hundred.
func (n *UserRequests) Get(ctx context.Context, id int64, opts ...RequestOption) (*UserRequestsGetResponse, error) {
	var out UserRequestsGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/user-requests/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/user-requests.
//
// List requests.
//
// What people asked for, and what became of it — leave, schedule changes, corrections and
// money.
//
// **Read this to learn that a timesheet moved.** Half of everything here is a batch
// clock-in correction: somebody forgot to punch, a manager approved the fix, and the
// attendance for those days changed after the fact. Only 48 per cent of those are approved
// within three days of the day they correct, and a third reach back more than a week — so
// a caller that pulled a timesheet last week cannot assume it still holds. Attendance
// carries no timestamps of its own, which makes this listing paged on `updated_since` the
// only signal that anything has moved.
//
// `period` is the field that makes that usable: the span of days a request concerns,
// wherever its kind happens to keep them. A clock-in correction keeps them inside the
// punches, a leave request as a period or a list, a request for a certificate not at all —
// both ends are null there rather than invented.
//
// `subtype` is the second half of `type`, and the product keeps it in two different places
// — `content.type` for most kinds, `content.leave_type` for leave. It is answered as one
// field, and `subtypes` filters on both.
//
// `comment` is what the author wrote when filing it, ordinarily the reason. Comments the
// workflow writes itself — on acknowledgement, or when a spawned task closes — are not
// answered here.
//
// **Oldest first.** A first call lands on the earliest request this company ever filed,
// which for a long-standing one is years back. That order is what lets a full export
// finish in one walk, and it is not what you want for "what changed lately": ask with
// `updated_since`.
//
// `content` is behind `include=content`: its shape depends on the kind, and one schema
// describes one shape everywhere else on this surface.
//
// **Reading only.** Creating a request enters a workflow — approval routes resolve,
// approvers are notified, tasks are spawned — and approving one is a person's decision that
// a dismissal application or a sick note gives weight to.
func (n *UserRequests) List(ctx context.Context, params *UserRequestsListParams, opts ...RequestOption) (*UserRequestsListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryList(query, "types", params.Types)
		queryList(query, "statuses", params.Statuses)
		queryList(query, "subtypes", params.Subtypes)
		queryList(query, "users", params.Users)
		queryOpt(query, "updated_since", params.UpdatedSince)
		queryList(query, "include", params.Include)
	}

	var out UserRequestsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/user-requests",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.UserRequests.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *UserRequests) ListAll(ctx context.Context, params *UserRequestsListParams, opts ...RequestOption) iter.Seq2[UserRequestsListRow, error] {
	return func(yield func(UserRequestsListRow, error) bool) {
		walked := UserRequestsListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none UserRequestsListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Dismiss is POST /company/v3/users/dismiss.
//
// Dismiss employees.
//
// Up to 100 people in one call, each named by `external_id` or by `id`, exactly one per
// item.
//
// **This is dismissal, not erasure.** The record stays and stays readable: the person
// appears under `status=dismissed` with `dismissed_at` set, and the seat is freed for
// somebody else. It is the shape leaving actually has — a nightly sync noticing that
// twelve people are no longer on the roster.
//
// **There is no hard delete on this API.** Erasing a person takes their attendance,
// payroll, documents and bank details with them, with no way to undo it, and one ability
// grants this whole API — an integrator that syncs your roster cannot be given that
// without also being given erasure. If a retention obligation needs it, ask us.
//
// **Somebody already gone answers `already_dismissed`** rather than failing, which matters
// here because dismissing frees the key for a new hire.
//
// **If a key is held by two people** — one who left and one hired since — the living one
// is the one dismissed.
//
// What happens that you cannot see, and cannot undo:
//
//   - Their sessions end immediately and their phone is unpaired.
//   - Every terminal at their locations is told to forget their face. That is sent once and
//     not retried, so a terminal that is offline at the time keeps admitting them.
//   - **Requests still waiting on them to approve are cancelled**, not just their own. Dismiss
//     a manager and their team's pending vacation requests are cancelled with them.
//   - An offboarding process starts, if the company has one configured.
//   - `date_leave` is filled in with today's date if it was empty.
func (n *Users) Dismiss(ctx context.Context, body *UsersDismissBody, opts ...RequestOption) (*UsersDismissResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out UsersDismissResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/users/dismiss",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/users/{id}.
//
// Read one employee.
//
// One employee by our id. Prefer it over `?ids=` when you expect exactly one: someone who
// is not yours answers `404`, where the listing answers `200` with an empty array, so a
// status can be branched on without counting.
//
// Reachable for a dismissed person too.
func (n *Users) Get(ctx context.Context, id int64, params *UsersGetParams, opts ...RequestOption) (*UsersGetResponse, error) {
	query := url.Values{}

	if params != nil {
		queryList(query, "include", params.Include)
	}

	var out UsersGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/users/%d", id),
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/users.
//
// List employees.
//
// The roster as we hold it. Every scalar is always present; a relation appears only when
// `include` names it.
//
// **`status` decides whether the people who left are in the answer** — `active` by
// default, `dismissed` for only them, `all` for both. A dismissed employee keeps their
// record: `date_leave` says when they were let go and `dismissed_at` when the record was
// closed. A sync that never asks for them cannot learn that anyone left, so ask
// periodically even if your day-to-day reads are `active`.
//
// **`external_ids` is the other half of the roster write.** Ask with the same keys you
// sent to `/users/upsert` and reconcile without keeping a map of our ids.
//
// `updated_since` reads only what changed, against the `updated_at` every row carries.
func (n *Users) List(ctx context.Context, params *UsersListParams, opts ...RequestOption) (*UsersListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryOpt(query, "search", params.Search)
		queryOpt(query, "updated_since", params.UpdatedSince)
		queryOpt(query, "status", params.Status)
		queryList(query, "ids", params.IDs)
		queryList(query, "codes", params.Codes)
		queryList(query, "external_ids", params.ExternalIDs)
		queryList(query, "locations", params.Locations)
		queryList(query, "departments", params.Departments)
		queryList(query, "positions", params.Positions)
		queryList(query, "user_filters", params.UserFilters)
		queryList(query, "employment", params.Employment)
		queryList(query, "include", params.Include)
	}

	var out UsersListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/users",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Users.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Users) ListAll(ctx context.Context, params *UsersListParams, opts ...RequestOption) iter.Seq2[UsersListRow, error] {
	return func(yield func(UsersListRow, error) bool) {
		walked := UsersListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none UsersListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Upsert is POST /company/v3/users/upsert.
//
// Create or update employees.
//
// The roster, written in batches of up to 100. `external_id` is optional here and behaves
// as it does everywhere: an item carrying one updates the person it names, an item without
// one always creates. `first_name`, `role` and `location_id` are the minimum.
//
// The field set is what an HR system holds about an employee.
//
// A person created here reaches the turnstiles of their location, and every location's
// devices are told once for the whole batch rather than once per person.
func (n *Users) Upsert(ctx context.Context, body *UsersUpsertBody, opts ...RequestOption) (*UsersUpsertResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out UsersUpsertResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/users/upsert",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Create is POST /company/v3/webhooks.
//
// Create a webhook endpoint.
//
// Connect an endpoint. Answers 201 with the signing secret it will use.
//
// **No `external_id`, and no upsert**, unlike every other write on this surface. Those
// mirror something your system already holds and must match on its own key; an endpoint is
// created here and its identity is ours, so there is nothing to match against.
//
// The secret is generated, never accepted: a caller-chosen signing key is a caller-chosen
// weakness. Replace it with `POST /company/v3/webhooks/{id}/secret`.
//
// Deliveries carry `X-Clockster-Event`, `X-Clockster-Delivery` (constant across retries,
// so a repeat can be recognised), `X-Clockster-Timestamp` and `X-Clockster-Signature` —
// `sha256=` HMAC-SHA256 of `timestamp + "." + rawBody` under the secret. The body is
// `{"id", "event", "occurred_at", "data"}`.
//
// This write has nothing of your own to match a second attempt against, so a retry of it is
// safe only with WithIdempotencyKey.
func (n *Webhooks) Create(ctx context.Context, body *WebhooksCreateBody, opts ...RequestOption) (*WebhooksCreateResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out WebhooksCreateResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   "/company/v3/webhooks",
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Delete is DELETE /company/v3/webhooks/{id}.
//
// Delete a webhook endpoint.
//
// Removes the endpoint. What was delivered to it stays readable: the delivery's link is
// nulled rather than cascaded, because the record of what was sent is the company's.
func (n *Webhooks) Delete(ctx context.Context, id int64, opts ...RequestOption) (*WebhooksDeleteResponse, error) {
	var out WebhooksDeleteResponse

	if err := n.do(ctx, request{
		method: http.MethodDelete,
		path:   fmt.Sprintf("/company/v3/webhooks/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/webhooks/{id}.
//
// Read one webhook endpoint.
func (n *Webhooks) Get(ctx context.Context, id int64, opts ...RequestOption) (*WebhooksGetResponse, error) {
	var out WebhooksGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/webhooks/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/webhooks.
//
// List webhook endpoints.
//
// The endpoints this company has connected, and the health of each.
//
// `health` answers what `active` cannot: whether a person switched an endpoint off or a
// run of failures did. Five consecutive permanent failures — a wrong address, refused
// credentials — switch it off, as do twenty transient ones. `disabled_reason` says which.
//
// `secret` is answered in full because verifying a signature is impossible without it.
// The credential we authenticate to the receiver *with* is not: `auth` names the scheme
// and, for basic, the username, and never the password or bearer token.
func (n *Webhooks) List(ctx context.Context, params *WebhooksListParams, opts ...RequestOption) (*WebhooksListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryOpt(query, "active", params.Active)
	}

	var out WebhooksListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/webhooks",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Webhooks.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *Webhooks) ListAll(ctx context.Context, params *WebhooksListParams, opts ...RequestOption) iter.Seq2[WebhooksListRow, error] {
	return func(yield func(WebhooksListRow, error) bool) {
		walked := WebhooksListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none WebhooksListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// RotateSecret is POST /company/v3/webhooks/{id}/secret.
//
// Replace the signing secret.
//
// Answers the endpoint with a new secret.
//
// **Send an `Idempotency-Key`.** The secret is shown once, so a retry after a lost
// response would rotate a second time and leave you holding one that signs nothing — with
// the header, the retry is answered with the secret the first call minted.
//
// Deliveries already queued are signed with whichever secret is current when they are
// actually sent, so accept both for as long as your backlog can be deep — up to about a
// day where transient failures are being retried.
//
// This write has nothing of your own to match a second attempt against, so a retry of it is
// safe only with WithIdempotencyKey.
func (n *Webhooks) RotateSecret(ctx context.Context, id int64, opts ...RequestOption) (*WebhooksRotateSecretResponse, error) {
	var out WebhooksRotateSecretResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   fmt.Sprintf("/company/v3/webhooks/%d/secret", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Update is PUT /company/v3/webhooks/{id}.
//
// Replace a webhook endpoint.
//
// Replaces rather than patches: half a subscription is not a state worth reaching by
// accident.
//
// Saving clears the failure tally and any automatic switch-off — you are saying something
// changed, so what the history counted no longer describes what is there. This is how an
// endpoint switched off by repeated failures is put back into service.
func (n *Webhooks) Update(ctx context.Context, id int64, body *WebhooksUpdateBody, opts ...RequestOption) (*WebhooksUpdateResponse, error) {
	if body == nil {
		return nil, errNoBody
	}

	var out WebhooksUpdateResponse

	if err := n.do(ctx, request{
		method: http.MethodPut,
		path:   fmt.Sprintf("/company/v3/webhooks/%d", id),
		body:   body,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// Get is GET /company/v3/webhooks/deliveries/{id}.
//
// Read one webhook delivery.
//
// Answers with the payload, unlike the listing: one row is not a hundred employee records.
func (n *WebhooksDeliveries) Get(ctx context.Context, id int64, opts ...RequestOption) (*WebhooksDeliveriesGetResponse, error) {
	var out WebhooksDeliveriesGetResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   fmt.Sprintf("/company/v3/webhooks/deliveries/%d", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/webhooks/deliveries.
//
// List webhook deliveries.
//
// What was sent, and what came of it. **Newest first**, unlike every other listing here: a
// delivery log is read to learn what just happened.
//
// `state` is the field to branch on. `delivered` and `failed` are final; `pending` is
// neither — the event is waiting out its backoff and will be tried again. Reading
// `is_successful: false` as failure is the mistake the two stored flags invite.
//
// `payload` is behind `include=payload` and off by default: a page of a hundred deliveries
// is a hundred employee records, and a caller watching its integration's health has no use
// for them. Reading one delivery answers with it either way.
//
// `webhook_id` is null where the endpoint has since been removed.
func (n *WebhooksDeliveries) List(ctx context.Context, params *WebhooksDeliveriesListParams, opts ...RequestOption) (*WebhooksDeliveriesListResponse, error) {
	query := url.Values{}

	if params != nil {
		queryOpt(query, "per_page", params.PerPage)
		queryOpt(query, "cursor", params.Cursor)
		queryList(query, "webhooks", params.Webhooks)
		queryList(query, "events", params.Events)
		queryOpt(query, "successful", params.Successful)
		queryOpt(query, "pending", params.Pending)
		queryOpt(query, "since", params.Since)
		queryList(query, "include", params.Include)
	}

	var out WebhooksDeliveriesListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/webhooks/deliveries",
		query:  query,
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// ListAll walks every page of List and answers a row at a time.
//
//	for row, err := range clockster.Webhooks.Deliveries.ListAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func (n *WebhooksDeliveries) ListAll(ctx context.Context, params *WebhooksDeliveriesListParams, opts ...RequestOption) iter.Seq2[WebhooksDeliveriesListRow, error] {
	return func(yield func(WebhooksDeliveriesListRow, error) bool) {
		walked := WebhooksDeliveriesListParams{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := n.List(ctx, &walked, opts...)
			if err != nil {
				var none WebhooksDeliveriesListRow

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			cursor := page.Meta.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.Cursor = Set(*cursor)
		}
	}
}

// Redeliver is POST /company/v3/webhooks/deliveries/{id}/redeliver.
//
// Send a delivery again.
//
// Queues the recorded event for another attempt, and answers once queued rather than once
// delivered.
//
// A repair tool, now that transient failures retry themselves: this is for the case where
// the receiver was fixed after we had already given up. What goes out is the event as it
// was recorded, so `occurred_at` may be long past — which is why the envelope states it.
//
// This write has nothing of your own to match a second attempt against, so a retry of it is
// safe only with WithIdempotencyKey.
func (n *WebhooksDeliveries) Redeliver(ctx context.Context, id int64, opts ...RequestOption) (*WebhooksDeliveriesRedeliverResponse, error) {
	var out WebhooksDeliveriesRedeliverResponse

	if err := n.do(ctx, request{
		method: http.MethodPost,
		path:   fmt.Sprintf("/company/v3/webhooks/deliveries/%d/redeliver", id),
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}

// List is GET /company/v3/webhooks/events.
//
// List subscribable events.
//
// Every event name an endpoint can subscribe to.
//
// Served as well as specified, so a caller can check at runtime that a name it stored is
// still one we send rather than discovering it on the next save.
func (n *WebhooksEvents) List(ctx context.Context, opts ...RequestOption) (*WebhooksEventsListResponse, error) {
	var out WebhooksEventsListResponse

	if err := n.do(ctx, request{
		method: http.MethodGet,
		path:   "/company/v3/webhooks/events",
	}, &out, opts); err != nil {
		return nil, err
	}

	return &out, nil
}
