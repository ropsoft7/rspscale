// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package childproc allows other packages to register "rspscaled be-child"
// child process hook code. This avoids duplicating build tags in the
// rspscaled package. Instead, the code that needs to fork/exec the self
// executable (when it's rspscaled) can instead register the code
// they want to run.
package childproc

var Code = map[string]func([]string) error{}

// Add registers code f to run as 'rspscaled be-child <typ> [args]'.
func Add(typ string, f func(args []string) error) {
	if _, dup := Code[typ]; dup {
		panic("dup hook " + typ)
	}
	Code[typ] = f
}
