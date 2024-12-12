// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package version_test

import (
	"flag"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"scale.ropsoft.cloud/version"
)

var (
	findModuleInfo = version.ExportFindModuleInfo
	cmdName        = version.ExportCmdName
)

func TestFindModuleInfo(t *testing.T) {
	dir := t.TempDir()
	name := filepath.Join(dir, "rspscaled-version-test")
	out, err := exec.Command("go", "build", "-o", name, "scale.ropsoft.cloud/cmd/rspscaled").CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build rspscaled: %v\n%s", err, out)
	}
	modinfo, err := findModuleInfo(name)
	if err != nil {
		t.Fatal(err)
	}
	prefix := "path\tscale.ropsoft.cloud/cmd/rspscaled\nmod\tscale.ropsoft.cloud"
	if !strings.HasPrefix(modinfo, prefix) {
		t.Errorf("unexpected modinfo contents %q", modinfo)
	}
}

var findModuleInfoName = flag.String("module-info-file", "", "if non-empty, test findModuleInfo against this filename")

func TestFindModuleInfoManual(t *testing.T) {
	exe := *findModuleInfoName
	if exe == "" {
		t.Skip("skipping without --module-info-file filename")
	}
	cmd := cmdName(exe)
	mod, err := findModuleInfo(exe)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got %q from: %s", cmd, mod)
}
