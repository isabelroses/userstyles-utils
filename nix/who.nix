{pkgs ? import <nixpkgs> {}}:
with pkgs;
  writeShellApplication {
    name = "who";

    runtimeInputs = [
      yq
      gum
      curl
      coreutils
    ];

    text = builtins.readFile ../src/who.sh;
  }
