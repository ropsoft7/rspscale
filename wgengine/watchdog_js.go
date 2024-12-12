// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build js

package wgengine

import "scale.ropsoft.cloud/net/dns/resolver"

type watchdogEngine struct {
	Engine
	wrap Engine
}

func (e *watchdogEngine) GetResolver() (r *resolver.Resolver, ok bool) {
	return nil, false
}
