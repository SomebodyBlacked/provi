{ pkgs, ... }: {
  channel = "stable-23.11"; # or "unstable"
  packages = [ pkgs.go ];
  services.mysql.enable = true;
  services.mysql.package = pkgs.mysql80;
  idx = {
    extensions = [
      "golang.go"
      "Equinusocio.vsc-material-theme"
      "equinusocio.vsc-material-theme-icons"
      "rangav.vscode-thunder-client"
      "redhat.vscode-yaml"
    ];
  };
}
