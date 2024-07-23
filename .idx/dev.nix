# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  # Which nixpkgs channel to use.
  channel = "stable-23.11"; # or "unstable"
  services.docker.enable = true;
  # Use https://search.nixos.org/packages to find packages
  packages = [
    pkgs.go
    pkgs.protobuf
    pkgs.docker
    pkgs.grpc-tools
    pkgs.nodejs_20
    pkgs.nodePackages.nodemon
    pkgs.protoc-gen-go
    pkgs.grpc
  ];
  # Sets environment variables in the workspace
  env = {};
  idx = {
    # Search for the extensions you want on https://open-vsx.org/ and use "publisher.id"
    extensions = [
      "golang.go"
    ];
    workspace = {
      onCreate = {
      };
    };
    # Enable previews and customize configuration
    previews = {
    };
  };
}
