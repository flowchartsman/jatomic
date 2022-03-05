package jatomic

import (
	"encoding/json"
	"sync/atomic"
)

// A Bool is an atomic boolean value.
// The zero value is false.
type Bool struct {
	atomic.Bool
}

// MarshalJSON  by calling Load()
func (x Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON by calling Store()
func (x *Bool) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}

// An Int32 is an atomic int32. The zero value is zero.
type Int32 struct {
	atomic.Int32
}

// MarshalJSON  by calling Load()
func (x Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON by calling Store()
func (x *Int32) UnmarshalJSON(b []byte) error {
	var v int32
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}

// An Int64 is an atomic int64. The zero value is zero.
type Int64 struct {
	atomic.Int64
}

// MarshalJSON  by calling Load()
func (x Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON by calling Store()
func (x *Int64) UnmarshalJSON(b []byte) error {
	var v int64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}

// An Uint32 is an atomic uint32. The zero value is zero.
type Uint32 struct {
	atomic.Uint32
}

// MarshalJSON  by calling Load()
func (x Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON by calling Store()
func (x *Uint32) UnmarshalJSON(b []byte) error {
	var v uint32
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}

// An Uint64 is an atomic uint64. The zero value is zero.
type Uint64 struct {
	atomic.Uint64
}

// MarshalJSON  by calling Load()
func (x Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON by calling Store()
func (x *Uint64) UnmarshalJSON(b []byte) error {
	var v uint64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}
