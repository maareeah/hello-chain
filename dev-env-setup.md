# Setup local dev
Use below cli's to setup your own cosmos-skd, tendermint, ignite versions.

```
    mkdir -p ~/my-devlop
    git clone git@github.com:tendermint/cosmos-sdk.git
    git checkout v0.46.7 && cd ..
    ln -s /Users/xxxx/my-devlop/cosmos-dev/cosmos-sdk /Users/xxxx/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.46.7

    git clone git@github.com:tendermint/tendermint.git
    cd tendermint && git checkout v0.34.24 && cd ..
    ln -s /Users/xxxx/my-devlop/cosmos-dev/tendermint  /Users/xxx/go/pkg/mod/github.com/tendermint/tendermint@v0.34.24
```

# Build your own ignite-CLI
https://github.com/ignite/cli/blob/main/dev-env-setup.md

```
    Ignite CLI version:             v0.25.2-dev
    Ignite CLI build date:          2023-01-10T18:20:28Z
    Ignite CLI source hash:         63a750d4a6b93d4f8a697511696076b892a0e60a
    Ignite CLI config version:      v1
    Cosmos SDK version:             v0.46.7
    Your OS:                        darwin
    Your arch:                      amd64
    Your Node.js version:           v16.15.1
    Your go version:                go version go1.19.5 darwin/amd64
    Your uname -a:                  Darwin MacBook-Pro.local 22.2.0 Darwin Kernel Version 22.2.0: Fri Nov 11 02:03:51 PST 2022; root:xnu-8792.61.2~4/RELEASE_ARM64_T6000 x86_64
    Your cwd:                       /Users/xxxx/my-devlop/cosmos-dev/ignite/cli
    Is on Gitpod:                   false

```

# Hello chain

https://docs.ignite.com/guide/hello
