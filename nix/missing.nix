{pkgs ? import <nixpkgs> {}}:
with pkgs;
  writeShellApplication {
    name = "missing";

    runtimeInputs = [
      yq
      curl
      coreutils
    ];

    text = builtins.readFile ../src/missing.sh;
  }
