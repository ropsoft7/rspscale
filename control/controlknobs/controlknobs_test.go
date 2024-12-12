// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package controlknobs

import (
	"reflect"
	"testing"
)

func TestAsDebugJSON(t *testing.T) {
	var nilPtr *Knobs
	if got := nilPtr.AsDebugJSON(); got != nil {
		t.Errorf("AsDebugJSON(nil) = %v; want nil", got)
	}
	k := new(Knobs)
	got := k.AsDebugJSON()
	if want := reflect.TypeFor[Knobs]().NumField(); len(got) != want {
		t.Errorf("AsDebugJSON map has %d fields; want %v", len(got), want)
	}
}
