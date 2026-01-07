{
  description = "supersonic";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      pkgs = import nixpkgs { system = "x86_64-linux"; };
    in
    {
      devShell.x86_64-linux = pkgs.mkShell {
        buildInputs = [
          pkgs.pkg-config

          pkgs.go

          pkgs.libglvnd
          pkgs.mpv-unwrapped

          pkgs.xorg.libXxf86vm
          pkgs.xorg.libX11

          pkgs.wayland
          pkgs.wayland-protocols
          pkgs.libxkbcommon
        ];
      };
    };
}
