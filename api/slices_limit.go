// Copyright 2018-20 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package api

import (
	"errors"
	"fmt"
)

var (
	ErrRangeInvalid      = errors.New("range invalid")
	ErrLengthUndefined   = errors.New("limit undefined")
	ErrLengthNotProvided = errors.New("end is nil so length must be provided")
)

// ScopeLimit is used to specify a range. Both Start and End are inclusive.
// A nil value means no limit, so a Start of nil means 0
// and an End of nil means no limit.
// The End value must always be equal to or larger than Start.
// Negative values are acceptable. A value of -2 means the second last row.
type ScopeLimit struct {
	Start *int
	End   *int
}

// String implements Stringer interface.
func (r ScopeLimit) String() string {
	if r.Start == nil {
		if r.End == nil {
			return "ScopeLimit:nil—nil"
		}
		return fmt.Sprintf("ScopeLimit:nil—%d", *r.End)
	}
	if r.End == nil {
		return fmt.Sprintf("ScopeLimit:%d—nil", *r.Start)
	}
	return fmt.Sprintf("ScopeLimit:%d—%d", *r.Start, *r.End)
}

// NRows returns the number of rows contained by ScopeLimit.
// If End is nil, then length must be provided.
func (r *ScopeLimit) NRows(length ...int) (int, error) {
	if len(length) > 0 {
		s, e, err := r.Limits(length[0])
		if err != nil {
			return 0, err
		}
		return e - s + 1, nil
	}

	if r.End == nil {
		return 0, ErrLengthNotProvided
	}

	var s int
	if r.Start != nil {
		s = *r.Start
	}
	if s < 0 || *r.End < 0 {
		return 0, ErrRangeInvalid
	}
	if *r.End < s {
		return 0, ErrRangeInvalid
	}
	return *r.End - s + 1, nil
}

// Limits is used to return the start and end limits of a ScopeLimit
// object for a given Dataframe or Series with length number of rows.
func (r *ScopeLimit) Limits(length int) (s int, e int, _ error) {
	if length <= 0 {
		return 0, 0, ErrLengthUndefined
	}

	if r.Start == nil {
		s = 0
	} else {
		if *r.Start < 0 {
			// negative
			s = length + *r.Start
		} else {
			s = *r.Start
		}
	}

	if r.End == nil {
		e = length - 1
	} else {
		if *r.End < 0 {
			// negative
			e = length + *r.End
		} else {
			e = *r.End
		}
	}
	if s < 0 || e < 0 {
		return 0, 0, ErrRangeInvalid
	}
	if s > e {
		return 0, 0, ErrRangeInvalid
	}
	if s >= length || e >= length {
		return 0, 0, ErrRangeInvalid
	}
	return
}

func (r *ScopeLimit) Limited(length int) (start, end int) {
	s, e, err := r.Limits(length)
	if err != nil {
		panic(err)
	}
	return s, e
}

// RangeFinite returns a ScopeLimit that has a finite span.
func RangeFinite(start int, end ...int) ScopeLimit {
	r := ScopeLimit{
		Start: &start,
	}
	if len(end) > 0 {
		r.End = &end[0]
	}
	return r
}

// IntsToRanges will convert an already (ascending) ordered list of ints to a slice of Ranges.
//
// Example:
//
//	import "sort"
//	ints := []int{2,4,5,6,8,10,11,45,46}
//	sort.Ints(ints)
//
//	fmt.Println(IntsToRanges(ints))
//	// Output: R{2,2}, R{4,6}, R{8,8}, R{10,11}, R{45,46}
func IntsToRanges(ints []int) []ScopeLimit {
	var out []ScopeLimit
OUTER:
	for i := 0; i < len(ints); i++ {
		v1 := ints[i]

		j := i + 1
		for {
			if j >= len(ints) {
				// j doesn't exist
				v2 := ints[j-1]
				out = append(out, ScopeLimit{Start: &v1, End: &v2})
				break OUTER
			} else {
				// j does exist
				v2 := ints[j]
				prevVal := ints[j-1]

				if (v2 != prevVal) && (v2 != prevVal+1) {
					out = append(out, ScopeLimit{Start: &v1, End: &prevVal})
					i = j - 1
					break
				}
				j++
				continue
			}
		}
	}

	return out
}
