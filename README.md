<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/symbol" width="720"></p>

# symbol

[![ci](https://github.com/go-composites/symbol/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/symbol/actions/workflows/ci.yml)

The interned-identifier composite of [go-composites](https://github.com/go-composites).
A `Symbol` is an immutable, interned name — the Go equivalent of Ruby's
`:name`. Two symbols built from the same string are the SAME underlying
instance, so identity comparison is meaningful and cheap, and a `Symbol`
never goes `nil`.

## Install

```sh
go get github.com/go-composites/symbol
```

## Usage

```golang
package main

import (
    "fmt"

    Symbol "github.com/go-composites/symbol/src"
)

func main() {
    foo := Symbol.New("foo")
    alsoFoo := Symbol.New("foo")
    bar := Symbol.New("bar")

    fmt.Println(foo.Inspect())        // :foo
    fmt.Println(foo == alsoFoo)       // true  — interned, same instance
    fmt.Println(foo.Equal(alsoFoo))   // true
    fmt.Println(foo.Equal(bar))       // false
    fmt.Println(Symbol.Null().IsNull()) // true
}
```

## API

### Constructors

- `New(name string) Interface` — the interned `Symbol` for `name`. Calling
  `New` twice with the same `name` returns the SAME underlying instance, so
  `New("foo") == New("foo")`. Interning is guarded by a `sync.Mutex` and is
  safe for concurrent callers.
- `Null() Interface` — the Null-Object `Symbol`.

### Methods

| method | returns | notes |
| --- | --- | --- |
| `ToGoString()` | Go `string` | the symbol's name (`""` for the Null-Object) |
| `Name()` | Go `string` | alias for `ToGoString()` |
| `Equal(other)` | Go `bool` | `true` iff `other` is the same interned name |
| `IsNull()` | Go `bool` | `true` only for the Null-Object |
| `Inspect()` | Go `string` | one-line `:name` view (`:null` for the Null-Object) |

## Null-Object

`Null()` returns the never-nil Null-Object `Symbol`: it has an empty name,
`IsNull()` reports `true`, `Inspect()` renders `:null`, and `Equal` is always
`false` — a null symbol is never equal to any symbol, including itself.

## License

BSD-3-Clause © the go-composites/symbol authors.
