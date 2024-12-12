// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package netns

import (
	"testing"

	"scale.ropsoft.cloud/net/netmon"
)

func TestGetInterfaceIndex(t *testing.T) {
	oldVal := bindToInterfaceByRoute.Load()
	t.Cleanup(func() { bindToInterfaceByRoute.Store(oldVal) })
	bindToInterfaceByRoute.Store(true)

	tests := []struct {
		name string
		addr string
		err  string
	}{
		{
			name: "IP_and_port",
			addr: "8.8.8.8:53",
		},
		{
			name: "bare_ip",
			addr: "8.8.8.8",
		},
		{
			name: "invalid",
			addr: "!!!!!",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			idx, err := getInterfaceIndex(t.Logf, nil, tc.addr)
			if err != nil {
				if tc.err == "" {
					t.Fatalf("got unexpected error: %v", err)
				}
				if errstr := err.Error(); errstr != tc.err {
					t.Errorf("expected error %q, got %q", errstr, tc.err)
				}
			} else {
				t.Logf("getInterfaceIndex(%q) = %d", tc.addr, idx)
				if tc.err != "" {
					t.Fatalf("wanted error %q", tc.err)
				}
				if idx < 0 {
					t.Fatalf("got invalid index %d", idx)
				}
			}
		})
	}

	t.Run("NoRspscale", func(t *testing.T) {
		tsif, err := rspscaleInterface()
		if err != nil {
			t.Fatal(err)
		}
		if tsif == nil {
			t.Skip("no rspscale interface on this machine")
		}

		defaultIdx, err := netmon.DefaultRouteInterfaceIndex()
		if err != nil {
			t.Fatal(err)
		}

		idx, err := getInterfaceIndex(t.Logf, nil, "100.100.100.100:53")
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("rspscaleIdx=%d defaultIdx=%d idx=%d", tsif.Index, defaultIdx, idx)

		if idx == tsif.Index {
			t.Fatalf("got idx=%d; wanted not Rspscale interface", idx)
		} else if idx != defaultIdx {
			t.Fatalf("got idx=%d, want %d", idx, defaultIdx)
		}
	})
}
