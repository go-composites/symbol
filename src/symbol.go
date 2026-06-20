package Symbol

import (
	"fmt"
	"sync"
)

// Interface is the structural contract of an interned, immutable identifier —
// the go-composites equivalent of Ruby's Symbol (`:name`). Two symbols built
// from the same name are the SAME underlying instance, so identity comparison
// is meaningful and cheap.
type Interface interface {
	// ToGoString returns the symbol's name as a Go string.
	ToGoString() string
	// Name is an alias for ToGoString.
	Name() string
	// Equal reports whether the given symbol denotes the same interned name.
	Equal(Interface) bool
	// IsNull reports whether this is the Null-Object symbol.
	IsNull() bool
	// Inspect renders a one-line `:name` view of the symbol.
	Inspect() string
}

// data is the immutable payload of a real symbol. Instances are never created
// directly: New owns the only constructor path so that interning holds.
type data struct {
	name string
}

var (
	registryMutex sync.Mutex
	registry      = map[string]*data{}
)

/*
New returns the interned Symbol for name. Calling New with the same name always
yields the SAME underlying instance, so New("foo") == New("foo"). The registry
is guarded by a sync.Mutex, making interning safe for concurrent callers.
*/
func New(name string) Interface {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	if existing, ok := registry[name]; ok {
		return existing
	}
	created := &data{name: name}
	registry[name] = created
	return created
}

// ToGoString returns the symbol's name.
func (d *data) ToGoString() string {
	return d.name
}

// Name is an alias for ToGoString.
func (d *data) Name() string {
	return d.name
}

/*
Equal reports whether other denotes the same interned name as the receiver.
Because symbols are interned, this is equivalent to pointer identity for real
symbols; comparing names also makes Equal robust against the Null-Object.
*/
func (d *data) Equal(other Interface) bool {
	if other.IsNull() {
		return false
	}
	return d.name == other.Name()
}

// IsNull reports that this is a real (non-null) Symbol.
func (d *data) IsNull() bool {
	return false
}

// Inspect renders the symbol as the literal `:name`.
func (d *data) Inspect() string {
	return fmt.Sprintf(":%s", d.name)
}

// null is the Null-Object variant of a Symbol: a placeholder that honours the
// full Interface without ever being nil. It has an empty name and is never
// equal to any symbol, including itself.
type null struct{}

/*
Null returns the Null-Object Symbol.
*/
func Null() Interface {
	return &null{}
}

// ToGoString returns the empty name for the null Symbol.
func (n *null) ToGoString() string {
	return ""
}

// Name is an alias for ToGoString.
func (n *null) Name() string {
	return ""
}

// Equal is always false for the null Symbol.
func (n *null) Equal(other Interface) bool {
	return false
}

// IsNull reports that this is the null Symbol.
func (n *null) IsNull() bool {
	return true
}

// Inspect renders the null Symbol as the literal `:null`.
func (n *null) Inspect() string {
	return ":null"
}
