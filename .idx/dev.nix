{ pkgs, ... }: {
  channel = "stable-23.11"; # or "unstable"
  packages = [ pkgs.go ];
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
