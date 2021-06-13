# CmdBox Command Module Template

[![GoDoc](https://godoc.org/cmdbox-_isosec?status.svg)](https://godoc.org/cmdbox-_isosec)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This simple utility returns the current UTC time in YYYYMMDDhhmmss
format. This can be used to create a unique folder/file or as a general
time stamp. 

It is kept in sync with the [latest CmdBox tagged
version](https://github.com/rwxrob/cmdbox).

## Install 

This command can be installed as a standalone program or composed into a
CmdBox composite monolith.

Use `go install` to install as a standalone:

```
go install github.com/rossim2i2/cmdbox-isosec/isosec@latest
```

Use `import` with a blank identifier to be composed:

```go
import (
  "github.com/rwxrob/cmdbox"
  _ "github.com/rossim2i2/cmdbox-isosec"
)
```

## Embedded Documentation

See the [`cmd.go`](cmd.go) file itself for additional embedded
documentation about this command.

## Usage

```
isosec
isosec help
isosec version
```

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).

