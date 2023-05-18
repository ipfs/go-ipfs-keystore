# go-ipfs-keystore

> go-ipfs-keystore implements keystores for ipfs
[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://protocol.ai)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![Travis CI](https://travis-ci.com/ipfs/go-ipfs-keystore.svg?branch=master)](https://travis-ci.com/ipfs/go-ipfs-keystore)
[![Go Reference](https://pkg.go.dev/badge/github.com/ipfs/go-ipfs-keystore.svg)](https://pkg.go.dev/github.com/ipfs/go-ipfs-keystore)

## ❗ This repo is no longer maintained.
👉 We highly recommend switching to the maintained version at https://github.com/ipfs/boxo/tree/main/keystore.
🏎️ Good news!  There is [tooling and documentation](https://github.com/ipfs/boxo#migrating-to-boxo) to expedite a switch in your repo. 

⚠️ If you continue using this repo, please note that security fixes will not be provided (unless someone steps in to maintain it).

📚 Learn more, including how to take the maintainership mantle or ask questions, [here](https://github.com/ipfs/boxo/wiki/Copied-or-Migrated-Repos-FAQ).


go-ipfs-keystore provides the Keystore interface for key management.  Keystores support adding, retrieving, and deleting keys as well as listing all keys and checking for membership.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-ipfs-keystore` works like a regular Go module:
```
> go get github.com/ipfs/go-ipfs-keystore
```

## Usage
```
import "github.com/ipfs/go-ipfs-keystore"
```

## Contribute

PRs accepted.

## License

This project is dual-licensed under Apache 2.0 and MIT terms:

- Apache License, Version 2.0, ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
- MIT license ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)
