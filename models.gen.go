// Code generated from openapi/company-v3.json; DO NOT EDIT.

// Every shape the Company API answers with or accepts.
//
// A field the document marks optional is a pointer in an answer and an Opt in a request: a
// relation is absent from an answer unless `include` asked for it, and a field left out of a
// write keeps whatever is stored where a null one clears it.

package clockster

import (
	"io"
)

// MeResponse is part of what GET /company/v3/me answers.
type MeResponse struct {
	Data MeData `json:"data"`
}

// MeData is part of what GET /company/v3/me answers.
type MeData struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// AttendanceListParams is the query of GET /company/v3/attendance.
type AttendanceListParams struct {
	// Start of the window, inclusive (YYYY-MM-DD).
	DateFrom string `json:"date_from"`

	// End of the window, inclusive (YYYY-MM-DD).
	DateTo string `json:"date_to"`

	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Only rows belonging to these people, by id.
	Users []int64 `json:"users,omitzero"`

	// Only these locations, by id.
	Locations []int64 `json:"locations,omitzero"`

	// Only rows in these states.
	// One of "out", "in", "break".
	Statuses []string `json:"statuses,omitzero"`

	// Only marks recorded this way — a device, the mobile app, or a person entering them by hand.
	// One of "device", "mobile", "frontend", "api", "system".
	Sources []string `json:"sources,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "user", "location", "attachments".
	Include []string `json:"include,omitzero"`
}

// AttendanceListResponse is part of what GET /company/v3/attendance answers.
type AttendanceListResponse struct {
	Data  []AttendanceListRow `json:"data"`
	Links PageLinks           `json:"links"`
	Meta  PageMeta            `json:"meta"`
}

// AttendanceListRow is part of what GET /company/v3/attendance answers.
type AttendanceListRow struct {
	ID          int64                         `json:"id"`
	UserID      int64                         `json:"user_id"`
	LocationID  *int64                        `json:"location_id"`
	Datetime    string                        `json:"datetime"`
	Status      string                        `json:"status"`
	Source      string                        `json:"source"`
	Latitude    *float64                      `json:"latitude"`
	Longitude   *float64                      `json:"longitude"`
	Address     *string                       `json:"address"`
	Comment     *string                       `json:"comment"`
	User        *EmployeeShort                `json:"user"`
	Location    *AttendanceListRowLocation    `json:"location"`
	Attachments []AttendanceListRowAttachment `json:"attachments"`
}

type EmployeeShort struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Code       *string `json:"code"`
	FirstName  string  `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName   string  `json:"last_name"`
}

// AttendanceListRowLocation is part of what GET /company/v3/attendance answers.
type AttendanceListRowLocation struct {
	ID          int64    `json:"id"`
	ExternalID  *string  `json:"external_id"`
	Code        *string  `json:"code"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	Radius      int64    `json:"radius"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

// AttendanceListRowAttachment is part of what GET /company/v3/attendance answers.
type AttendanceListRowAttachment struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Format      string  `json:"format"`
	URL         string  `json:"url"`
	CreatedAt   string  `json:"created_at"`
}

type PageLinks struct {
	First any     `json:"first"`
	Last  any     `json:"last"`
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
}

type PageMeta struct {
	Path       string  `json:"path"`
	PerPage    int64   `json:"per_page"`
	NextCursor *string `json:"next_cursor"`
	PrevCursor *string `json:"prev_cursor"`
}

// AttendanceRecordBody is part of what POST /company/v3/attendance takes.
type AttendanceRecordBody struct {
	Attendance []AttendanceRecordAttendanceItem `json:"attendance"`
}

// AttendanceRecordAttendanceItem is part of what POST /company/v3/attendance takes.
type AttendanceRecordAttendanceItem struct {
	UserID     int64      `json:"user_id"`
	LocationID Opt[int64] `json:"location_id,omitzero"`
	ShiftID    Opt[int64] `json:"shift_id,omitzero"`

	// One of "out", "in", "break".
	Status string `json:"status"`

	// An instant, ISO 8601 with an offset.
	Datetime string      `json:"datetime"`
	Comment  Opt[string] `json:"comment,omitzero"`
}

// AttendanceRecordResponse is part of what POST /company/v3/attendance answers.
type AttendanceRecordResponse struct {
	Data []AttendanceRecordRow `json:"data"`
}

// AttendanceRecordRow is part of what POST /company/v3/attendance answers.
type AttendanceRecordRow struct {
	UserID   int64  `json:"user_id"`
	Datetime string `json:"datetime"`
	Status   string `json:"status"`
	ID       int64  `json:"id"`
	Result   string `json:"result"`
}

// DepartmentsDeleteResponse is part of what DELETE /company/v3/departments/{id} answers.
type DepartmentsDeleteResponse struct {
	Data DeleteOutcome `json:"data"`
}

type DeleteOutcome struct {
	ID     int64  `json:"id"`
	Result string `json:"result"`
}

// DepartmentsGetParams is the query of GET /company/v3/departments/{id}.
type DepartmentsGetParams struct {
	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "managers".
	Include []string `json:"include,omitzero"`
}

// DepartmentsGetResponse is part of what GET /company/v3/departments/{id} answers.
type DepartmentsGetResponse struct {
	Data DepartmentsGetData `json:"data"`
}

// DepartmentsGetData is part of what GET /company/v3/departments/{id} answers.
type DepartmentsGetData struct {
	ID          int64           `json:"id"`
	ExternalID  *string         `json:"external_id"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Managers    []EmployeeShort `json:"managers"`
}

// DepartmentsListParams is the query of GET /company/v3/departments.
type DepartmentsListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Free text over the names the section lists.
	Search Opt[string] `json:"search,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "managers".
	Include []string `json:"include,omitzero"`
}

// DepartmentsListResponse is part of what GET /company/v3/departments answers.
type DepartmentsListResponse struct {
	Data  []DepartmentsListRow `json:"data"`
	Links PageLinks            `json:"links"`
	Meta  PageMeta             `json:"meta"`
}

// DepartmentsListRow is part of what GET /company/v3/departments answers.
type DepartmentsListRow struct {
	ID          int64           `json:"id"`
	ExternalID  *string         `json:"external_id"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Managers    []EmployeeShort `json:"managers"`
}

// DepartmentsUpsertBody is part of what POST /company/v3/departments/upsert takes.
type DepartmentsUpsertBody struct {
	Items []DepartmentsUpsertItem `json:"items"`
}

// DepartmentsUpsertItem is part of what POST /company/v3/departments/upsert takes.
type DepartmentsUpsertItem struct {
	ExternalID  string      `json:"external_id"`
	Title       string      `json:"title"`
	Description Opt[string] `json:"description,omitzero"`
}

// DepartmentsUpsertResponse is part of what POST /company/v3/departments/upsert answers.
type DepartmentsUpsertResponse struct {
	Data []UpsertOutcome `json:"data"`
}

type UpsertOutcome struct {
	ExternalID *string `json:"external_id"`
	ID         int64   `json:"id"`
	Result     string  `json:"result"`
}

// DocumentsDeleteResponse is part of what DELETE /company/v3/documents/{id} answers.
type DocumentsDeleteResponse struct {
	Data DeleteOutcome `json:"data"`
}

// DocumentsGetParams is the query of GET /company/v3/documents/{id}.
type DocumentsGetParams struct {
	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "attachments", "signers", "labor_contract".
	Include []string `json:"include,omitzero"`
}

// DocumentsGetResponse is part of what GET /company/v3/documents/{id} answers.
type DocumentsGetResponse struct {
	Data DocumentsGetData `json:"data"`
}

// DocumentsGetData is part of what GET /company/v3/documents/{id} answers.
type DocumentsGetData struct {
	ID               int64                          `json:"id"`
	ExternalID       *string                        `json:"external_id"`
	UserID           int64                          `json:"user_id"`
	AuthorID         int64                          `json:"author_id"`
	Party            string                         `json:"party"`
	Type             string                         `json:"type"`
	Name             string                         `json:"name"`
	ContractNumber   *string                        `json:"contract_number"`
	EmploymentType   *string                        `json:"employment_type"`
	StartDate        *string                        `json:"start_date"`
	EndDate          *string                        `json:"end_date"`
	ExpirationDate   *string                        `json:"expiration_date"`
	ParentDocumentID *int64                         `json:"parent_document_id"`
	Signature        DocumentsGetDataSignature      `json:"signature"`
	CreatedAt        string                         `json:"created_at"`
	UpdatedAt        string                         `json:"updated_at"`
	Attachments      []DocumentsGetDataAttachment   `json:"attachments"`
	Signers          []DocumentsGetDataSigner       `json:"signers"`
	LaborContract    *DocumentsGetDataLaborContract `json:"labor_contract"`
}

// DocumentsGetDataSignature is part of what GET /company/v3/documents/{id} answers.
type DocumentsGetDataSignature struct {
	State       string  `json:"state"`
	CompletedAt *string `json:"completed_at"`
}

// DocumentsGetDataAttachment is part of what GET /company/v3/documents/{id} answers.
type DocumentsGetDataAttachment struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Format      string  `json:"format"`
	URL         string  `json:"url"`
	CreatedAt   string  `json:"created_at"`
}

// DocumentsGetDataSigner is part of what GET /company/v3/documents/{id} answers.
type DocumentsGetDataSigner struct {
	Type      string  `json:"type"`
	Status    string  `json:"status"`
	SignedAt  *string `json:"signed_at"`
	SignedVia *string `json:"signed_via"`
	PartyID   int64   `json:"party_id"`
	PartyType string  `json:"party_type"`
}

// DocumentsGetDataLaborContract is part of what GET /company/v3/documents/{id} answers.
type DocumentsGetDataLaborContract struct {
	ID                 int64   `json:"id"`
	ExternalContractID string  `json:"external_contract_id"`
	IIN                string  `json:"iin"`
	ContractNumber     *string `json:"contract_number"`
	ContractDate       *string `json:"contract_date"`
	BeginDate          *string `json:"begin_date"`
	EndDate            *string `json:"end_date"`
	TerminationDate    *string `json:"termination_date"`
	EstablishedPost    string  `json:"established_post"`
	HasContractFile    bool    `json:"has_contract_file"`
}

// DocumentsListParams is the query of GET /company/v3/documents.
type DocumentsListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Whose side of the document to read.
	// One of "employee", "counterparty".
	Party Opt[string] `json:"party,omitzero"`

	// Only rows carrying these keys of yours. The other half of an upsert: write with your key, read
	// back with it.
	ExternalIDs []string `json:"external_ids,omitzero"`

	// Only rows belonging to these people, by id.
	Users []int64 `json:"users,omitzero"`

	// Only these locations, by id.
	Locations []int64 `json:"locations,omitzero"`

	// Only these departments, by id.
	Departments []int64 `json:"departments,omitzero"`

	// Only these positions, by id.
	Positions []int64 `json:"positions,omitzero"`

	// Only these user filters, by id.
	UserFilters []int64 `json:"user_filters,omitzero"`

	// Only rows of these types.
	// One of "passport", "cv", "diploma", "medical", "photo", "other", "medical_book",
	// "employment_agreement", "termination_of_employment_agreement", "equipment_agreement",
	// "application", "order", "supplementary_agreement", "job_description", "nda",
	// "non_compete_agreement", "data_processing_agreement", "act_of_service_acceptance",
	// "health_and_safety_briefing", "shift_schedule", "letter", "vacation_schedule", "contract",
	// "agreement", "goods_release_note", "reconciliation_act", "return_to_supplier".
	Types []string `json:"types,omitzero"`

	// Only documents covering these employment terms.
	// One of "full_time", "part_time", "irregular_hours", "contract_1", "contract_2",
	// "apprenticeship", "traineeship", "piece_rate", "probation", "outstaffing".
	EmploymentTypes []string `json:"employment_types,omitzero"`

	// Free text over the names the section lists.
	Search Opt[string] `json:"search,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`

	// Expiring after this date (YYYY-MM-DD) — how you find what is about to lapse.
	ExpiresAfter Opt[string] `json:"expires_after,omitzero"`

	// Expiring before this date (YYYY-MM-DD).
	ExpiresBefore Opt[string] `json:"expires_before,omitzero"`

	// In force at or after this date (YYYY-MM-DD).
	EffectiveFrom Opt[string] `json:"effective_from,omitzero"`

	// In force at or before this date (YYYY-MM-DD).
	EffectiveTo Opt[string] `json:"effective_to,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "attachments", "signers", "labor_contract".
	Include []string `json:"include,omitzero"`
}

// DocumentsListResponse is part of what GET /company/v3/documents answers.
type DocumentsListResponse struct {
	Data  []DocumentsListRow `json:"data"`
	Links PageLinks          `json:"links"`
	Meta  PageMeta           `json:"meta"`
}

// DocumentsListRow is part of what GET /company/v3/documents answers.
type DocumentsListRow struct {
	ID               int64                          `json:"id"`
	ExternalID       *string                        `json:"external_id"`
	UserID           int64                          `json:"user_id"`
	AuthorID         int64                          `json:"author_id"`
	Party            string                         `json:"party"`
	Type             string                         `json:"type"`
	Name             string                         `json:"name"`
	ContractNumber   *string                        `json:"contract_number"`
	EmploymentType   *string                        `json:"employment_type"`
	StartDate        *string                        `json:"start_date"`
	EndDate          *string                        `json:"end_date"`
	ExpirationDate   *string                        `json:"expiration_date"`
	ParentDocumentID *int64                         `json:"parent_document_id"`
	Signature        DocumentsListRowSignature      `json:"signature"`
	CreatedAt        string                         `json:"created_at"`
	UpdatedAt        string                         `json:"updated_at"`
	Attachments      []DocumentsListRowAttachment   `json:"attachments"`
	Signers          []DocumentsListRowSigner       `json:"signers"`
	LaborContract    *DocumentsListRowLaborContract `json:"labor_contract"`
}

// DocumentsListRowSignature is part of what GET /company/v3/documents answers.
type DocumentsListRowSignature struct {
	State       string  `json:"state"`
	CompletedAt *string `json:"completed_at"`
}

// DocumentsListRowAttachment is part of what GET /company/v3/documents answers.
type DocumentsListRowAttachment struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Format      string  `json:"format"`
	URL         string  `json:"url"`
	CreatedAt   string  `json:"created_at"`
}

// DocumentsListRowSigner is part of what GET /company/v3/documents answers.
type DocumentsListRowSigner struct {
	Type      string  `json:"type"`
	Status    string  `json:"status"`
	SignedAt  *string `json:"signed_at"`
	SignedVia *string `json:"signed_via"`
	PartyID   int64   `json:"party_id"`
	PartyType string  `json:"party_type"`
}

// DocumentsListRowLaborContract is part of what GET /company/v3/documents answers.
type DocumentsListRowLaborContract struct {
	ID                 int64   `json:"id"`
	ExternalContractID string  `json:"external_contract_id"`
	IIN                string  `json:"iin"`
	ContractNumber     *string `json:"contract_number"`
	ContractDate       *string `json:"contract_date"`
	BeginDate          *string `json:"begin_date"`
	EndDate            *string `json:"end_date"`
	TerminationDate    *string `json:"termination_date"`
	EstablishedPost    string  `json:"established_post"`
	HasContractFile    bool    `json:"has_contract_file"`
}

// DocumentsUpsertBody is part of what POST /company/v3/documents/upsert takes.
type DocumentsUpsertBody struct {
	Documents []DocumentsUpsertDocument `json:"documents"`
}

// DocumentsUpsertDocument is part of what POST /company/v3/documents/upsert takes.
type DocumentsUpsertDocument struct {
	ExternalID string `json:"external_id"`

	// One of "passport", "cv", "diploma", "medical", "photo", "other", "medical_book",
	// "employment_agreement", "termination_of_employment_agreement", "equipment_agreement",
	// "application", "order", "supplementary_agreement", "job_description", "nda",
	// "non_compete_agreement", "data_processing_agreement", "act_of_service_acceptance",
	// "health_and_safety_briefing", "shift_schedule", "letter", "vacation_schedule", "contract",
	// "agreement", "goods_release_note", "reconciliation_act", "return_to_supplier".
	Type           string      `json:"type"`
	UserID         int64       `json:"user_id"`
	Name           Opt[string] `json:"name,omitzero"`
	ContractNumber Opt[string] `json:"contract_number,omitzero"`

	// One of "full_time", "part_time", "irregular_hours", "contract_1", "contract_2",
	// "apprenticeship", "traineeship", "piece_rate", "probation", "outstaffing".
	EmploymentType Opt[string] `json:"employment_type,omitzero"`

	// A plain date, YYYY-MM-DD.
	StartDate Opt[string] `json:"start_date,omitzero"`

	// A plain date, YYYY-MM-DD.
	EndDate Opt[string] `json:"end_date,omitzero"`

	// A plain date, YYYY-MM-DD.
	ExpirationDate   Opt[string] `json:"expiration_date,omitzero"`
	ParentExternalID Opt[string] `json:"parent_external_id,omitzero"`
	FileID           Opt[int64]  `json:"file_id,omitzero"`
}

// DocumentsUpsertResponse is part of what POST /company/v3/documents/upsert answers.
type DocumentsUpsertResponse struct {
	Data []UpsertOutcome `json:"data"`
}

// FilesUploadForm is what POST /company/v3/files takes. It carries bytes rather than JSON,
// which is what makes this the one operation sent as multipart.
type FilesUploadForm struct {
	// File is the bytes to store, read to the end before the request is sent.
	File io.Reader

	// Filename is the name the bytes travel under. Empty is sent as "upload".
	Filename string

	Name Opt[string] `json:"name,omitzero"`
}

// FilesUploadResponse is part of what POST /company/v3/files answers.
type FilesUploadResponse struct {
	Data FilesUploadData `json:"data"`
}

// FilesUploadData is part of what POST /company/v3/files answers.
type FilesUploadData struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Format    string `json:"format"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

// LocationsDeleteResponse is part of what DELETE /company/v3/locations/{id} answers.
type LocationsDeleteResponse struct {
	Data DeleteOutcome `json:"data"`
}

// LocationsGetParams is the query of GET /company/v3/locations/{id}.
type LocationsGetParams struct {
	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "managers".
	Include []string `json:"include,omitzero"`
}

// LocationsGetResponse is part of what GET /company/v3/locations/{id} answers.
type LocationsGetResponse struct {
	Data LocationsGetData `json:"data"`
}

// LocationsGetData is part of what GET /company/v3/locations/{id} answers.
type LocationsGetData struct {
	ID          int64           `json:"id"`
	ExternalID  *string         `json:"external_id"`
	Code        *string         `json:"code"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	Latitude    *float64        `json:"latitude"`
	Longitude   *float64        `json:"longitude"`
	Radius      int64           `json:"radius"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Managers    []EmployeeShort `json:"managers"`
}

// LocationsListParams is the query of GET /company/v3/locations.
type LocationsListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Free text over the names the section lists.
	Search Opt[string] `json:"search,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "managers".
	Include []string `json:"include,omitzero"`

	// Only these employee codes.
	Codes []string `json:"codes,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`
}

// LocationsListResponse is part of what GET /company/v3/locations answers.
type LocationsListResponse struct {
	Data  []LocationsListRow `json:"data"`
	Links PageLinks          `json:"links"`
	Meta  PageMeta           `json:"meta"`
}

// LocationsListRow is part of what GET /company/v3/locations answers.
type LocationsListRow struct {
	ID          int64           `json:"id"`
	ExternalID  *string         `json:"external_id"`
	Code        *string         `json:"code"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	Latitude    *float64        `json:"latitude"`
	Longitude   *float64        `json:"longitude"`
	Radius      int64           `json:"radius"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Managers    []EmployeeShort `json:"managers"`
}

// LocationsUpsertBody is part of what POST /company/v3/locations/upsert takes.
type LocationsUpsertBody struct {
	Items []LocationsUpsertItem `json:"items"`
}

// LocationsUpsertItem is part of what POST /company/v3/locations/upsert takes.
type LocationsUpsertItem struct {
	ExternalID  string       `json:"external_id"`
	Title       string       `json:"title"`
	Description Opt[string]  `json:"description,omitzero"`
	Code        Opt[string]  `json:"code,omitzero"`
	Latitude    Opt[float64] `json:"latitude,omitzero"`
	Longitude   Opt[float64] `json:"longitude,omitzero"`
	Radius      Opt[int64]   `json:"radius,omitzero"`
}

// LocationsUpsertResponse is part of what POST /company/v3/locations/upsert answers.
type LocationsUpsertResponse struct {
	Data []UpsertOutcome `json:"data"`
}

// PayrollPayslipsListParams is the query of GET /company/v3/payroll/payslips.
type PayrollPayslipsListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Only rows belonging to these people, by id.
	Users []int64 `json:"users,omitzero"`

	// Only rows in these states.
	// One of "draft", "approved", "paid".
	Statuses []string `json:"statuses,omitzero"`

	// Only these months, as YYYY-MM.
	Months []string `json:"months,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`
}

// PayrollPayslipsListResponse is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListResponse struct {
	Data []PayrollPayslipsListRow `json:"data"`
	Meta PayrollPayslipsListMeta  `json:"meta"`
}

// PayrollPayslipsListRow is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListRow struct {
	ID         int64                             `json:"id"`
	User       PayrollPayslipsListRowUser        `json:"user"`
	AuthorID   int64                             `json:"author_id"`
	Period     PayrollPayslipsListRowPeriod      `json:"period"`
	Status     string                            `json:"status"`
	Currency   *string                           `json:"currency"`
	TakeHome   float64                           `json:"take_home"`
	Ctc        int64                             `json:"ctc"`
	Salary     PayrollPayslipsListRowSalary      `json:"salary"`
	Additions  []PayrollPayslipsListRowAddition  `json:"additions"`
	Deductions []PayrollPayslipsListRowDeduction `json:"deductions"`
	Allowances []PayrollPayslipsListRowAllowance `json:"allowances"`
	LoanRepaid int64                             `json:"loan_repaid"`
	UpdatedAt  string                            `json:"updated_at"`
}

// PayrollPayslipsListRowUser is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListRowUser struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
}

// PayrollPayslipsListRowPeriod is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListRowPeriod struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Month string `json:"month"`
}

// PayrollPayslipsListRowSalary is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListRowSalary struct {
	BasicRate  *float64 `json:"basic_rate"`
	BasicType  *string  `json:"basic_type"`
	WorkDays   *int64   `json:"work_days"`
	WorkedDays *int64   `json:"worked_days"`
}

// PayrollPayslipsListRowAddition is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListRowAddition struct {
	Title   string  `json:"title"`
	Type    string  `json:"type"`
	Value   int64   `json:"value"`
	PreTax  bool    `json:"pre_tax"`
	Comment *string `json:"comment"`
}

// PayrollPayslipsListRowDeduction is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListRowDeduction struct {
	Title   string  `json:"title"`
	Type    string  `json:"type"`
	Value   float64 `json:"value"`
	PreTax  bool    `json:"pre_tax"`
	Comment *string `json:"comment"`
}

// PayrollPayslipsListRowAllowance is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListRowAllowance struct {
	Title   string  `json:"title"`
	Type    string  `json:"type"`
	Value   int64   `json:"value"`
	PreTax  bool    `json:"pre_tax"`
	Comment *string `json:"comment"`
}

// PayrollPayslipsListMeta is part of what GET /company/v3/payroll/payslips answers.
type PayrollPayslipsListMeta struct {
	PerPage    int64   `json:"per_page"`
	NextCursor *string `json:"next_cursor"`
	PrevCursor *string `json:"prev_cursor"`
}

// PositionsDeleteResponse is part of what DELETE /company/v3/positions/{id} answers.
type PositionsDeleteResponse struct {
	Data DeleteOutcome `json:"data"`
}

// PositionsGetParams is the query of GET /company/v3/positions/{id}.
type PositionsGetParams struct {
	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	Include []string `json:"include,omitzero"`
}

// PositionsGetResponse is part of what GET /company/v3/positions/{id} answers.
type PositionsGetResponse struct {
	Data PositionsGetData `json:"data"`
}

// PositionsGetData is part of what GET /company/v3/positions/{id} answers.
type PositionsGetData struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// PositionsListParams is the query of GET /company/v3/positions.
type PositionsListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Free text over the names the section lists.
	Search Opt[string] `json:"search,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	Include []string `json:"include,omitzero"`
}

// PositionsListResponse is part of what GET /company/v3/positions answers.
type PositionsListResponse struct {
	Data  []PositionsListRow `json:"data"`
	Links PageLinks          `json:"links"`
	Meta  PageMeta           `json:"meta"`
}

// PositionsListRow is part of what GET /company/v3/positions answers.
type PositionsListRow struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// PositionsUpsertBody is part of what POST /company/v3/positions/upsert takes.
type PositionsUpsertBody struct {
	Items []PositionsUpsertItem `json:"items"`
}

// PositionsUpsertItem is part of what POST /company/v3/positions/upsert takes.
type PositionsUpsertItem struct {
	ExternalID  string      `json:"external_id"`
	Title       string      `json:"title"`
	Description Opt[string] `json:"description,omitzero"`
}

// PositionsUpsertResponse is part of what POST /company/v3/positions/upsert answers.
type PositionsUpsertResponse struct {
	Data []UpsertOutcome `json:"data"`
}

// SchedulesCreateBody is part of what POST /company/v3/schedules takes.
type SchedulesCreateBody struct {
	Schedules []SchedulesCreateSchedule `json:"schedules"`
}

// SchedulesCreateSchedule is one of WorkSchedule, FreeSchedule, LeaveSchedule.
//
// Which one is read off "type", so that field carries the value named beside each type
// below rather than any of the values it accepts.
type SchedulesCreateSchedule interface {
	isSchedulesCreateSchedule()
}

func (WorkSchedule) isSchedulesCreateSchedule()  {}
func (FreeSchedule) isSchedulesCreateSchedule()  {}
func (LeaveSchedule) isSchedulesCreateSchedule() {}

type WorkSchedule struct {
	// One of "work", "free", "leave".
	Type         string     `json:"type"`
	Dates        []string   `json:"dates"`
	Users        []int64    `json:"users"`
	LocationID   Opt[int64] `json:"location_id,omitzero"`
	DepartmentID Opt[int64] `json:"department_id,omitzero"`
	PositionID   Opt[int64] `json:"position_id,omitzero"`
	Timezone     string     `json:"timezone"`

	// An instant, ISO 8601 with an offset.
	Start Opt[string] `json:"start,omitzero"`

	// An instant, ISO 8601 with an offset.
	End        Opt[string]         `json:"end,omitzero"`
	BreakTime  Opt[int64]          `json:"break_time,omitzero"`
	GraceStart Opt[int64]          `json:"grace_start,omitzero"`
	GraceEnd   Opt[int64]          `json:"grace_end,omitzero"`
	Shifts     []WorkScheduleShift `json:"shifts,omitzero"`
}

type WorkScheduleShift struct {
	// An instant, ISO 8601 with an offset.
	Start string `json:"start"`

	// An instant, ISO 8601 with an offset.
	End          string     `json:"end"`
	LocationID   Opt[int64] `json:"location_id,omitzero"`
	DepartmentID Opt[int64] `json:"department_id,omitzero"`
	PositionID   Opt[int64] `json:"position_id,omitzero"`
}

type FreeSchedule struct {
	// One of "work", "free", "leave".
	Type         string     `json:"type"`
	Dates        []string   `json:"dates"`
	Users        []int64    `json:"users"`
	LocationID   Opt[int64] `json:"location_id,omitzero"`
	DepartmentID Opt[int64] `json:"department_id,omitzero"`
	PositionID   Opt[int64] `json:"position_id,omitzero"`
	Timezone     string     `json:"timezone"`

	// An instant, ISO 8601 with an offset.
	Start string `json:"start"`

	// An instant, ISO 8601 with an offset.
	End         string     `json:"end"`
	TimePlanned Opt[int64] `json:"time_planned,omitzero"`
}

type LeaveSchedule struct {
	// One of "work", "free", "leave".
	Type         string     `json:"type"`
	Dates        []string   `json:"dates"`
	Users        []int64    `json:"users"`
	LocationID   Opt[int64] `json:"location_id,omitzero"`
	DepartmentID Opt[int64] `json:"department_id,omitzero"`
	PositionID   Opt[int64] `json:"position_id,omitzero"`

	// One of "annual", "unpaid", "sick", "unpaid_sick", "maternity", "paternity", "special",
	// "day_off", "compensatory", "personal", "emergency", "unexcused_absence".
	LeaveType string `json:"leave_type"`
}

// SchedulesCreateResponse is part of what POST /company/v3/schedules answers.
type SchedulesCreateResponse struct {
	Data []SchedulesCreateRow `json:"data"`
}

// SchedulesCreateRow is part of what POST /company/v3/schedules answers.
type SchedulesCreateRow struct {
	ID           int64                     `json:"id"`
	Type         string                    `json:"type"`
	LeaveType    *string                   `json:"leave_type"`
	IsSplit      bool                      `json:"is_split"`
	Dates        []string                  `json:"dates"`
	Timezone     *string                   `json:"timezone"`
	Start        *string                   `json:"start"`
	End          *string                   `json:"end"`
	TimePlanned  int64                     `json:"time_planned"`
	BreakTime    int64                     `json:"break_time"`
	GraceStart   int64                     `json:"grace_start"`
	GraceEnd     int64                     `json:"grace_end"`
	LocationID   *int64                    `json:"location_id"`
	DepartmentID *int64                    `json:"department_id"`
	PositionID   *int64                    `json:"position_id"`
	Shifts       []SchedulesCreateRowShift `json:"shifts"`
	Users        []SchedulesCreateRowUser  `json:"users"`
}

// SchedulesCreateRowShift is part of what POST /company/v3/schedules answers.
type SchedulesCreateRowShift struct {
	ID           int64  `json:"id"`
	Start        string `json:"start"`
	End          string `json:"end"`
	TimePlanned  int64  `json:"time_planned"`
	LocationID   *int64 `json:"location_id"`
	DepartmentID *int64 `json:"department_id"`
	PositionID   *int64 `json:"position_id"`
}

// SchedulesCreateRowUser is part of what POST /company/v3/schedules answers.
type SchedulesCreateRowUser struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
}

// SchedulesDeleteResponse is part of what DELETE /company/v3/schedules/{id} answers.
type SchedulesDeleteResponse struct {
	Data DeleteOutcome `json:"data"`
}

// SchedulesGetResponse is part of what GET /company/v3/schedules/{id} answers.
type SchedulesGetResponse struct {
	Data SchedulesGetData `json:"data"`
}

// SchedulesGetData is part of what GET /company/v3/schedules/{id} answers.
type SchedulesGetData struct {
	ID           int64                   `json:"id"`
	Type         string                  `json:"type"`
	LeaveType    *string                 `json:"leave_type"`
	IsSplit      bool                    `json:"is_split"`
	Dates        []string                `json:"dates"`
	Timezone     string                  `json:"timezone"`
	Start        string                  `json:"start"`
	End          string                  `json:"end"`
	TimePlanned  int64                   `json:"time_planned"`
	BreakTime    int64                   `json:"break_time"`
	GraceStart   int64                   `json:"grace_start"`
	GraceEnd     int64                   `json:"grace_end"`
	LocationID   *int64                  `json:"location_id"`
	DepartmentID *int64                  `json:"department_id"`
	PositionID   *int64                  `json:"position_id"`
	Shifts       []SchedulesGetDataShift `json:"shifts"`
	Users        []SchedulesGetDataUser  `json:"users"`
}

// SchedulesGetDataShift is part of what GET /company/v3/schedules/{id} answers.
type SchedulesGetDataShift struct {
	ID           int64  `json:"id"`
	Start        string `json:"start"`
	End          string `json:"end"`
	TimePlanned  int64  `json:"time_planned"`
	LocationID   *int64 `json:"location_id"`
	DepartmentID *int64 `json:"department_id"`
	PositionID   *int64 `json:"position_id"`
}

// SchedulesGetDataUser is part of what GET /company/v3/schedules/{id} answers.
type SchedulesGetDataUser struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
}

// TasksGetParams is the query of GET /company/v3/tasks/{id}.
type TasksGetParams struct {
	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "items", "managers", "user", "author".
	Include []string `json:"include,omitzero"`
}

// TasksGetResponse is part of what GET /company/v3/tasks/{id} answers.
type TasksGetResponse struct {
	Data TasksGetData `json:"data"`
}

// TasksGetData is part of what GET /company/v3/tasks/{id} answers.
type TasksGetData struct {
	ID           int64              `json:"id"`
	ExternalID   *string            `json:"external_id"`
	Title        string             `json:"title"`
	Description  *string            `json:"description"`
	Status       string             `json:"status"`
	Active       bool               `json:"active"`
	Priority     int64              `json:"priority"`
	UserID       int64              `json:"user_id"`
	AuthorID     int64              `json:"author_id"`
	CategoryID   *int64             `json:"category_id"`
	LocationID   *int64             `json:"location_id"`
	DepartmentID *int64             `json:"department_id"`
	PositionID   *int64             `json:"position_id"`
	DueDate      string             `json:"due_date"`
	TimeStart    string             `json:"time_start"`
	TimeEnd      string             `json:"time_end"`
	Timezone     string             `json:"timezone"`
	KpiPlan      int64              `json:"kpi_plan"`
	KpiFact      int64              `json:"kpi_fact"`
	TimeWorked   int64              `json:"time_worked"`
	StartedAt    *string            `json:"started_at"`
	FinishedAt   *string            `json:"finished_at"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	Items        []TasksGetDataItem `json:"items"`
	Managers     []EmployeeShort    `json:"managers"`
	User         *EmployeeShort     `json:"user"`
	Author       *EmployeeShort     `json:"author"`
}

// TasksGetDataItem is part of what GET /company/v3/tasks/{id} answers.
type TasksGetDataItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Order       int64  `json:"order"`
	IsCompleted bool   `json:"is_completed"`
}

// TasksListParams is the query of GET /company/v3/tasks.
type TasksListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Only rows carrying these keys of yours. The other half of an upsert: write with your key, read
	// back with it.
	ExternalIDs []string `json:"external_ids,omitzero"`

	// Only rows belonging to these people, by id.
	Users []int64 `json:"users,omitzero"`

	// Only rows in these categories.
	Categories []int64 `json:"categories,omitzero"`

	// Only rows in these states.
	// One of "created", "started", "paused", "completed", "incompleted", "pastdue".
	Statuses []string `json:"statuses,omitzero"`

	// Only rows switched on (`true`) or off (`false`). Omit for both.
	Active Opt[bool] `json:"active,omitzero"`

	// Free text over the names the section lists.
	Search Opt[string] `json:"search,omitzero"`

	// Due at or after this date (YYYY-MM-DD).
	DueFrom Opt[string] `json:"due_from,omitzero"`

	// Due at or before this date (YYYY-MM-DD).
	DueTo Opt[string] `json:"due_to,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "items", "managers", "user", "author".
	Include []string `json:"include,omitzero"`
}

// TasksListResponse is part of what GET /company/v3/tasks answers.
type TasksListResponse struct {
	Data []TasksListRow `json:"data"`
	Meta TasksListMeta  `json:"meta"`
}

// TasksListRow is part of what GET /company/v3/tasks answers.
type TasksListRow struct {
	ID           int64              `json:"id"`
	ExternalID   *string            `json:"external_id"`
	Title        string             `json:"title"`
	Description  *string            `json:"description"`
	Status       string             `json:"status"`
	Active       bool               `json:"active"`
	Priority     int64              `json:"priority"`
	UserID       int64              `json:"user_id"`
	AuthorID     int64              `json:"author_id"`
	CategoryID   *int64             `json:"category_id"`
	LocationID   *int64             `json:"location_id"`
	DepartmentID *int64             `json:"department_id"`
	PositionID   *int64             `json:"position_id"`
	DueDate      string             `json:"due_date"`
	TimeStart    string             `json:"time_start"`
	TimeEnd      string             `json:"time_end"`
	Timezone     string             `json:"timezone"`
	KpiPlan      int64              `json:"kpi_plan"`
	KpiFact      int64              `json:"kpi_fact"`
	TimeWorked   int64              `json:"time_worked"`
	StartedAt    *string            `json:"started_at"`
	FinishedAt   *string            `json:"finished_at"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	Items        []TasksListRowItem `json:"items"`
	Managers     []EmployeeShort    `json:"managers"`
	User         *EmployeeShort     `json:"user"`
	Author       *EmployeeShort     `json:"author"`
}

// TasksListRowItem is part of what GET /company/v3/tasks answers.
type TasksListRowItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Order       int64  `json:"order"`
	IsCompleted bool   `json:"is_completed"`
}

// TasksListMeta is part of what GET /company/v3/tasks answers.
type TasksListMeta struct {
	PerPage    int64   `json:"per_page"`
	NextCursor *string `json:"next_cursor"`
	PrevCursor *string `json:"prev_cursor"`
}

// TasksUpsertBody is part of what POST /company/v3/tasks/upsert takes.
type TasksUpsertBody struct {
	Tasks []TasksUpsertTask `json:"tasks"`
}

// TasksUpsertTask is part of what POST /company/v3/tasks/upsert takes.
type TasksUpsertTask struct {
	ExternalID   string      `json:"external_id"`
	Title        string      `json:"title"`
	Description  Opt[string] `json:"description,omitzero"`
	UserID       int64       `json:"user_id"`
	CategoryID   Opt[int64]  `json:"category_id,omitzero"`
	LocationID   Opt[int64]  `json:"location_id,omitzero"`
	DepartmentID Opt[int64]  `json:"department_id,omitzero"`
	PositionID   Opt[int64]  `json:"position_id,omitzero"`

	// A plain date, YYYY-MM-DD.
	DueDate Opt[string] `json:"due_date,omitzero"`

	// An instant, ISO 8601 with an offset.
	TimeStart Opt[string] `json:"time_start,omitzero"`

	// An instant, ISO 8601 with an offset.
	TimeEnd  Opt[string]           `json:"time_end,omitzero"`
	Timezone Opt[string]           `json:"timezone,omitzero"`
	Priority Opt[int64]            `json:"priority,omitzero"`
	Active   Opt[bool]             `json:"active,omitzero"`
	KpiPlan  Opt[float64]          `json:"kpi_plan,omitzero"`
	Managers []int64               `json:"managers,omitzero"`
	Items    []TasksUpsertTaskItem `json:"items,omitzero"`
}

// TasksUpsertTaskItem is part of what POST /company/v3/tasks/upsert takes.
type TasksUpsertTaskItem struct {
	Title string     `json:"title"`
	Order Opt[int64] `json:"order,omitzero"`
}

// TasksUpsertResponse is part of what POST /company/v3/tasks/upsert answers.
type TasksUpsertResponse struct {
	Data []UpsertOutcome `json:"data"`
}

// TimesheetsListParams is the query of GET /company/v3/timesheets.
type TimesheetsListParams struct {
	// Start of the window, inclusive (YYYY-MM-DD).
	DateFrom string `json:"date_from"`

	// End of the window, inclusive (YYYY-MM-DD).
	DateTo string `json:"date_to"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Only rows belonging to these people, by id.
	Users []int64 `json:"users,omitzero"`

	// Only these locations, by id.
	Locations []int64 `json:"locations,omitzero"`

	// Only these departments, by id.
	Departments []int64 `json:"departments,omitzero"`

	// Only these positions, by id.
	Positions []int64 `json:"positions,omitzero"`

	// Only people on these employment terms.
	Employment Opt[string] `json:"employment,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "actual", "variance", "user", "location", "department", "position".
	Include []string `json:"include,omitzero"`
}

// TimesheetsListResponse is part of what GET /company/v3/timesheets answers.
type TimesheetsListResponse struct {
	Data []TimesheetsListRow `json:"data"`
	Meta TimesheetsListMeta  `json:"meta"`
}

// TimesheetsListRow is part of what GET /company/v3/timesheets answers.
type TimesheetsListRow struct {
	Date     string                     `json:"date"`
	User     TimesheetsListRowUser      `json:"user"`
	Planned  *TimesheetsListRowPlanned  `json:"planned"`
	Actual   *TimesheetsListRowActual   `json:"actual"`
	Variance *TimesheetsListRowVariance `json:"variance"`
}

// TimesheetsListRowUser is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowUser struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Code       *string `json:"code"`
	FirstName  *string `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName   *string `json:"last_name"`
}

// TimesheetsListRowPlanned is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlanned struct {
	ScheduleID   int64                               `json:"schedule_id"`
	Type         string                              `json:"type"`
	LeaveType    *string                             `json:"leave_type"`
	IsSplit      bool                                `json:"is_split"`
	Timezone     string                              `json:"timezone"`
	Start        string                              `json:"start"`
	End          string                              `json:"end"`
	TimePlanned  int64                               `json:"time_planned"`
	BreakTime    int64                               `json:"break_time"`
	GraceStart   int64                               `json:"grace_start"`
	GraceEnd     int64                               `json:"grace_end"`
	Shifts       []TimesheetsListRowPlannedShift     `json:"shifts"`
	LocationID   *int64                              `json:"location_id"`
	DepartmentID *int64                              `json:"department_id"`
	PositionID   *int64                              `json:"position_id"`
	Location     *TimesheetsListRowPlannedLocation   `json:"location"`
	Department   *TimesheetsListRowPlannedDepartment `json:"department"`
	Position     *TimesheetsListRowPlannedPosition   `json:"position"`
}

// TimesheetsListRowPlannedShift is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlannedShift struct {
	ID           int64                                    `json:"id"`
	Code         *string                                  `json:"code"`
	Start        string                                   `json:"start"`
	End          string                                   `json:"end"`
	TimePlanned  int64                                    `json:"time_planned"`
	LocationID   *int64                                   `json:"location_id"`
	DepartmentID *int64                                   `json:"department_id"`
	PositionID   *int64                                   `json:"position_id"`
	Location     *TimesheetsListRowPlannedShiftLocation   `json:"location"`
	Department   *TimesheetsListRowPlannedShiftDepartment `json:"department"`
	Position     *TimesheetsListRowPlannedShiftPosition   `json:"position"`
}

// TimesheetsListRowPlannedShiftLocation is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlannedShiftLocation struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Title      string  `json:"title"`
}

// TimesheetsListRowPlannedShiftDepartment is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlannedShiftDepartment struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Title      string  `json:"title"`
}

// TimesheetsListRowPlannedShiftPosition is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlannedShiftPosition struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Title      string  `json:"title"`
}

// TimesheetsListRowPlannedLocation is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlannedLocation struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Title      string  `json:"title"`
}

// TimesheetsListRowPlannedDepartment is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlannedDepartment struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Title      string  `json:"title"`
}

// TimesheetsListRowPlannedPosition is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowPlannedPosition struct {
	ID         int64   `json:"id"`
	ExternalID *string `json:"external_id"`
	Title      string  `json:"title"`
}

// TimesheetsListRowActual is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowActual struct {
	In               *string                        `json:"in"`
	Out              *string                        `json:"out"`
	TimeWorked       int64                          `json:"time_worked"`
	TimeBreak        int64                          `json:"time_break"`
	TimeWorkedDayOff int64                          `json:"time_worked_day_off"`
	TimeNight        int64                          `json:"time_night"`
	Shifts           []TimesheetsListRowActualShift `json:"shifts"`
}

// TimesheetsListRowActualShift is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowActualShift struct {
	ID         int64  `json:"id"`
	In         string `json:"in"`
	Out        string `json:"out"`
	TimeWorked int64  `json:"time_worked"`
}

// TimesheetsListRowVariance is part of what GET /company/v3/timesheets answers.
type TimesheetsListRowVariance struct {
	TimeLate        int64 `json:"time_late"`
	TimeEarlyLeft   int64 `json:"time_early_left"`
	TimeOverworked  int64 `json:"time_overworked"`
	TimeUnderworked int64 `json:"time_underworked"`
}

// TimesheetsListMeta is part of what GET /company/v3/timesheets answers.
type TimesheetsListMeta struct {
	UsersPerPage int64   `json:"users_per_page"`
	NextCursor   *string `json:"next_cursor"`
	PrevCursor   *string `json:"prev_cursor"`
}

// UserFiltersDeleteResponse is part of what DELETE /company/v3/user-filters/{id} answers.
type UserFiltersDeleteResponse struct {
	Data DeleteOutcome `json:"data"`
}

// UserFiltersGetParams is the query of GET /company/v3/user-filters/{id}.
type UserFiltersGetParams struct {
	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "managers".
	Include []string `json:"include,omitzero"`
}

// UserFiltersGetResponse is part of what GET /company/v3/user-filters/{id} answers.
type UserFiltersGetResponse struct {
	Data UserFiltersGetData `json:"data"`
}

// UserFiltersGetData is part of what GET /company/v3/user-filters/{id} answers.
type UserFiltersGetData struct {
	ID          int64           `json:"id"`
	ExternalID  *string         `json:"external_id"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Managers    []EmployeeShort `json:"managers"`
}

// UserFiltersListParams is the query of GET /company/v3/user-filters.
type UserFiltersListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Free text over the names the section lists.
	Search Opt[string] `json:"search,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "managers".
	Include []string `json:"include,omitzero"`
}

// UserFiltersListResponse is part of what GET /company/v3/user-filters answers.
type UserFiltersListResponse struct {
	Data  []UserFiltersListRow `json:"data"`
	Links PageLinks            `json:"links"`
	Meta  PageMeta             `json:"meta"`
}

// UserFiltersListRow is part of what GET /company/v3/user-filters answers.
type UserFiltersListRow struct {
	ID          int64           `json:"id"`
	ExternalID  *string         `json:"external_id"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Managers    []EmployeeShort `json:"managers"`
}

// UserFiltersUpsertBody is part of what POST /company/v3/user-filters/upsert takes.
type UserFiltersUpsertBody struct {
	Items []UserFiltersUpsertItem `json:"items"`
}

// UserFiltersUpsertItem is part of what POST /company/v3/user-filters/upsert takes.
type UserFiltersUpsertItem struct {
	ExternalID  string      `json:"external_id"`
	Title       string      `json:"title"`
	Description Opt[string] `json:"description,omitzero"`
}

// UserFiltersUpsertResponse is part of what POST /company/v3/user-filters/upsert answers.
type UserFiltersUpsertResponse struct {
	Data []UpsertOutcome `json:"data"`
}

// UserRequestsGetResponse is part of what GET /company/v3/user-requests/{id} answers.
type UserRequestsGetResponse struct {
	Data UserRequestsGetData `json:"data"`
}

// UserRequestsGetData is part of what GET /company/v3/user-requests/{id} answers.
type UserRequestsGetData struct {
	ID        string                     `json:"id"`
	Type      string                     `json:"type"`
	Subtype   *string                    `json:"subtype"`
	Status    string                     `json:"status"`
	UserID    int64                      `json:"user_id"`
	AuthorID  int64                      `json:"author_id"`
	Period    UserRequestsGetDataPeriod  `json:"period"`
	Comment   *string                    `json:"comment"`
	Amount    *float64                   `json:"amount"`
	Currency  *string                    `json:"currency"`
	CreatedAt string                     `json:"created_at"`
	UpdatedAt string                     `json:"updated_at"`
	Content   UserRequestsGetDataContent `json:"content"`
}

// UserRequestsGetDataPeriod is part of what GET /company/v3/user-requests/{id} answers.
type UserRequestsGetDataPeriod struct {
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
}

// UserRequestsGetDataContent is part of what GET /company/v3/user-requests/{id} answers.
type UserRequestsGetDataContent struct {
	Type     string                              `json:"type"`
	Clockins []UserRequestsGetDataContentClockin `json:"clockins"`
}

// UserRequestsGetDataContentClockin is part of what GET /company/v3/user-requests/{id} answers.
type UserRequestsGetDataContentClockin struct {
	Status   string `json:"status"`
	Datetime string `json:"datetime"`
}

// UserRequestsListParams is the query of GET /company/v3/user-requests.
type UserRequestsListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Only rows of these types.
	// One of "leave", "work", "general", "finance".
	Types []string `json:"types,omitzero"`

	// Only rows in these states.
	// One of "pending", "accepted", "rejected", "cancelled", "approval", "execution", "signing".
	Statuses []string `json:"statuses,omitzero"`

	// Only rows of these subtypes, which narrow a type further.
	Subtypes []string `json:"subtypes,omitzero"`

	// Only rows belonging to these people, by id.
	Users []int64 `json:"users,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "content", "user", "author".
	Include []string `json:"include,omitzero"`
}

// UserRequestsListResponse is part of what GET /company/v3/user-requests answers.
type UserRequestsListResponse struct {
	Data  []UserRequestsListRow `json:"data"`
	Links PageLinks             `json:"links"`
	Meta  PageMeta              `json:"meta"`
}

// UserRequestsListRow is part of what GET /company/v3/user-requests answers.
type UserRequestsListRow struct {
	ID        string                      `json:"id"`
	Type      string                      `json:"type"`
	Subtype   *string                     `json:"subtype"`
	Status    string                      `json:"status"`
	UserID    int64                       `json:"user_id"`
	AuthorID  int64                       `json:"author_id"`
	Period    UserRequestsListRowPeriod   `json:"period"`
	Comment   *string                     `json:"comment"`
	Amount    *float64                    `json:"amount"`
	Currency  *string                     `json:"currency"`
	CreatedAt string                      `json:"created_at"`
	UpdatedAt string                      `json:"updated_at"`
	User      *EmployeeShort              `json:"user"`
	Author    *EmployeeShort              `json:"author"`
	Content   *UserRequestsListRowContent `json:"content"`
}

// UserRequestsListRowPeriod is part of what GET /company/v3/user-requests answers.
type UserRequestsListRowPeriod struct {
	DateStart *string `json:"date_start"`
	DateEnd   *string `json:"date_end"`
}

// UserRequestsListRowContent is part of what GET /company/v3/user-requests answers.
type UserRequestsListRowContent struct {
	Type       string                              `json:"type"`
	Clockins   []UserRequestsListRowContentClockin `json:"clockins"`
	Amount     *float64                            `json:"amount"`
	CurrencyID *int64                              `json:"currency_id"`
}

// UserRequestsListRowContentClockin is part of what GET /company/v3/user-requests answers.
type UserRequestsListRowContentClockin struct {
	Status   string `json:"status"`
	Datetime string `json:"datetime"`
}

// UsersDismissBody is part of what POST /company/v3/users/dismiss takes.
type UsersDismissBody struct {
	Users []UsersDismissUser `json:"users"`
}

// UsersDismissUser is part of what POST /company/v3/users/dismiss takes.
type UsersDismissUser struct {
	ExternalID Opt[string] `json:"external_id,omitzero"`
	ID         Opt[int64]  `json:"id,omitzero"`
}

// UsersDismissResponse is part of what POST /company/v3/users/dismiss answers.
type UsersDismissResponse struct {
	Data []UpsertOutcome `json:"data"`
}

// UsersGetParams is the query of GET /company/v3/users/{id}.
type UsersGetParams struct {
	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "location", "locations", "department", "position", "user_filters", "dismissal", "meta".
	Include []string `json:"include,omitzero"`
}

// UsersGetResponse is part of what GET /company/v3/users/{id} answers.
type UsersGetResponse struct {
	Data UsersGetData `json:"data"`
}

// UsersGetData is part of what GET /company/v3/users/{id} answers.
type UsersGetData struct {
	ID          int64                    `json:"id"`
	ExternalID  *string                  `json:"external_id"`
	Code        *string                  `json:"code"`
	FirstName   string                   `json:"first_name"`
	MiddleName  *string                  `json:"middle_name"`
	LastName    string                   `json:"last_name"`
	Email       string                   `json:"email"`
	Phone       *string                  `json:"phone"`
	ExtraPhone  *string                  `json:"extra_phone"`
	Role        *string                  `json:"role"`
	Gender      string                   `json:"gender"`
	NationalID  string                   `json:"national_id"`
	TaxID       *string                  `json:"tax_id"`
	InsuranceID *string                  `json:"insurance_id"`
	Employment  string                   `json:"employment"`
	Locale      string                   `json:"locale"`
	Timezone    string                   `json:"timezone"`
	DateBirth   string                   `json:"date_birth"`
	DateHire    string                   `json:"date_hire"`
	DateLeave   *string                  `json:"date_leave"`
	Photo       *string                  `json:"photo"`
	CreatedAt   string                   `json:"created_at"`
	UpdatedAt   string                   `json:"updated_at"`
	DismissedAt *string                  `json:"dismissed_at"`
	Location    *UsersGetDataLocation    `json:"location"`
	Locations   []UsersGetDataLocation   `json:"locations"`
	Department  *UsersGetDataDepartment  `json:"department"`
	Position    *UsersGetDataPosition    `json:"position"`
	UserFilters []UsersGetDataUserFilter `json:"user_filters"`
	Dismissal   *UsersGetDataDismissal   `json:"dismissal"`
	Meta        *UsersGetDataMeta        `json:"meta"`
}

// UsersGetDataLocation is part of what GET /company/v3/users/{id} answers.
type UsersGetDataLocation struct {
	ID          int64    `json:"id"`
	ExternalID  *string  `json:"external_id"`
	Code        *string  `json:"code"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	Radius      int64    `json:"radius"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

// UsersGetDataDepartment is part of what GET /company/v3/users/{id} answers.
type UsersGetDataDepartment struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UsersGetDataPosition is part of what GET /company/v3/users/{id} answers.
type UsersGetDataPosition struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UsersGetDataUserFilter is part of what GET /company/v3/users/{id} answers.
type UsersGetDataUserFilter struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UsersGetDataDismissal is part of what GET /company/v3/users/{id} answers.
type UsersGetDataDismissal struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// UsersGetDataMeta is part of what GET /company/v3/users/{id} answers.
type UsersGetDataMeta struct {
	BirthPlace           string   `json:"birth_place"`
	MaritalStatus        string   `json:"marital_status"`
	Religion             *string  `json:"religion"`
	BloodType            *string  `json:"blood_type"`
	Children             *int64   `json:"children"`
	ContactName          *string  `json:"contact_name"`
	Relationship         *string  `json:"relationship"`
	Phone                *string  `json:"phone"`
	DomicileAddress      *string  `json:"domicile_address"`
	DomicileCity         *string  `json:"domicile_city"`
	DomicileProvince     *string  `json:"domicile_province"`
	DomicileDistrict     *string  `json:"domicile_district"`
	DomicilePostalCode   *string  `json:"domicile_postal_code"`
	DocumentAddress      *string  `json:"document_address"`
	DocumentCity         *string  `json:"document_city"`
	DocumentProvince     *string  `json:"document_province"`
	DocumentDistrict     *string  `json:"document_district"`
	DocumentPostalCode   *string  `json:"document_postal_code"`
	EducationLevel       *string  `json:"education_level"`
	EducationInstitution *string  `json:"education_institution"`
	EducationMajor       *string  `json:"education_major"`
	GraduationYear       *int64   `json:"graduation_year"`
	EducationGpa         *float64 `json:"education_gpa"`
	FullName             *string  `json:"full_name"`
	AliasName            *string  `json:"alias_name"`
	LocalName            *string  `json:"local_name"`
	Nationality          *string  `json:"nationality"`
	MarriageDate         *string  `json:"marriage_date"`
	RetireAge            *int64   `json:"retire_age"`
	RetireDate           *string  `json:"retire_date"`
	EthnicOrigin         *string  `json:"ethnic_origin"`
	ContractNumber       *string  `json:"contract_number"`
	HealthInsuranceID    *string  `json:"health_insurance_id"`
	GovSavingsID         *string  `json:"gov_savings_id"`
	GovSavingsAcc        *string  `json:"gov_savings_acc"`
}

// UsersListParams is the query of GET /company/v3/users.
type UsersListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Free text over the names the section lists.
	Search Opt[string] `json:"search,omitzero"`

	// Only rows changed at or after this instant (ISO 8601). The cheap way to sync: ask for what
	// moved, not for everything.
	UpdatedSince Opt[string] `json:"updated_since,omitzero"`

	// Whether the people who left are in the answer: `active`, `dismissed`, or `all`.
	// One of "active", "dismissed", "all".
	Status Opt[string] `json:"status,omitzero"`

	// Only these ids.
	IDs []int64 `json:"ids,omitzero"`

	// Only these employee codes.
	Codes []string `json:"codes,omitzero"`

	// Only rows carrying these keys of yours. The other half of an upsert: write with your key, read
	// back with it.
	ExternalIDs []string `json:"external_ids,omitzero"`

	// Only these locations, by id.
	Locations []int64 `json:"locations,omitzero"`

	// Only these departments, by id.
	Departments []int64 `json:"departments,omitzero"`

	// Only these positions, by id.
	Positions []int64 `json:"positions,omitzero"`

	// Only these user filters, by id.
	UserFilters []int64 `json:"user_filters,omitzero"`

	// Only people on these employment terms.
	Employment []string `json:"employment,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "location", "locations", "department", "position", "user_filters", "dismissal", "meta".
	Include []string `json:"include,omitzero"`
}

// UsersListResponse is part of what GET /company/v3/users answers.
type UsersListResponse struct {
	Data  []UsersListRow `json:"data"`
	Links PageLinks      `json:"links"`
	Meta  PageMeta       `json:"meta"`
}

// UsersListRow is part of what GET /company/v3/users answers.
type UsersListRow struct {
	ID          int64                    `json:"id"`
	ExternalID  *string                  `json:"external_id"`
	Code        *string                  `json:"code"`
	FirstName   string                   `json:"first_name"`
	MiddleName  *string                  `json:"middle_name"`
	LastName    string                   `json:"last_name"`
	Email       string                   `json:"email"`
	Phone       *string                  `json:"phone"`
	ExtraPhone  *string                  `json:"extra_phone"`
	Role        *string                  `json:"role"`
	Gender      string                   `json:"gender"`
	NationalID  string                   `json:"national_id"`
	TaxID       *string                  `json:"tax_id"`
	InsuranceID *string                  `json:"insurance_id"`
	Employment  string                   `json:"employment"`
	Locale      string                   `json:"locale"`
	Timezone    string                   `json:"timezone"`
	DateBirth   string                   `json:"date_birth"`
	DateHire    string                   `json:"date_hire"`
	DateLeave   *string                  `json:"date_leave"`
	Photo       *string                  `json:"photo"`
	CreatedAt   string                   `json:"created_at"`
	UpdatedAt   string                   `json:"updated_at"`
	DismissedAt *string                  `json:"dismissed_at"`
	Dismissal   *UsersListRowDismissal   `json:"dismissal"`
	Location    *UsersListRowLocation    `json:"location"`
	Locations   []UsersListRowLocation   `json:"locations"`
	Department  *UsersListRowDepartment  `json:"department"`
	Position    *UsersListRowPosition    `json:"position"`
	UserFilters []UsersListRowUserFilter `json:"user_filters"`
	Meta        *UsersListRowMeta        `json:"meta"`
}

// UsersListRowDismissal is part of what GET /company/v3/users answers.
type UsersListRowDismissal struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// UsersListRowLocation is part of what GET /company/v3/users answers.
type UsersListRowLocation struct {
	ID          int64    `json:"id"`
	ExternalID  *string  `json:"external_id"`
	Code        *string  `json:"code"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	Radius      int64    `json:"radius"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

// UsersListRowDepartment is part of what GET /company/v3/users answers.
type UsersListRowDepartment struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UsersListRowPosition is part of what GET /company/v3/users answers.
type UsersListRowPosition struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UsersListRowUserFilter is part of what GET /company/v3/users answers.
type UsersListRowUserFilter struct {
	ID          int64   `json:"id"`
	ExternalID  *string `json:"external_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UsersListRowMeta is part of what GET /company/v3/users answers.
type UsersListRowMeta struct {
	BirthPlace           string   `json:"birth_place"`
	MaritalStatus        string   `json:"marital_status"`
	Religion             *string  `json:"religion"`
	BloodType            *string  `json:"blood_type"`
	Children             *int64   `json:"children"`
	ContactName          *string  `json:"contact_name"`
	Relationship         *string  `json:"relationship"`
	Phone                *string  `json:"phone"`
	DomicileAddress      *string  `json:"domicile_address"`
	DomicileCity         *string  `json:"domicile_city"`
	DomicileProvince     *string  `json:"domicile_province"`
	DomicileDistrict     *string  `json:"domicile_district"`
	DomicilePostalCode   *string  `json:"domicile_postal_code"`
	DocumentAddress      *string  `json:"document_address"`
	DocumentCity         *string  `json:"document_city"`
	DocumentProvince     *string  `json:"document_province"`
	DocumentDistrict     *string  `json:"document_district"`
	DocumentPostalCode   *string  `json:"document_postal_code"`
	EducationLevel       *string  `json:"education_level"`
	EducationInstitution *string  `json:"education_institution"`
	EducationMajor       *string  `json:"education_major"`
	GraduationYear       *int64   `json:"graduation_year"`
	EducationGpa         *float64 `json:"education_gpa"`
	FullName             *string  `json:"full_name"`
	AliasName            *string  `json:"alias_name"`
	LocalName            *string  `json:"local_name"`
	Nationality          *string  `json:"nationality"`
	MarriageDate         *string  `json:"marriage_date"`
	RetireAge            *int64   `json:"retire_age"`
	RetireDate           *string  `json:"retire_date"`
	EthnicOrigin         *string  `json:"ethnic_origin"`
	ContractNumber       *string  `json:"contract_number"`
	HealthInsuranceID    *string  `json:"health_insurance_id"`
	GovSavingsID         *string  `json:"gov_savings_id"`
	GovSavingsAcc        *string  `json:"gov_savings_acc"`
}

// UsersUpsertBody is part of what POST /company/v3/users/upsert takes.
type UsersUpsertBody struct {
	Users []UsersUpsertUser `json:"users"`
}

// UsersUpsertUser is part of what POST /company/v3/users/upsert takes.
type UsersUpsertUser struct {
	ExternalID Opt[string] `json:"external_id,omitzero"`
	FirstName  string      `json:"first_name"`
	MiddleName Opt[string] `json:"middle_name,omitzero"`
	LastName   Opt[string] `json:"last_name,omitzero"`
	Code       Opt[string] `json:"code,omitzero"`
	Email      Opt[string] `json:"email,omitzero"`
	Phone      Opt[string] `json:"phone,omitzero"`
	ExtraPhone Opt[string] `json:"extra_phone,omitzero"`

	// One of "admin", "employee".
	Role string `json:"role"`

	// One of "male", "female", "other".
	Gender Opt[string] `json:"gender,omitzero"`

	// One of "en", "ru", "kk", "uk", "id", "uz", "az", "fr", "vi", "zh".
	Locale   Opt[string] `json:"locale,omitzero"`
	Timezone Opt[string] `json:"timezone,omitzero"`

	// A plain date, YYYY-MM-DD.
	DateHire Opt[string] `json:"date_hire,omitzero"`

	// A plain date, YYYY-MM-DD.
	DateLeave Opt[string] `json:"date_leave,omitzero"`

	// A plain date, YYYY-MM-DD.
	DateBirth   Opt[string] `json:"date_birth,omitzero"`
	NationalID  Opt[string] `json:"national_id,omitzero"`
	TaxID       Opt[string] `json:"tax_id,omitzero"`
	InsuranceID Opt[string] `json:"insurance_id,omitzero"`

	// One of "full_time", "part_time", "irregular_hours", "contract_1", "contract_2",
	// "apprenticeship", "traineeship", "piece_rate", "probation", "outstaffing".
	Employment     Opt[string] `json:"employment,omitzero"`
	Responsibility Opt[string] `json:"responsibility,omitzero"`
	LocationID     int64       `json:"location_id"`
	Locations      []int64     `json:"locations,omitzero"`
	DepartmentID   Opt[int64]  `json:"department_id,omitzero"`
	PositionID     Opt[int64]  `json:"position_id,omitzero"`
	UserFilters    []int64     `json:"user_filters,omitzero"`
}

// UsersUpsertResponse is part of what POST /company/v3/users/upsert answers.
type UsersUpsertResponse struct {
	Data []UpsertOutcome `json:"data"`
}

// WebhooksCreateBody is part of what POST /company/v3/webhooks takes.
type WebhooksCreateBody struct {
	Title        Opt[string] `json:"title,omitzero"`
	URL          string      `json:"url"`
	ContactEmail Opt[string] `json:"contact_email,omitzero"`

	// One of "user.created", "user.updated", "user.deleted", "user.restored", "user.purged",
	// "location.created", "location.updated", "location.deleted", "department.created",
	// "department.updated", "department.deleted", "position.created", "position.updated",
	// "position.deleted", "task.created", "task.completed", "task.approved", "task.rejected",
	// "task.deleted".
	Events    []string                 `json:"events"`
	AuthBasic *WebhooksCreateAuthBasic `json:"auth_basic,omitzero"`
	AuthToken Opt[string]              `json:"auth_token,omitzero"`
	Active    bool                     `json:"active"`
}

// WebhooksCreateAuthBasic is part of what POST /company/v3/webhooks takes.
type WebhooksCreateAuthBasic struct {
	Username Opt[string] `json:"username,omitzero"`
	Password Opt[string] `json:"password,omitzero"`
}

// WebhooksCreateResponse is part of what POST /company/v3/webhooks answers.
type WebhooksCreateResponse struct {
	Data WebhooksCreateData `json:"data"`
}

// WebhooksCreateData is part of what POST /company/v3/webhooks answers.
type WebhooksCreateData struct {
	ID           int64                    `json:"id"`
	Title        string                   `json:"title"`
	URL          string                   `json:"url"`
	ContactEmail string                   `json:"contact_email"`
	Secret       string                   `json:"secret"`
	Events       []string                 `json:"events"`
	Auth         WebhooksCreateDataAuth   `json:"auth"`
	Active       bool                     `json:"active"`
	Health       WebhooksCreateDataHealth `json:"health"`
}

// WebhooksCreateDataAuth is part of what POST /company/v3/webhooks answers.
type WebhooksCreateDataAuth struct {
	Type     string  `json:"type"`
	Username *string `json:"username"`
}

// WebhooksCreateDataHealth is part of what POST /company/v3/webhooks answers.
type WebhooksCreateDataHealth struct {
	ConsecutiveFailures int64   `json:"consecutive_failures"`
	LastSuccessAt       *string `json:"last_success_at"`
	LastFailureAt       *string `json:"last_failure_at"`
	LastFailureStatus   *int64  `json:"last_failure_status"`
	DisabledAt          *string `json:"disabled_at"`
	DisabledReason      *string `json:"disabled_reason"`
}

// WebhooksDeleteResponse is part of what DELETE /company/v3/webhooks/{id} answers.
type WebhooksDeleteResponse struct {
	Data DeleteOutcome `json:"data"`
}

// WebhooksGetResponse is part of what GET /company/v3/webhooks/{id} answers.
type WebhooksGetResponse struct {
	Data WebhooksGetData `json:"data"`
}

// WebhooksGetData is part of what GET /company/v3/webhooks/{id} answers.
type WebhooksGetData struct {
	ID           int64                 `json:"id"`
	Title        string                `json:"title"`
	URL          string                `json:"url"`
	ContactEmail string                `json:"contact_email"`
	Secret       string                `json:"secret"`
	Events       []string              `json:"events"`
	Auth         WebhooksGetDataAuth   `json:"auth"`
	Active       bool                  `json:"active"`
	Health       WebhooksGetDataHealth `json:"health"`
}

// WebhooksGetDataAuth is part of what GET /company/v3/webhooks/{id} answers.
type WebhooksGetDataAuth struct {
	Type     string  `json:"type"`
	Username *string `json:"username"`
}

// WebhooksGetDataHealth is part of what GET /company/v3/webhooks/{id} answers.
type WebhooksGetDataHealth struct {
	ConsecutiveFailures int64   `json:"consecutive_failures"`
	LastSuccessAt       *string `json:"last_success_at"`
	LastFailureAt       *string `json:"last_failure_at"`
	LastFailureStatus   *int64  `json:"last_failure_status"`
	DisabledAt          *string `json:"disabled_at"`
	DisabledReason      *string `json:"disabled_reason"`
}

// WebhooksListParams is the query of GET /company/v3/webhooks.
type WebhooksListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Only rows switched on (`true`) or off (`false`). Omit for both.
	Active Opt[bool] `json:"active,omitzero"`
}

// WebhooksListResponse is part of what GET /company/v3/webhooks answers.
type WebhooksListResponse struct {
	Data  []WebhooksListRow `json:"data"`
	Links PageLinks         `json:"links"`
	Meta  PageMeta          `json:"meta"`
}

// WebhooksListRow is part of what GET /company/v3/webhooks answers.
type WebhooksListRow struct {
	ID           int64                 `json:"id"`
	Title        string                `json:"title"`
	URL          string                `json:"url"`
	ContactEmail string                `json:"contact_email"`
	Secret       string                `json:"secret"`
	Events       []string              `json:"events"`
	Auth         WebhooksListRowAuth   `json:"auth"`
	Active       bool                  `json:"active"`
	Health       WebhooksListRowHealth `json:"health"`
}

// WebhooksListRowAuth is part of what GET /company/v3/webhooks answers.
type WebhooksListRowAuth struct {
	Type     string  `json:"type"`
	Username *string `json:"username"`
}

// WebhooksListRowHealth is part of what GET /company/v3/webhooks answers.
type WebhooksListRowHealth struct {
	ConsecutiveFailures int64   `json:"consecutive_failures"`
	LastSuccessAt       *string `json:"last_success_at"`
	LastFailureAt       *string `json:"last_failure_at"`
	LastFailureStatus   *int64  `json:"last_failure_status"`
	DisabledAt          *string `json:"disabled_at"`
	DisabledReason      *string `json:"disabled_reason"`
}

// WebhooksRotateSecretResponse is part of what POST /company/v3/webhooks/{id}/secret answers.
type WebhooksRotateSecretResponse struct {
	Data WebhooksRotateSecretData `json:"data"`
}

// WebhooksRotateSecretData is part of what POST /company/v3/webhooks/{id}/secret answers.
type WebhooksRotateSecretData struct {
	ID           int64                          `json:"id"`
	Title        string                         `json:"title"`
	URL          string                         `json:"url"`
	ContactEmail string                         `json:"contact_email"`
	Secret       string                         `json:"secret"`
	Events       []string                       `json:"events"`
	Auth         WebhooksRotateSecretDataAuth   `json:"auth"`
	Active       bool                           `json:"active"`
	Health       WebhooksRotateSecretDataHealth `json:"health"`
}

// WebhooksRotateSecretDataAuth is part of what POST /company/v3/webhooks/{id}/secret answers.
type WebhooksRotateSecretDataAuth struct {
	Type     string  `json:"type"`
	Username *string `json:"username"`
}

// WebhooksRotateSecretDataHealth is part of what POST /company/v3/webhooks/{id}/secret answers.
type WebhooksRotateSecretDataHealth struct {
	ConsecutiveFailures int64   `json:"consecutive_failures"`
	LastSuccessAt       *string `json:"last_success_at"`
	LastFailureAt       *string `json:"last_failure_at"`
	LastFailureStatus   *int64  `json:"last_failure_status"`
	DisabledAt          *string `json:"disabled_at"`
	DisabledReason      *string `json:"disabled_reason"`
}

// WebhooksUpdateBody is part of what PUT /company/v3/webhooks/{id} takes.
type WebhooksUpdateBody struct {
	Title        Opt[string] `json:"title,omitzero"`
	URL          string      `json:"url"`
	ContactEmail Opt[string] `json:"contact_email,omitzero"`

	// One of "user.created", "user.updated", "user.deleted", "user.restored", "user.purged",
	// "location.created", "location.updated", "location.deleted", "department.created",
	// "department.updated", "department.deleted", "position.created", "position.updated",
	// "position.deleted", "task.created", "task.completed", "task.approved", "task.rejected",
	// "task.deleted".
	Events    []string                 `json:"events"`
	AuthBasic *WebhooksUpdateAuthBasic `json:"auth_basic,omitzero"`
	AuthToken Opt[string]              `json:"auth_token,omitzero"`
	Active    bool                     `json:"active"`
}

// WebhooksUpdateAuthBasic is part of what PUT /company/v3/webhooks/{id} takes.
type WebhooksUpdateAuthBasic struct {
	Username Opt[string] `json:"username,omitzero"`
	Password Opt[string] `json:"password,omitzero"`
}

// WebhooksUpdateResponse is part of what PUT /company/v3/webhooks/{id} answers.
type WebhooksUpdateResponse struct {
	Data WebhooksUpdateData `json:"data"`
}

// WebhooksUpdateData is part of what PUT /company/v3/webhooks/{id} answers.
type WebhooksUpdateData struct {
	ID           int64                    `json:"id"`
	Title        string                   `json:"title"`
	URL          string                   `json:"url"`
	ContactEmail string                   `json:"contact_email"`
	Secret       string                   `json:"secret"`
	Events       []string                 `json:"events"`
	Auth         WebhooksUpdateDataAuth   `json:"auth"`
	Active       bool                     `json:"active"`
	Health       WebhooksUpdateDataHealth `json:"health"`
}

// WebhooksUpdateDataAuth is part of what PUT /company/v3/webhooks/{id} answers.
type WebhooksUpdateDataAuth struct {
	Type     string  `json:"type"`
	Username *string `json:"username"`
}

// WebhooksUpdateDataHealth is part of what PUT /company/v3/webhooks/{id} answers.
type WebhooksUpdateDataHealth struct {
	ConsecutiveFailures int64   `json:"consecutive_failures"`
	LastSuccessAt       *string `json:"last_success_at"`
	LastFailureAt       *string `json:"last_failure_at"`
	LastFailureStatus   *int64  `json:"last_failure_status"`
	DisabledAt          *string `json:"disabled_at"`
	DisabledReason      *string `json:"disabled_reason"`
}

// WebhooksDeliveriesGetResponse is part of what GET /company/v3/webhooks/deliveries/{id} answers.
type WebhooksDeliveriesGetResponse struct {
	Data WebhooksDeliveriesGetData `json:"data"`
}

// WebhooksDeliveriesGetData is part of what GET /company/v3/webhooks/deliveries/{id} answers.
type WebhooksDeliveriesGetData struct {
	ID             int64                            `json:"id"`
	WebhookID      int64                            `json:"webhook_id"`
	Event          string                           `json:"event"`
	State          string                           `json:"state"`
	Attempts       int64                            `json:"attempts"`
	ResponseStatus int64                            `json:"response_status"`
	FailureReason  *string                          `json:"failure_reason"`
	OccurredAt     string                           `json:"occurred_at"`
	NextAttemptAt  *string                          `json:"next_attempt_at"`
	ProcessedAt    string                           `json:"processed_at"`
	Payload        WebhooksDeliveriesGetDataPayload `json:"payload"`
}

// WebhooksDeliveriesGetDataPayload is part of what GET /company/v3/webhooks/deliveries/{id} answers.
type WebhooksDeliveriesGetDataPayload struct {
	ID int64 `json:"id"`
}

// WebhooksDeliveriesListParams is the query of GET /company/v3/webhooks/deliveries.
type WebhooksDeliveriesListParams struct {
	// How many rows one page holds. Defaults to 50.
	PerPage Opt[int64] `json:"per_page,omitzero"`

	// The `meta.next_cursor` of the previous page. Omit it for the first. A cursor is bound to the
	// filters it was issued under — change them and start again.
	Cursor Opt[string] `json:"cursor,omitzero"`

	// Only deliveries to these endpoints, by id.
	Webhooks []int64 `json:"webhooks,omitzero"`

	// Only deliveries of these events.
	// One of "user.created", "user.updated", "user.deleted", "user.restored", "user.purged",
	// "location.created", "location.updated", "location.deleted", "department.created",
	// "department.updated", "department.deleted", "position.created", "position.updated",
	// "position.deleted", "task.created", "task.completed", "task.approved", "task.rejected",
	// "task.deleted".
	Events []string `json:"events,omitzero"`

	// Only deliveries that were accepted (`true`) or that were not (`false`).
	Successful Opt[bool] `json:"successful,omitzero"`

	// Only deliveries still waiting on a retry (`true`), or only those finished with (`false`).
	Pending Opt[bool] `json:"pending,omitzero"`

	// Only rows from this instant onward (ISO 8601).
	Since Opt[string] `json:"since,omitzero"`

	// Relations to load, comma-separated. Anything not named is absent from the answer rather than
	// null.
	// One of "payload".
	Include []string `json:"include,omitzero"`
}

// WebhooksDeliveriesListResponse is part of what GET /company/v3/webhooks/deliveries answers.
type WebhooksDeliveriesListResponse struct {
	Data  []WebhooksDeliveriesListRow `json:"data"`
	Links PageLinks                   `json:"links"`
	Meta  PageMeta                    `json:"meta"`
}

// WebhooksDeliveriesListRow is part of what GET /company/v3/webhooks/deliveries answers.
type WebhooksDeliveriesListRow struct {
	ID             int64                             `json:"id"`
	WebhookID      int64                             `json:"webhook_id"`
	Event          string                            `json:"event"`
	State          string                            `json:"state"`
	Attempts       int64                             `json:"attempts"`
	ResponseStatus int64                             `json:"response_status"`
	FailureReason  *string                           `json:"failure_reason"`
	OccurredAt     string                            `json:"occurred_at"`
	NextAttemptAt  *string                           `json:"next_attempt_at"`
	ProcessedAt    *string                           `json:"processed_at"`
	Payload        *WebhooksDeliveriesListRowPayload `json:"payload"`
}

// WebhooksDeliveriesListRowPayload is part of what GET /company/v3/webhooks/deliveries answers.
type WebhooksDeliveriesListRowPayload struct {
	ID int64 `json:"id"`
}

// WebhooksDeliveriesRedeliverResponse is part of what POST /company/v3/webhooks/deliveries/{id}/redeliver answers.
type WebhooksDeliveriesRedeliverResponse struct {
	Data DeleteOutcome `json:"data"`
}

// WebhooksEventsListResponse is part of what GET /company/v3/webhooks/events answers.
type WebhooksEventsListResponse struct {
	Data []string `json:"data"`
}
