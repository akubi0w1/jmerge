# jmerge

## Overview

jmerge can merge multiple json files.

## Install

```
# jmerge
$ go get github.com/akubi0w1/jmerge

# jmerge-cli - go get
$ go get github.com/akubi0w1/jmerge

# jmerge-cli
$ brew tap akubi0w1/tap
$ brew install akubi0w1/tap/jmerge-cli
```

## Examples

### jmerge

See [example/merge][example/merge] for merge example.

```go
package main

import (
	"fmt"

	"github.com/akubi0w1/jmerge"
)

func main() {
	basePath := "./base.json"
	overlayPath := "./overlay.json"

	out, err := jmerge.MergeJSONByFile(basePath, overlayPath, jmerge.MergeModeIgnore, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
```

```json
// base:
{
  "host": "127.0.0.1",
  "port": "8080",
  "mysql": {
    "user": "root",
    "password": "password",
    "host": "127.0.0.1",
    "port": "3306",
    "database": "main"
  }
}

// overlay:
{
  "port": "80",
  "mysql": {
    "user": "worker",
    "password": "akubi"
  }
}

// result
{
  "host": "127.0.0.1",
  "mysql": {
    "database": "main",
    "host": "127.0.0.1",
    "password": "akubi",
    "port": "3306",
    "user": "worker"
  },
  "port": "80"
}
```

### jmerge cli

See [cli/example][cli/example] for merge example using cli.

File structure:

```
./
├── base
│   └── http
│       └── setting.json
├── output
│   └── dev
│       └── http
│           └── setting.json
└── overlay
    ├── dev
    │   ├── http
    │   │   └── setting.json
    │   └── jmerge.yaml
    └── prd
        ├── http
        │   └── setting.json
        └── jmerge.yaml
```

Config:

```yaml
# config: jmerge.yaml

namespace: dev

# path of base file
base: ../../base

# output path for after merge
output: ../../output

# whether to format output
format: true

# merge target
merges:
  # mode is add or ignore
  - mode: add
    targets:
      - http/setting.json
```

Execute:

```shell
# execute cli
$ jmerge-cli merge
```

Usege:

```
merge multiple json files into one.

Usage:
  jmerge-cli merge [flags]

Flags:
  -c, --config string   config file path (default "jmerge.yaml")
  -h, --help            help for merge
```


[//]:#(refs)
[example/merge]: ./example/merge
[cli/example]: ./cli/example
[//]:#(refs)