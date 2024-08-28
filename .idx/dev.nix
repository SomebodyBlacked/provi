{ pkgs, ... }: {
  channel = "stable-23.11"; # or "unstable"
  packages = [
    pkgs.go
  ];
  env = { };
  idx = {
    extensions = [
      "golang.go"
      "Equinusocio.vsc-material-theme"
      "equinusocio.vsc-material-theme-icons"
      "rangav.vscode-thunder-client"
    ];
    workspace = {
      onCreate = {
        default.openFiles = [ "server.go" ];
      };
    };
  };
}
