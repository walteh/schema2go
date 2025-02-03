// 📦 originally copied by copyrc
// 🔗 source: https://raw.githubusercontent.com/fsamin/go-dump/c877b85e7fc1e8c59fad34c8f986e786ea76dc01/dump.go
// 📝 license: Apache-2.0
// ℹ️ see .copyrc.lock for more details

package dump

import (
	"io"
	"os"
)

// Dump displays the passed parameter properties to standard out such as complete types and all
// pointer addresses used to indirect to the final value.
// See Fdump if you would prefer dumping to an arbitrary io.Writer or Sdump to
// get the formatted result as a string.
func Dump(i interface{}, formatters ...KeyFormatterFunc) error {
	return Fdump(os.Stdout, i, formatters...)
}

// Sdump returns a string with the passed arguments formatted exactly the same as Dump.
func Sdump(i interface{}, formatters ...KeyFormatterFunc) (string, error) {
	if formatters == nil {
		formatters = []KeyFormatterFunc{WithDefaultFormatter()}
	}
	e := NewDefaultEncoder()
	e.Formatters = formatters
	return e.Sdump(i)
}

// Fdump formats and displays the passed arguments to io.Writer w. It formats exactly the same as Dump.
func Fdump(w io.Writer, i interface{}, formatters ...KeyFormatterFunc) error {
	if formatters == nil {
		formatters = []KeyFormatterFunc{WithDefaultFormatter()}
	}
	e := NewEncoder(w)
	e.Formatters = formatters
	return e.Fdump(i)
}

// ToMap dumps argument as a map[string]interface{}
func ToMap(i interface{}, formatters ...KeyFormatterFunc) (map[string]interface{}, error) {
	if formatters == nil {
		formatters = []KeyFormatterFunc{WithDefaultFormatter()}
	}
	e := NewDefaultEncoder()
	e.Formatters = formatters
	return e.ToMap(i)
}

// ToStringMap formats the argument as a map[string]string. It formats exactly the same as Dump.
func ToStringMap(i interface{}, formatters ...KeyFormatterFunc) (map[string]string, error) {
	if formatters == nil {
		formatters = []KeyFormatterFunc{WithDefaultFormatter()}
	}
	e := NewDefaultEncoder()
	e.Formatters = formatters
	return e.ToStringMap(i)
}

// MustSdump is a helper that wraps a call to a function returning (string, error)
// and panics if the error is non-nil.
func MustSdump(i interface{}, formatters ...KeyFormatterFunc) string {
	enc := NewDefaultEncoder()
	enc.Formatters = formatters
	s, err := enc.Sdump(i)
	if err != nil {
		panic(err)
	}
	return s
}
