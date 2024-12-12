// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"runtime"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
	"scale.ropsoft.cloud/clientupdate"
	"scale.ropsoft.cloud/version"
	"scale.ropsoft.cloud/version/distro"
)

var updateCmd = &ffcli.Command{
	Name:       "update",
	ShortUsage: "rspscale update",
	ShortHelp:  "Update Rspscale to the latest/different version",
	Exec:       runUpdate,
	FlagSet: (func() *flag.FlagSet {
		fs := newFlagSet("update")
		fs.BoolVar(&updateArgs.yes, "yes", false, "update without interactive prompts")
		fs.BoolVar(&updateArgs.dryRun, "dry-run", false, "print what update would do without doing it, or prompts")
		// These flags are not supported on several systems that only provide
		// the latest version of Rspscale:
		//
		//  - Arch (and other pacman-based distros)
		//  - Alpine (and other apk-based distros)
		//  - FreeBSD (and other pkg-based distros)
		//  - Unraid/QNAP/Synology
		//  - macOS
		if distro.Get() != distro.Arch &&
			distro.Get() != distro.Alpine &&
			distro.Get() != distro.QNAP &&
			distro.Get() != distro.Synology &&
			runtime.GOOS != "freebsd" &&
			runtime.GOOS != "darwin" {
			fs.StringVar(&updateArgs.track, "track", "", `which track to check for updates: "stable" or "unstable" (dev); empty means same as current`)
			fs.StringVar(&updateArgs.version, "version", "", `explicit version to update/downgrade to`)
		}
		return fs
	})(),
}

var updateArgs struct {
	yes     bool
	dryRun  bool
	track   string // explicit track; empty means same as current
	version string // explicit version; empty means auto
}

func runUpdate(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}
	if updateArgs.version != "" && updateArgs.track != "" {
		return errors.New("cannot specify both --version and --track")
	}
	err := clientupdate.Update(clientupdate.Arguments{
		Version: updateArgs.version,
		Track:   updateArgs.track,
		Logf:    func(f string, a ...any) { printf(f+"\n", a...) },
		Stdout:  Stdout,
		Stderr:  Stderr,
		Confirm: confirmUpdate,
	})
	if errors.Is(err, errors.ErrUnsupported) {
		return errors.New("The 'update' command is not supported on this platform; see https://scale.ropsoft.cloud/s/client-updates")
	}
	return err
}

func confirmUpdate(ver string) bool {
	if updateArgs.yes {
		fmt.Printf("Updating Rspscale from %v to %v; --yes given, continuing without prompts.\n", version.Short(), ver)
		return true
	}

	if updateArgs.dryRun {
		fmt.Printf("Current: %v, Latest: %v\n", version.Short(), ver)
		return false
	}

	msg := fmt.Sprintf("This will update Rspscale from %v to %v. Continue?", version.Short(), ver)
	return promptYesNo(msg)
}

// PromptYesNo takes a question and prompts the user to answer the
// question with a yes or no. It appends a [y/n] to the message.
func promptYesNo(msg string) bool {
	fmt.Print(msg + " [y/n] ")
	var resp string
	fmt.Scanln(&resp)
	resp = strings.ToLower(resp)
	switch resp {
	case "y", "yes", "sure":
		return true
	}
	return false
}
