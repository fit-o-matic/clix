{
  description = "clix - extensible building blocks for CLI applications";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        go = pkgs.go_1_26;
      in
      {
        devShells.default = pkgs.mkShell {
          packages = [
            go
            pkgs.gopls
            pkgs.golangci-lint
            pkgs.gotools # goimports, etc.
          ];
        };
      });
}
