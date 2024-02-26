{
  callPackage,
  coreutils,
  curl,
  yq,
}: let
  mainPkg = callPackage ./default.nix {};
in
  mainPkg.overrideAttrs (oa: {
    nativeBuildInputs =
      [
        coreutils
        curl
        yq
      ]
      ++ (oa.nativeBuildInputs or []);
  })
