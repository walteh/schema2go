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

// writeDotPath writes the current path and colon
func (d *Dumper) writeDotPath() {
	if len(d.path) > 0 {
		d.buf.WriteString(strings.Join(d.path, "."))
		if d.UseTabWriter {
			d.buf.WriteString("\t= ")
		} else {
			d.buf.WriteString(": ")
		}
	}
}

func (d *Dumper) shouldWriteDotPath(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct, reflect.Interface, reflect.Map, reflect.Slice:
		return false
	default:
		return true
	}
}

// writeNilValue writes a nil value in dot notation format
func (d *Dumper) writeNilValue(v reflect.Value) {
	d.writeDotPath()
	d.buf.WriteString("(nil)")
	d.buf.WriteByte('\n')
}

// dumpMapWithDotNotation handles map dumping with dot notation format
func (d *Dumper) dumpMapWithDotNotation(v reflect.Value) {
	d.initRootPath(v)

	if v.IsNil() {
		d.writeNilValue(v)
		return
	}

	keys := v.MapKeys()
	if len(keys) == 0 {
		d.writeDotPath()
		d.buf.WriteByte('\n')
		return
	}

	for _, k := range keys {
		// Add map key to path
		keyStr := fmt.Sprintf("%v", k.Interface())
		d.path = append(d.path, keyStr)

		// Write path and value
		d.writeDotPath()
		d.dump(v.MapIndex(k))
		d.buf.WriteByte('\n')

		// Remove map key from path
		if len(d.path) > 0 {
			d.path = d.path[:len(d.path)-1]
		}
	}
}

// dumpStructWithDotNotation handles struct dumping with dot notation format
func (d *Dumper) dumpStructWithDotNotation(v reflect.Value) {
	d.initRootPath(v)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields if HidePrivateFields is true
		if !fieldType.IsExported() && d.HidePrivateFields {
			continue
		}

		// Add field name to path
		d.path = append(d.path, fieldType.Name)

		// Write path and value
		// only write the path
		if d.shouldWriteDotPath(field) {
			d.writeDotPath()
		}
		// d.buf.WriteString(fmt.Sprintf(" {%s} ", field.Kind()))

		if field.Kind() == reflect.Ptr && field.IsNil() {
			d.buf.WriteString("(nil)")
		} else if !field.CanInterface() && d.HidePrivateFields {
			d.buf.WriteString("(unexported)")
		} else {
			d.dump(field)
		}
		if d.shouldWriteDotPath(field) {
			d.buf.WriteByte('\n')
		}

		// Remove field name from path
		if len(d.path) > 0 {
			d.path = d.path[:len(d.path)-1]
		}
	}
}

// dumpSliceWithDotNotation handles slice/array dumping with dot notation format
func (d *Dumper) dumpSliceWithDotNotation(v reflect.Value) {
	d.initRootPath(v)

	if v.IsNil() {
		d.writeNilValue(v)
		return
	}

	if v.Len() == 0 {
		d.writeDotPath()
		d.buf.WriteByte('\n')
		return
	}

	for i := 0; i < v.Len(); i++ {
		// Add array index to path
		if len(d.path) > 0 {
			lastPart := d.path[len(d.path)-1]
			d.path[len(d.path)-1] = fmt.Sprintf("%s[%d]", lastPart, i)
		} else {
			d.path = append(d.path, fmt.Sprintf("[%d]", i))
		}
		idx := v.Index(i)
		// Write path and value
		if d.shouldWriteDotPath(idx) {
			d.writeDotPath()
		}

		d.dump(v.Index(i))

		if d.shouldWriteDotPath(idx) {
			d.buf.WriteByte('\n')
		}

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
