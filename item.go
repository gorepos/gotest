package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func ParseLine(line string) (Item, error) {
	event := Item{}
	err := json.Unmarshal([]byte(line), &event)
	if err != nil {
		return Item{}, fmt.Errorf("parsing JSON: %w\ninput: %s", err, line)
	}
	return event, nil
}

const (
	ActionPass   = "pass"
	ActionFail   = "fail"
	ActionSkip   = "skip"
	ActionStart  = "start"
	ActionRun    = "run"
	ActionOutput = "output"
)

// Item represents a single test event
type Item struct {
	Action  string
	Package string
	Test    string
	Output  string
	Elapsed float64
}

// String formats a test Item for display
func (i Item) String() string {
	return fmt.Sprintf(i.indent()+" %s %s (%.2fs)", i.icon(), i.PrettyName(), i.Elapsed)
}

func (i Item) PrettyName() string {
	if i.Test == "" {
		return ""
	}

	parts := strings.Split(i.Test, "/")

	// Remove "Test" prefix
	parts[0] = strings.TrimPrefix(parts[0], "Test")

	// Convert camel case to human readable
	for i := 0; i < len(parts); i++ {
		// If it's a file name or other identifier, underline it
		if strings.Contains(parts[i], ".") {
			cl := color.New(color.Underline)
			parts[i] = cl.Sprint(parts[i])
		} else {
			// Otherwise, convert to human-readable format
			parts[i] = convertCamelCase(parts[i])
		}
	}
	return strings.Join(parts, " ")

}

func (i Item) indent() string {
	depth := strings.Count(i.Test, "/")
	return strings.Repeat("   ", depth)
}

func (i Item) icon() string {
	icon := color.RedString("x")
	if i.Action == ActionPass {
		icon = color.GreenString("✔")
	}
	return icon
}

func (i Item) IsTestResult() bool {
	if i.Test == "" {
		return false
	}
	if i.Action == ActionPass || i.Action == ActionFail {
		return true
	}
	return false
}

func (i Item) IsPackageResult() bool {
	if i.Test != "" {
		return false
	}
	if i.Action == ActionPass || i.Action == ActionFail {
		return true
	}
	return false
}
