# Changelog

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
