{pkgs ? import <nixpkgs> {}}:
with pkgs;
  writeShellApplication {
    name = "count";
    runtimeInputs = [
      coreutils
      curl
      yq
    ];
    text = builtins.readFile ./count.sh;
  }
