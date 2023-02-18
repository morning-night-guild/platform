# .docker

## Setup

`.env` is required to Makefile.  
`.env` file is prepared to switch the port number in each developer's environment.

```shell
touch .env
make env
```

## Help

```shell
make
```

## Troubleshoot

- If you encounter the following error, delete frontend/node_modules and re-run make up

```
Specifically the "@esbuild/darwin-arm64" package is present but this platform
needs the "@esbuild/linux-arm64" package instead. People often get into this
situation by installing esbuild on Windows or macOS and copying "node_modules"
into a Docker image that runs Linux, or by copying "node_modules" between
Windows and WSL environments.
```

