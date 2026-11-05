# Changelog

## 0.8.0

### May break your build

Nothing. The value is additive.

### Changed in the API

Reaches you whether or not you update this package.

- Webhooks announce `task.reopened`: a `completed` or `incompleted` task reopened in the app, which
  makes it `started` again. Subscribe to it in `events` of `POST /webhooks` and `PUT /webhooks/{id}`;
  `GET /webhooks/events` lists it, and the `events` filter of `GET /webhooks/deliveries` takes it,
  now up to 20.
- `POST /users/dismiss`: a request still waiting on someone dismissed to approve skips their turn
  and moves on to the next approver. It is cancelled only when they were the last approver left to
  decide, and approvals they already gave stand. Before, every such request was cancelled.
- `DELETE /checklists/{id}`: files on the checklist itself no longer keep it from being deleted,
  whoever uploaded them, and go with it. A file on an item still does.

### New

- `WebhooksEventTaskReopened` in `sets.gen.go`, and the value in `WebhooksEventValues()`.

## 0.7.0

### May break your build

Nothing. Four operations and five sets are new.

### New

- `client.Checklists`: `List`, `ListAll`, `Get`, `Create` and `Delete` over `/checklists` — the
  checklists your company issues, one row shared by everyone a checklist is aimed at. `Create` files
  up to 100 of today's checklists, each once under your `external_id`: a key already on file answers
  `exists`, so a resend after a timeout is safe. `Delete` takes back one nobody has worked on yet;
  after that it answers `409` with `checklist_locked`.
- These calls need `checklists:read` and `checklists:write` on the API key, unless it is `full`.
- `ChecklistsInclude`, `ChecklistsItemType`, `ChecklistsPriority`, `ChecklistsStatus` and
  `ChecklistsType` in `sets.gen.go`, each with its `...Values()`.

## 0.6.0

### May break your build

Nothing. The value is additive.

### Changed in the API

Reaches you whether or not you update this package.

- A single adjustment's `type` also takes `withholding` — in the `types` filter, which now takes up
  to seven, and in `POST /payroll/single-adjustments`. It comes off the take-home pay alone, as a
  loan repayment does, and leaves the company's cost as it is, where `single_deduction_post_tax`
  lowers both.
- `DELETE /documents/{id}` refuses a document sent to a counterparty, by e-mail or a copied signing
  link: `409` with `document_sent_to_counterparty`. It is withdrawn in the app instead.

### New

- `PayrollSingleAdjustmentsTypeWithholding` in `sets.gen.go`, and the value in
  `PayrollSingleAdjustmentsTypeValues()`.

## 0.5.0

### May break your build

Nothing. The seventeen values are additive.

### Changed in the API

Reaches you whether or not you update this package.

- A company may hold up to ten API keys, each with full access or read or write access per
  section, and every operation names the scope it needs. A call outside the key's scopes answers
  `403` with `insufficient_scope`. A key issued before scopes has full access.
- The limit of 100 requests a minute is counted against the company and shared by all of its keys.
- A document's `type` also takes `srts`, `vaccination`, `social_id`, `disability_certificate`,
  `large_family_certificate`, `asp_certificate`, `tech_passport`, `pension`, `rk_passport`,
  `student_card`, `vnzh`, `pcr_certificate`, `lbg_card`, `insurance_policy`, `hunter`, `oralman`
  and `attorney` — in the `types` filter, which now takes up to 49, and in
  `POST /documents/upsert`.
- The API is also served to AI agents as an MCP server at `https://api.clockster.com/company/mcp`;
  the document's introduction says how to connect one.

### New

- A constant per new value in `sets.gen.go`, and the seventeen in `DocumentsTypeValues()`.

## 0.4.0

### May break your build

Nothing. Three operations and one set are new.

### Changed in the API

Reaches you whether or not you update this package.

- `GET /payroll/payslips` answers the `external_id` of a dismissed employee. It was `null` on their
  payslips.

### New

- `client.Payroll.SingleAdjustments`: `List`, `ListAll`, `Create` and `Delete` over
  `/payroll/single-adjustments` — one-off additions and deductions that the next calculation of a
  payslip takes in. `Create` files up to 100 at a time, all or nothing; pass
  `clockster.WithIdempotencyKey` so a retry does not file them twice.
- A constant per type of `PayrollSingleAdjustmentsType` in `sets.gen.go`, and the six in
  `PayrollSingleAdjustmentsTypeValues()`.

## 0.3.0

### May break your build

Nothing. The five values are additive.

### Changed in the API

Reaches you whether or not you update this package.

- A document's `type` also takes `driver_license`, `birth_certificate`, `marriage_certificate`,
  `divorce_certificate` and `change_fio_certificate` — in the `types` filter and in
  `POST /documents/upsert`.

### New

- A constant per new value in `sets.gen.go`, and the five in `DocumentsTypeValues()`.

## 0.2.0

### May break your build

Nothing. The fields stayed `string`, and the constants are new names beside them.

### Changed in the API

Reaches you whether or not you update this package.

- A location's `radius` must be between 50 and 700. A value outside that is refused.
- `radius` cannot be cleared. Leave the key out to keep what is stored.
- `?employment=` takes only the ten terms of the set. An unknown one is refused rather than
  answering an empty page.

### New

- `sets.gen.go`: a constant per value of every set of values, and `UsersRoleValues()` beside each.
- Every request body field carries a doc comment, and `Priority` documents its two values.
