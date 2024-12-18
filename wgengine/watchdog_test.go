// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package wgengine

import (
	"runtime"
	"testing"
	"time"

	"scale.ropsoft.cloud/health"
	"scale.ropsoft.cloud/util/usermetric"
)

func TestWatchdog(t *testing.T) {
	t.Parallel()

	var maxWaitMultiple time.Duration = 1
	if runtime.GOOS == "darwin" {
		// Work around slow close syscalls on Big Sur with content filter Network Extensions installed.
		// See https://github.com/ropsoft7/rspscale/issues/1598.
		maxWaitMultiple = 15
	}

	t.Run("default watchdog does not fire", func(t *testing.T) {
		t.Parallel()
		ht := new(health.Tracker)
		reg := new(usermetric.Registry)
		e, err := NewFakeUserspaceEngine(t.Logf, 0, ht, reg)
		if err != nil {
			t.Fatal(err)
		}

		e = NewWatchdog(e)
		e.(*watchdogEngine).maxWait = maxWaitMultiple * 150 * time.Millisecond
		e.(*watchdogEngine).logf = t.Logf
		e.(*watchdogEngine).fatalf = t.Fatalf

		e.RequestStatus()
		e.RequestStatus()
		e.RequestStatus()
		e.Close()
	})
}
