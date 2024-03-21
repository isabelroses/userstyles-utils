{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  buildInputs = with pkgs; [
    yq
    curl
    shfmt
    coreutils
  ];
}
