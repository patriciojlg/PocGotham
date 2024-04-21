{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    devenv.url = "github:cachix/devenv";
  };

  outputs = { self, nixpkgs, devenv, ... } @ inputs:
  let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in
  {

    devShell.${system} = devenv.lib.mkShell {
      inherit inputs pkgs;

      modules = [
        ({ pkgs, config, ...}:
        {
          env.PROJECT_NAME = "gotth_poc";
          env.AWS_PROFILE = "monk3yd-rlab";
          env.AWS_ACCOUNT_ID = "963485456147";
          env.AWS_REGION_NAME = "us-east-1";
          env.PORT = ":8080";

          languages.go.enable = true;

          packages = [
            pkgs.pulumi
            pkgs.pulumiPackages.pulumi-language-go
            pkgs.air
          ];
        })
      ];
    };
  };
}
