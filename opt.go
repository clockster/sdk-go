package clockster

import "encoding/json"

// Opt is a field a request may leave out.
//
// This API reads an omitted field and a null one differently: an omitted one keeps whatever is
// stored, a null one clears it. A pointer says only one of those things, so optional fields are
// Opt and required ones that accept null are pointers.
//
//	clockster.UsersUpsertUser{
//		ExternalID: clockster.Set("HR-1"),
//		FirstName:  "Aisulu",
//		PositionID: clockster.Null[int64](), // clear the position
//		// DepartmentID is not set: not written, and the stored value stays.
//	}
type Opt[T any] struct {
	value T
	state optState
}

type optState uint8

const (
	optUnset optState = iota
	optValue
	optNull
)

// Set is a field carrying a value.
func Set[T any](value T) Opt[T] {
	return Opt[T]{value: value, state: optValue}
}

// Null is a field written as null, which clears what is stored.
func Null[T any]() Opt[T] {
	return Opt[T]{state: optNull}
}

// Value answers the value and whether there is one — false for a field left out and for a null.
func (o Opt[T]) Value() (T, bool) {
	return o.value, o.state == optValue
}

// Or answers the value, or the fallback where there is none.
func (o Opt[T]) Or(fallback T) T {
	if o.state != optValue {
		return fallback
	}

	return o.value
}

// Sent reports whether the field goes on the wire at all, as a value or as a null.
func (o Opt[T]) Sent() bool {
	return o.state != optUnset
}

// IsNull reports whether the field is written as null.
func (o Opt[T]) IsNull() bool {
	return o.state == optNull
}

// IsZero is what the `omitzero` tag reads: a field nobody set is not written.
func (o Opt[T]) IsZero() bool {
	return o.state == optUnset
}

// MarshalJSON writes the value, or null. A field left out never reaches here: `omitzero` drops it
// before this is called.
func (o Opt[T]) MarshalJSON() ([]byte, error) {
	if o.state != optValue {
		return []byte("null"), nil
	}

	return json.Marshal(o.value)
}

func (o *Opt[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*o = Null[T]()

		return nil
	}

	var value T

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*o = Set(value)

	return nil
}

// Ptr is a value to point at, for a required field that accepts null.
func Ptr[T any](value T) *T {
	return &value
}

// Deref is what a pointer holds, or the zero value where it holds nothing. A response field is a
// pointer where the API may answer null or leave it out, which is most of them.
func Deref[T any](value *T) T {
	if value == nil {
		var zero T

		return zero
	}

	return *value
}
