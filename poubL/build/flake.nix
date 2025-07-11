{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  
  outputs = { nixpkgs, ... }: {
    devShells.aarch64-darwin.default = 
      let pkgs = nixpkgs.legacyPackages.aarch64-darwin;
      in pkgs.mkShell {
        packages = with pkgs; [
          libvirt
          go
          qemu
        ];
        
        shellHook = ''
          echo "Environnement Dagger activé sur macOS"
        '';
      };
  };
}
