package diff

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/k0kubun/pp/v3"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/pmezard/go-difflib/difflib"
)

func TypedDiffExportedOnly[T any](want T, got T) string {
	printer := pp.New()
	printer.SetExportedOnly(true)
	printer.SetColoringEnabled(false)

	return diffTyped(printer, want, got)
}

func TypedDiff[T any](want T, got T) string {
	printer := pp.New()
	printer.SetExportedOnly(false)
	printer.SetColoringEnabled(false)

	return diffTyped(printer, want, got)
}

func diffd(want string, got string) string {
	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(want),
		B:        difflib.SplitLines(got),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})

	return diff

}

// formatStartingWhitespace formats leading whitespace characters to be visible while maintaining proper spacing
// Example:
//
//	Input:  "    \t  hello"
//	Output: "····→···hello"
//
// Where:
//
//	· represents a space (Middle Dot U+00B7)
//	→ represents a tab (Rightwards Arrow U+2192)
func formatStartingWhitespace(s string, colord *color.Color) string {
	out := " |"
	for j, char := range s {
		switch char {
		case ' ':
			out += "⌷" // Middle Dot (U+00B7)
		case '\t':
			out += "└──▹" // Rightwards Arrow (U+2192)
		default:
			return color.New(color.Faint).Sprint(out) + colord.Sprint(s[j:])
		}
	}
	return color.New(color.Faint).Sprint(out)
}

func diffTyped[T any](printer *pp.PrettyPrinter, want T, got T) string {
	// Enable colors

	var abc string

	switch any(want).(type) {
	case reflect.Type:
		want := ConvolutedFormatReflectType(any(want).(reflect.Type))
		got := ConvolutedFormatReflectType(any(got).(reflect.Type))
		return diffTyped[string](printer, want, got)
	case reflect.Value:
		w := any(want).(reflect.Value)
		g := any(got).(reflect.Value)
		want := ConvolutedFormatReflectValue(w)
		got := ConvolutedFormatReflectValue(g)
		return diffTyped[any](printer, want, got)
	case string:
		abc = diffd(any(want).(string), any(got).(string))
	default:
		abc = diffd(printer.Sprint(want), printer.Sprint(got))
	}
	if abc == "" {
		return ""
	}

	prevNoColor := color.NoColor
	defer func() {
		color.NoColor = prevNoColor
	}()
	color.NoColor = false

	expectedPrefix := fmt.Sprintf("[%s] %s", color.New(color.FgRed, color.Bold).Sprint("exp"), color.New(color.Faint).Sprint(" -"))
	actualPrefix := fmt.Sprintf("[%s] %s", color.New(color.Bold, color.FgBlue).Sprint("act"), color.New(color.Faint).Sprint(" +"))

	abc = strings.ReplaceAll(abc, "--- Expected", fmt.Sprintf("%s %s [%s]", color.New(color.Faint).Sprint("---"), color.New(color.FgRed).Sprint("expected"), color.New(color.FgRed, color.Bold).Sprint("exp")))
	abc = strings.ReplaceAll(abc, "+++ Actual", fmt.Sprintf("%s %s [%s]", color.New(color.Faint).Sprint("+++"), color.New(color.FgBlue).Sprint("actual"), color.New(color.FgBlue, color.Bold).Sprint("act")))

	realignmain := []string{}
	for i, spltz := range strings.Split(abc, "\n@@") {
		if i == 0 {
			realignmain = append(realignmain, spltz)
		} else {
			first := ""

			realign := []string{}
			for j, found := range strings.Split(spltz, "\n") {
				if j == 0 {
					first = color.New(color.Faint).Sprint("@@" + found)
				} else {
					if strings.HasPrefix(found, "-") {
						realign = append(realign, expectedPrefix+formatStartingWhitespace(found[1:], color.New(color.FgRed)))
					} else if strings.HasPrefix(found, "+") {
						realign = append(realign, actualPrefix+formatStartingWhitespace(found[1:], color.New(color.FgBlue)))
					} else {
						realign = append(realign, strings.Repeat(" ", 8)+formatStartingWhitespace(found, color.New(color.Faint)))
					}
				}
			}

			realignmain = append(realignmain, first)
			realignmain = append(realignmain, realign...)
		}
	}
	str := "\n"
	str += strings.Join(realignmain, "\n")

	return str
}

func SingleLineStringDiff(want string, got string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(want, got, false)
	return dmp.DiffPrettyText(diffs)
}
