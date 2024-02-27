{pkgs ? import <nixpkgs> {}}:
with pkgs;
  writeShellApplication {
    name = "umeta";

    runtimeInputs = [
      yq
      gum
      curl
      coreutils
    ];

    text = builtins.readFile ../src/umeta.sh;
  }
