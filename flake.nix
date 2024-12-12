# flake.nix describes a Nix source repository that provides
# development builds of Rspscale and the fork of the Go compiler
# toolchain that Rspscale maintains. It also provides a development
# environment for working on rspscale, for use with "nix develop".
#
# For more information about this and why this file is useful, see:
# https://nixos.wiki/wiki/Flakes
#
# Also look into direnv: https://direnv.net/, this can make it so that you can
# automatically get your environment set up when you change folders into the
# project.
#
# WARNING: currently, the packages provided by this flake are brittle,
# and importing this flake into your own Nix configs is likely to
# leave you with broken builds periodically.
#
# The issue is that building Rspscale binaries uses the buildGoModule
# helper from nixpkgs. This helper demands to know the content hash of
# all of the Go dependencies of this repo, in the form of a Nix SRI
# hash. This hash isn't automatically kept in sync with changes made
# to go.mod yet, and so every time we update go.mod while hacking on
# Rspscale, this flake ends up with a broken build due to hash
# mismatches.
#
# Right now, this flake is intended for use by Rspscale developers,
# who are aware of this mismatch and willing to live with it. At some
# point, we'll add automation to keep the hashes more in sync, at
# which point this caveat should go away.
#
# See https://github.com/ropsoft7/rspscale/issues/6845 for tracking
# how to fix this mismatch.
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    # Used by shell.nix as a compat shim.
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    flake-compat,
  }: let
    # rspscaleRev is the git commit at which this flake was imported,
    # or the empty string when building from a local checkout of the
    # rspscale repo.
    rspscaleRev = self.rev or "";
    # rspscale takes a nixpkgs package set, and builds Rspscale from
    # the same commit as this flake. IOW, it provides "rspscale built
    # from HEAD", where HEAD is "whatever commit you imported the
    # flake at".
    #
    # This is currently unfortunately brittle, because we have to
    # specify vendorHash, and that sha changes any time we alter
    # go.mod. We don't want to force a nix dependency on everyone
    # hacking on Rspscale, so this flake is likely to have broken
    # builds periodically until someone comes through and manually
    # fixes them up. I sure wish there was a way to express "please
    # just trust the local go.mod, vendorHash has no benefit here",
    # but alas.
    #
    # So really, this flake is for rspscale devs to dogfood with, if
    # you're an end user you should be prepared for this flake to not
    # build periodically.
    rspscale = pkgs:
      pkgs.buildGo123Module rec {
        name = "rspscale";

        src = ./.;
        vendorHash = pkgs.lib.fileContents ./go.mod.sri;
        nativeBuildInputs = pkgs.lib.optionals pkgs.stdenv.isLinux [pkgs.makeWrapper];
        ldflags = ["-X scale.ropsoft.cloud/version.gitCommitStamp=${rspscaleRev}"];
        CGO_ENABLED = 0;
        subPackages = ["cmd/rspscale" "cmd/rspscaled"];
        doCheck = false;

        # NOTE: We strip the ${PORT} and $FLAGS because they are unset in the
        # environment and cause issues (specifically the unset PORT). At some
        # point, there should be a NixOS module that allows configuration of these
        # things, but for now, we hardcode the default of port 41641 (taken from
        # ./cmd/rspscaled/rspscaled.defaults).
        postInstall = pkgs.lib.optionalString pkgs.stdenv.isLinux ''
          wrapProgram $out/bin/rspscaled --prefix PATH : ${pkgs.lib.makeBinPath [pkgs.iproute2 pkgs.iptables pkgs.getent pkgs.shadow]}
          wrapProgram $out/bin/rspscale --suffix PATH : ${pkgs.lib.makeBinPath [pkgs.procps]}

          sed -i \
            -e "s#/usr/sbin#$out/bin#" \
            -e "/^EnvironmentFile/d" \
            -e 's/''${PORT}/41641/' \
            -e 's/$FLAGS//' \
            ./cmd/rspscaled/rspscaled.service

          install -D -m0444 -t $out/lib/systemd/system ./cmd/rspscaled/rspscaled.service
        '';
      };

    # This whole blob makes the rspscale package available for all
    # OS/CPU combos that nix supports, as well as a dev shell so that
    # "nix develop" and "nix-shell" give you a dev env.
    flakeForSystem = nixpkgs: system: let
      pkgs = nixpkgs.legacyPackages.${system};
      ts = rspscale pkgs;
    in {
      packages = {
        default = ts;
        rspscale = ts;
      };
      devShell = pkgs.mkShell {
        packages = with pkgs; [
          curl
          git
          gopls
          gotools
          graphviz
          perl
          go_1_23
          yarn

          # qemu and e2fsprogs are needed for natlab
          qemu
          e2fsprogs
        ];
      };
    };
  in
    flake-utils.lib.eachDefaultSystem (system: flakeForSystem nixpkgs system);
}
# nix-direnv cache busting line: sha256-xO1DuLWi6/lpA9ubA2ZYVJM+CkVNA5IaVGZxX9my0j0=
