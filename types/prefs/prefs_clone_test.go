// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by scale.ropsoft.cloud/cmd/cloner; DO NOT EDIT.

package prefs

import (
	"net/netip"

	"scale.ropsoft.cloud/types/ptr"
)

// Clone makes a deep copy of TestPrefs.
// The result aliases no memory with the original.
func (src *TestPrefs) Clone() *TestPrefs {
	if src == nil {
		return nil
	}
	dst := new(TestPrefs)
	*dst = *src
	dst.StringSlice = *src.StringSlice.Clone()
	dst.IntSlice = *src.IntSlice.Clone()
	dst.StringStringMap = *src.StringStringMap.Clone()
	dst.IntStringMap = *src.IntStringMap.Clone()
	dst.AddrIntMap = *src.AddrIntMap.Clone()
	dst.Bundle1 = *src.Bundle1.Clone()
	dst.Bundle2 = *src.Bundle2.Clone()
	dst.Generic = *src.Generic.Clone()
	dst.BundleList = *src.BundleList.Clone()
	dst.StringBundleMap = *src.StringBundleMap.Clone()
	dst.IntBundleMap = *src.IntBundleMap.Clone()
	dst.AddrBundleMap = *src.AddrBundleMap.Clone()
	return dst
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestPrefsCloneNeedsRegeneration = TestPrefs(struct {
	Int32Item       Item[int32]
	UInt64Item      Item[uint64]
	StringItem1     Item[string]
	StringItem2     Item[string]
	BoolItem1       Item[bool]
	BoolItem2       Item[bool]
	StringSlice     List[string]
	IntSlice        List[int]
	AddrItem        Item[netip.Addr]
	StringStringMap Map[string, string]
	IntStringMap    Map[int, string]
	AddrIntMap      Map[netip.Addr, int]
	Bundle1         Item[*TestBundle]
	Bundle2         Item[*TestBundle]
	Generic         Item[*TestGenericStruct[int]]
	BundleList      StructList[*TestBundle]
	StringBundleMap StructMap[string, *TestBundle]
	IntBundleMap    StructMap[int, *TestBundle]
	AddrBundleMap   StructMap[netip.Addr, *TestBundle]
	Group           TestPrefsGroup
}{})

// Clone makes a deep copy of TestBundle.
// The result aliases no memory with the original.
func (src *TestBundle) Clone() *TestBundle {
	if src == nil {
		return nil
	}
	dst := new(TestBundle)
	*dst = *src
	if dst.Nested != nil {
		dst.Nested = ptr.To(*src.Nested)
	}
	return dst
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestBundleCloneNeedsRegeneration = TestBundle(struct {
	Name   string
	Nested *TestValueStruct
}{})

// Clone makes a deep copy of TestValueStruct.
// The result aliases no memory with the original.
func (src *TestValueStruct) Clone() *TestValueStruct {
	if src == nil {
		return nil
	}
	dst := new(TestValueStruct)
	*dst = *src
	return dst
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestValueStructCloneNeedsRegeneration = TestValueStruct(struct {
	Value int
}{})

// Clone makes a deep copy of TestGenericStruct.
// The result aliases no memory with the original.
func (src *TestGenericStruct[T]) Clone() *TestGenericStruct[T] {
	if src == nil {
		return nil
	}
	dst := new(TestGenericStruct[T])
	*dst = *src
	return dst
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
func _TestGenericStructCloneNeedsRegeneration[T ImmutableType](TestGenericStruct[T]) {
	_TestGenericStructCloneNeedsRegeneration(struct {
		Value T
	}{})
}

// Clone makes a deep copy of TestPrefsGroup.
// The result aliases no memory with the original.
func (src *TestPrefsGroup) Clone() *TestPrefsGroup {
	if src == nil {
		return nil
	}
	dst := new(TestPrefsGroup)
	*dst = *src
	return dst
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestPrefsGroupCloneNeedsRegeneration = TestPrefsGroup(struct {
	FloatItem      Item[float64]
	TestStringItem Item[TestStringType]
}{})
