{pkgs ? import <nixpkgs> {}}:
with pkgs;
  writeShellApplication {
    name = "count";

    runtimeInputs = [
      yq
      curl
      coreutils
    ];

    text = builtins.readFile ../src/count.sh;
  }
