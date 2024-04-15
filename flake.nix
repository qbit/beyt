{
  description = "beyt: beat time in your systray";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs =
    { self
    , nixpkgs
    ,
    }:
    let
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in
    {
      overlay = _: prev: { inherit (self.packages.${prev.system}) beyt; };

      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          beyt = with pkgs; pkgs.buildGoModule rec {
            pname = "beyt";
            version = "v0.1.1";
            src = ./.;

            vendorHash = "sha256-q1cpnirvqCyYvbyuN4CFqHpc5coW17rYz9ce8blSpSg=";

            nativeBuildInputs = [ pkg-config copyDesktopItems ];
            buildInputs = [
              glfw
              libGL
              libGLU
              openssh
              pkg-config
              glibc
              xorg.libXcursor
              xorg.libXi
              xorg.libXinerama
              xorg.libXrandr
              xorg.libXxf86vm
              xorg.xinput
            ];

            desktopItems = [
              (makeDesktopItem {
                name = "traygent";
                exec = pname;
                icon = pname;
                desktopName = pname;
              })
            ];
          };
        });

      defaultPackage = forAllSystems (system: self.packages.${system}.beyt);
      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            shellHook = ''
              PS1='\u@\h:\@; '
              nix flake run github:qbit/xin#flake-warn
              echo "Go `${pkgs.go}/bin/go version`"
            '';
            buildInputs = with pkgs; [
              git
              go_1_21
              gopls
              go-tools
              glxinfo

              glfw
              glibc
              pkg-config
              xorg.libXcursor
              xorg.libXi
              xorg.libXinerama
              xorg.libXrandr
              xorg.libXxf86vm
              xorg.xinput
              graphviz

              go-font
            ];
          };
        });
    };
}
