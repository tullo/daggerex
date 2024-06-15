# SP
- https://go.googlesource.com/example/

## Dagger Examples
- https://github.com/dagger/dagger/tree/main/examples

## Dagger Go SDK
- https://docs.dagger.io/sdk/go

## Linux

```shell
curl -L https://dl.dagger.io/dagger/install.sh | BIN_DIR=$HOME/.local/bin sh

dagger version
> dagger v0.11.7 (registry.dagger.io/engine) linux/amd64
```

### Shell completion

dagger has built-in shell completion. This is how you can install it for:

BASH:
1. Ensure that you install `bash-completion` using your package manager.
2. Add dagger completion to your personal bash completions dir
```
  mkdir -p /home/anda/.local/share/bash-completion/completions
  dagger completion bash > /home/anda/.local/share/bash-completion/completions/dagger
```
ZSH:
1. Generate a _dagger completion script and write it to a file within your $FPATH, e.g.:
      `dagger completion zsh > /usr/local/share/zsh/site-functions/_dagger`
2. Ensure that the following is present in your ~/.zshrc:
  ```
      autoload -U compinit
      compinit -i
  ```
FISH:
1. Generate a dagger.fish completion script and write it to a file within fish completions, e.g.:
`dagger completion fish > ~/.config/fish/completions/dagger.fish`
```
