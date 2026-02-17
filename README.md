## go

**Handy Go libraries and utilities**

This repository contains a grab‑bag of small, focused Go packages that I use across projects:

- `pkg/stringz`: helpers for working with strings (e.g. `ToAlphaNumeric`).
- `pkg/structs`: high‑level helpers for turning structs into maps/slices and inspecting fields.
- `pkg/command`: utilities for conditionally running `exec.Cmd` instances (honouring a `DRY_RUN` env var).
- `pkg/httpx`: helpers for dealing with HTTP responses and errors.
- `pkg/logz`, `pkg/jsonx`, `pkg/jsontime`, `pkg/filez`, `pkg/slicez`, `pkg/mathx`, `pkg/reflectx`, etc: small, composable helpers for common tasks.

For full package documentation, see the Go reference for the module:
[pkg.go.dev/github.com/tartale/go](https://pkg.go.dev/github.com/tartale/go).

### Installation

Add the module to your project using `go get`:

```bash
go get github.com/tartale/go@latest
```

Or depend on a specific sub‑package:

```bash
go get github.com/tartale/go/pkg/stringz@latest
```

### Quick example

Using `stringz` and `structs` together:

```go
package main

import (
	"fmt"

	"github.com/tartale/go/pkg/stringz"
	"github.com/tartale/go/pkg/structs"
)

type User struct {
	Name  string `structs:"name"`
	Email string `structs:"email,omitempty"`
}

func main() {
	fmt.Println(stringz.ToAlphaNumeric("Hello, world!")) // "Helloworld"

	u := User{Name: "Tom", Email: "tom@example.com"}
	m := structs.Map(u)
	fmt.Printf("%#v\n", m)
}
```

### Development

- **Tests**: run all tests with `go test ./...`.
- **Packages**: see `pkg/` for the collection of small, focused libraries.

