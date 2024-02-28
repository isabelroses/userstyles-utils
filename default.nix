{buildGoModule}:
buildGoModule {
  pname = "uutils";
  version = "0.0.1";

  src = ./.;

  # sha256-0000000000000000000000000000000000000000000=
  vendorHash = "sha256-IhCYBlUwLx2OfWKHb6z5XGo3kMoow+BchHuvikm1KAY=";

  ldflags = ["-s" "-w"];
}
