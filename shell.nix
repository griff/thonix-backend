{nixpkgs ? (import /vagrant/nixpkgs {}) }:
with nixpkgs; callPackage ./default.nix {
  deps = ./test-deps.nix;
  thonix-frontend=thonix-frontend-head;
}