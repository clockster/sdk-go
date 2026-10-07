# Changelog

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
