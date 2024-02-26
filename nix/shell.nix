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
        yq
        curl
        shfmt
        coreutils
      ]
      ++ (oa.nativeBuildInputs or []);
  })
