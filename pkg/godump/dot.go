// Package godump provides pretty printing functionality
package godump

import (
	"fmt"
	"reflect"
	"strings"
)

// initRootPath initializes the root path for dot notation if needed
func (d *Dumper) initRootPath(val reflect.Value) {
	if d.DotNotation && len(d.path) == 0 {
		// Add type name as root for dot notation
		if val.Kind() == reflect.Ptr {
			d.path = append(d.path, val.Elem().Type().Name())
		} else {
			d.path = append(d.path, val.Type().Name())
		}
	}
}

// dumpMapWithDotNotation handles map dumping with dot notation format
func (d *Dumper) dumpMapWithDotNotation(v reflect.Value) {
	d.initRootPath(v)
	keys := v.MapKeys()
	for _, k := range keys {
		// Add map key to path
		keyStr := fmt.Sprintf("%v", k.Interface())
		d.path = append(d.path, keyStr)

		// Write path
		if len(d.path) > 0 {
			d.buf.WriteString(strings.Join(d.path, "."))
			d.buf.WriteString(": ")
		}

		// Dump value
		d.dump(v.MapIndex(k))
		d.buf.WriteByte('\n')

		// Remove map key from path
		d.path = d.path[:len(d.path)-1]
	}
}

// dumpStructWithDotNotation handles struct dumping with dot notation format
func (d *Dumper) dumpStructWithDotNotation(v reflect.Value) {
	d.initRootPath(v)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.CanInterface() && d.HidePrivateFields {
			continue
		}

		// Add field name to path
		d.path = append(d.path, t.Field(i).Name)

		// Write path
		if len(d.path) > 0 {
			d.buf.WriteString(strings.Join(d.path, "."))
			d.buf.WriteString(": ")
		}

		// Dump value
		d.dump(field)
		d.buf.WriteByte('\n')

		// Remove field name from path
		if len(d.path) > 0 {
			d.path = d.path[:len(d.path)-1]
		}
	}
}

// dumpSliceWithDotNotation handles slice/array dumping with dot notation format
func (d *Dumper) dumpSliceWithDotNotation(v reflect.Value) {
	d.initRootPath(v)
	for i := 0; i < v.Len(); i++ {
		// Add array index to path
		if len(d.path) > 0 {
			lastPart := d.path[len(d.path)-1]
			d.path[len(d.path)-1] = fmt.Sprintf("%s[%d]", lastPart, i)
		} else {
			d.path = append(d.path, fmt.Sprintf("[%d]", i))
		}

		// Write path
		if len(d.path) > 0 {
			d.buf.WriteString(strings.Join(d.path, "."))
			d.buf.WriteString(": ")
		}

		// Dump value
		d.dump(v.Index(i))
		d.buf.WriteByte('\n')

		// Restore path
		if len(d.path) > 0 {
			if len(d.path) == 1 {
				d.path = d.path[:0]
			} else {
				lastDot := strings.LastIndex(d.path[len(d.path)-1], "[")
				if lastDot >= 0 {
					d.path[len(d.path)-1] = d.path[len(d.path)-1][:lastDot]
				}
			}
		}
	}
}
