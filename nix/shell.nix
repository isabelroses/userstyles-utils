{
  yq,
  curl,
  shfmt,
  coreutils,
  callPackage,
}: let
  mainPkg = callPackage ./default.nix {};
in
  mainPkg.overrideAttrs (oa: {
    nativeBuildInputs =
      [
        coreutils
        curl
        yq
        shfmt
      ]
      ++ (oa.nativeBuildInputs or []);
  })
