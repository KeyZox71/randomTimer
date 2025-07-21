{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs =
    inputs@{ self, nixpkgs, ... }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forEachSupportedSystem =
        f:
        nixpkgs.lib.genAttrs supportedSystems (
          system:
          f {
            pkgs = import nixpkgs { inherit system; };
          }
        );
    in
    {
      packages = forEachSupportedSystem (
        { pkgs }:
        rec {
          default = randomTimer;
          randomTimer = pkgs.buildGoModule {
            src = self;
            pname = "randomTimer";
            subPackages = [ "cmd/randomTimer" ];
            version = "0.1.0";
            vendorHash = "sha256-mXx15hWdUJAM+LUBo49NU5gY0hNzz5T+9gZ1GlUYxM0="; 
          };
        }
      );
      devShells = forEachSupportedSystem (
        { pkgs }:
        {
          default = pkgs.mkShell {
            hardeningDisable = [ "all" ];
            nativeBuildInputs = with pkgs; [
            ];
            buildInputs = with pkgs; [
            ];
            packages = with pkgs; [
              nixd
              nixfmt-rfc-style

              go
              gopls
              self.packages.${pkgs.system}.default
            ];
          };
        }
      );
    };
}
