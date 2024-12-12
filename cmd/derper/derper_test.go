// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"scale.ropsoft.cloud/derp/derphttp"
	"scale.ropsoft.cloud/tstest/deptest"
)

func TestProdAutocertHostPolicy(t *testing.T) {
	tests := []struct {
		in     string
		wantOK bool
	}{
		{"derp.scale.ropsoft.cloud", true},
		{"derp.scale.ropsoft.cloud.", true},
		{"derp1.scale.ropsoft.cloud", true},
		{"derp1b.scale.ropsoft.cloud", true},
		{"derp2.scale.ropsoft.cloud", true},
		{"derp02.scale.ropsoft.cloud", true},
		{"derp-nyc.scale.ropsoft.cloud", true},
		{"derpfoo.scale.ropsoft.cloud", true},
		{"derp02.bar.scale.ropsoft.cloud", false},
		{"example.net", false},
	}
	for _, tt := range tests {
		got := prodAutocertHostPolicy(context.Background(), tt.in) == nil
		if got != tt.wantOK {
			t.Errorf("f(%q) = %v; want %v", tt.in, got, tt.wantOK)
		}
	}
}

func TestNoContent(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "no challenge",
		},
		{
			name:  "valid challenge",
			input: "input",
			want:  "response input",
		},
		{
			name:  "valid challenge hostname",
			input: "ts_derp99b.scale.ropsoft.cloud",
			want:  "response ts_derp99b.scale.ropsoft.cloud",
		},
		{
			name:  "invalid challenge",
			input: "foo\x00bar",
			want:  "",
		},
		{
			name:  "whitespace invalid challenge",
			input: "foo bar",
			want:  "",
		},
		{
			name:  "long challenge",
			input: strings.Repeat("x", 65),
			want:  "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "https://localhost/generate_204", nil)
			if tt.input != "" {
				req.Header.Set(derphttp.NoContentChallengeHeader, tt.input)
			}
			w := httptest.NewRecorder()
			derphttp.ServeNoContent(w, req)
			resp := w.Result()

			if tt.want == "" {
				if h, found := resp.Header[derphttp.NoContentResponseHeader]; found {
					t.Errorf("got %+v; expected no response header", h)
				}
				return
			}

			if got := resp.Header.Get(derphttp.NoContentResponseHeader); got != tt.want {
				t.Errorf("got %q; want %q", got, tt.want)
			}
		})
	}
}

func TestDeps(t *testing.T) {
	deptest.DepChecker{
		BadDeps: map[string]string{
			"testing":                            "do not use testing package in production code",
			"gvisor.dev/gvisor/pkg/buffer":       "https://github.com/ropsoft7/rspscale/issues/9756",
			"gvisor.dev/gvisor/pkg/cpuid":        "https://github.com/ropsoft7/rspscale/issues/9756",
			"gvisor.dev/gvisor/pkg/tcpip":        "https://github.com/ropsoft7/rspscale/issues/9756",
			"gvisor.dev/gvisor/pkg/tcpip/header": "https://github.com/ropsoft7/rspscale/issues/9756",
			"scale.ropsoft.cloud/net/packet":           "not needed in derper",
			"github.com/gaissmai/bart":           "not needed in derper",
			"database/sql/driver":                "not needed in derper", // previously came in via github.com/google/uuid
		},
	}.Check(t)
}

func TestTemplate(t *testing.T) {
	buf := &bytes.Buffer{}
	err := homePageTemplate.Execute(buf, templateData{
		ShowAbuseInfo: true,
		Disabled:      true,
		AllowDebug:    true,
	})
	if err != nil {
		t.Fatal(err)
	}

	str := buf.String()
	if !strings.Contains(str, "If you suspect abuse") {
		t.Error("Output is missing abuse mailto")
	}
	if !strings.Contains(str, "Rspscale Security Policies") {
		t.Error("Output is missing Rspscale Security Policies link")
	}
	if !strings.Contains(str, "Status:") {
		t.Error("Output is missing disabled status")
	}
	if !strings.Contains(str, "Debug info") {
		t.Error("Output is missing debug info")
	}
}
