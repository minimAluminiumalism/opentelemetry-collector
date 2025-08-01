// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pcommon

import (
	"iter"
	"slices"

	"go.opentelemetry.io/collector/pdata/internal"
)

// UInt64Slice represents a []uint64 slice.
// The instance of UInt64Slice can be assigned to multiple objects since it's immutable.
//
// Must use NewUInt64Slice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type UInt64Slice internal.UInt64Slice

func (ms UInt64Slice) getOrig() *[]uint64 {
	return internal.GetOrigUInt64Slice(internal.UInt64Slice(ms))
}

func (ms UInt64Slice) getState() *internal.State {
	return internal.GetUInt64SliceState(internal.UInt64Slice(ms))
}

// NewUInt64Slice creates a new empty UInt64Slice.
func NewUInt64Slice() UInt64Slice {
	orig := []uint64(nil)
	state := internal.StateMutable
	return UInt64Slice(internal.NewUInt64Slice(&orig, &state))
}

// AsRaw returns a copy of the []uint64 slice.
func (ms UInt64Slice) AsRaw() []uint64 {
	return internal.CopyOrigUInt64Slice(nil, *ms.getOrig())
}

// FromRaw copies raw []uint64 into the slice UInt64Slice.
func (ms UInt64Slice) FromRaw(val []uint64) {
	ms.getState().AssertMutable()
	*ms.getOrig() = internal.CopyOrigUInt64Slice(*ms.getOrig(), val)
}

// Len returns length of the []uint64 slice value.
// Equivalent of len(uInt64Slice).
func (ms UInt64Slice) Len() int {
	return len(*ms.getOrig())
}

// At returns an item from particular index.
// Equivalent of uInt64Slice[i].
func (ms UInt64Slice) At(i int) uint64 {
	return (*ms.getOrig())[i]
}

// All returns an iterator over index-value pairs in the slice.
func (ms UInt64Slice) All() iter.Seq2[int, uint64] {
	return func(yield func(int, uint64) bool) {
		for i := 0; i < ms.Len(); i++ {
			if !yield(i, ms.At(i)) {
				return
			}
		}
	}
}

// SetAt sets uint64 item at particular index.
// Equivalent of uInt64Slice[i] = val
func (ms UInt64Slice) SetAt(i int, val uint64) {
	ms.getState().AssertMutable()
	(*ms.getOrig())[i] = val
}

// EnsureCapacity ensures UInt64Slice has at least the specified capacity.
//  1. If the newCap <= cap, then is no change in capacity.
//  2. If the newCap > cap, then the slice capacity will be expanded to the provided value which will be equivalent of:
//     buf := make([]uint64, len(uInt64Slice), newCap)
//     copy(buf, uInt64Slice)
//     uInt64Slice = buf
func (ms UInt64Slice) EnsureCapacity(newCap int) {
	ms.getState().AssertMutable()
	oldCap := cap(*ms.getOrig())
	if newCap <= oldCap {
		return
	}

	newOrig := make([]uint64, len(*ms.getOrig()), newCap)
	copy(newOrig, *ms.getOrig())
	*ms.getOrig() = newOrig
}

// Append appends extra elements to UInt64Slice.
// Equivalent of uInt64Slice = append(uInt64Slice, elms...)
func (ms UInt64Slice) Append(elms ...uint64) {
	ms.getState().AssertMutable()
	*ms.getOrig() = append(*ms.getOrig(), elms...)
}

// MoveTo moves all elements from the current slice overriding the destination and
// resetting the current instance to its zero value.
func (ms UInt64Slice) MoveTo(dest UInt64Slice) {
	ms.getState().AssertMutable()
	dest.getState().AssertMutable()
	// If they point to the same data, they are the same, nothing to do.
	if ms.getOrig() == dest.getOrig() {
		return
	}
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = nil
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (ms UInt64Slice) MoveAndAppendTo(dest UInt64Slice) {
	ms.getState().AssertMutable()
	dest.getState().AssertMutable()
	if *dest.getOrig() == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.getOrig() = *ms.getOrig()
	} else {
		*dest.getOrig() = append(*dest.getOrig(), *ms.getOrig()...)
	}
	*ms.getOrig() = nil
}

// CopyTo copies all elements from the current slice overriding the destination.
func (ms UInt64Slice) CopyTo(dest UInt64Slice) {
	dest.getState().AssertMutable()
	*dest.getOrig() = internal.CopyOrigUInt64Slice(*dest.getOrig(), *ms.getOrig())
}

// Equal checks equality with another UInt64Slice
func (ms UInt64Slice) Equal(val UInt64Slice) bool {
	return slices.Equal(*ms.getOrig(), *val.getOrig())
}
