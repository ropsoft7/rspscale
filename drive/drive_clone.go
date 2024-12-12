// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by scale.ropsoft.cloud/cmd/cloner; DO NOT EDIT.

package drive

// Clone makes a deep copy of Share.
// The result aliases no memory with the original.
func (src *Share) Clone() *Share {
	if src == nil {
		return nil
	}
	dst := new(Share)
	*dst = *src
	dst.BookmarkData = append(src.BookmarkData[:0:0], src.BookmarkData...)
	return dst
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _ShareCloneNeedsRegeneration = Share(struct {
	Name         string
	Path         string
	As           string
	BookmarkData []byte
}{})

// Clone duplicates src into dst and reports whether it succeeded.
// To succeed, <src, dst> must be of types <*T, *T> or <*T, **T>,
// where T is one of Share.
func Clone(dst, src any) bool {
	switch src := src.(type) {
	case *Share:
		switch dst := dst.(type) {
		case *Share:
			*dst = *src.Clone()
			return true
		case **Share:
			*dst = src.Clone()
			return true
		}
	}
	return false
}