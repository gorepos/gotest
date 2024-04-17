package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Styles struct {
	Run,
	Skip,
	PackageStart,
	PackageResult,
	_ *color.Color
}

var DefaultStyles = Styles{
	Run:           color.New(color.FgHiBlack),
	Skip:          color.New(color.FgHiYellow),
	PackageStart:  color.New(color.Bold, color.FgHiBlue),
	PackageResult: color.New(color.FgBlue),
}

// Transformer is a struct that holds the I/O streams for the transformer, and a PASS flag
type Transformer struct {
	Stdin  io.Reader
	Stdout io.Writer
	PASS   bool
}

// NewTransformer creates a new Transformer with the default I/O streams
func NewTransformer() *Transformer {
	return &Transformer{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
	}
}

// Execute runs the 'go test -json <user-args>' command
// and transforms the output into human-readable form.
// If there are any errors, they are printed to Stdout.
// If there are any test failures, PASS is set to false.
func (t *Transformer) Execute(userArgs []string) {
	args := []string{"test", "-json", "-p", "1"}
	args = append(args, userArgs...)
	cmd := exec.Command("go", args...)
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		t.println(cmd.Args, err)
		return
	}
	cmd.Stderr = t.Stdout
	if err := cmd.Start(); err != nil {
		t.println(cmd.Args, err)
		return
	}
	t.Stdin = outPipe
	t.Transform()
	if err := cmd.Wait(); err != nil {
		t.PASS = false
		t.println(cmd.Args, err)
		return
	}
}

// print
func (t *Transformer) print(a ...interface{}) {
	// if first arg is a color, use it
	if len(a) > 0 {
		if c, ok := a[0].(*color.Color); ok {
			c.Fprint(t.Stdout, a[1:]...)
			return
		}
	}
	fmt.Fprint(t.Stdout, a...)
}

// println
func (t *Transformer) println(a ...interface{}) {
	if len(a) > 0 {
		if c, ok := a[0].(*color.Color); ok {
			c.Fprintln(t.Stdout, a[1:]...)
			return
		}
	}
	fmt.Fprintln(t.Stdout, a...)
}

// Transform reads the JSON output from Stdin, and prints it in a human-readable form
func (t *Transformer) Transform() {
	t.PASS = true
	scanner := bufio.NewScanner(t.Stdin)

	currentPackage := ""

	for scanner.Scan() {
		text := scanner.Text()
		item, err := ParseLine(text)
		if err != nil {
			t.PASS = false
			t.println(err)
			return
		}
		switch {
		case item.Action == ActionSkip:
			if item.Test != "" {
				cl := DefaultStyles.Skip
				t.print(cl, item.indent()+" ⏎ "+item.PrettyName(), " (skipped)\n")
			}

		case item.Action == ActionStart:
			//currentPackage = item.Package

		case item.Action == ActionRun:
			if currentPackage != item.Package {
				// start a new package
				t.println()
				cl := DefaultStyles.PackageStart
				t.println(cl, item.Package+"...")
				currentPackage = item.Package
			}
			cl := DefaultStyles.Run
			t.print(cl, item.indent()+"   "+item.PrettyName())
			t.println(cl, "...")

		case item.IsPackageResult():
			cl := DefaultStyles.PackageResult
			//t.print(item.indent() + " " + item.icon())
			//t.print(cl, item.indent()+" "+item.Package)
			t.print(cl, item.Package, " ", item.icon())

			t.print(fmt.Sprintf(" (%.2fs)\n", item.Elapsed))

		case item.Action == ActionOutput:

			// skip non-informative output
			if item.Output == "" {
				continue
			}
			if strings.HasPrefix(item.Output, "===") {
				continue
			}
			if strings.HasPrefix(item.Output, "---") {
				continue
			}
			if strings.Contains(item.Output, "[no test files]") &&
				strings.HasPrefix(item.Output, "?") {
				continue
			}
			if strings.HasPrefix(item.Output, "ok  \t"+item.Package) {
				continue
			}
			if strings.TrimSpace(item.Output) == "PASS" {
				continue
			}
			if strings.TrimSpace(item.Output) == "FAIL" {
				continue
			}

			t.print(item.Output)

		case item.IsTestResult():

			t.println(item.String())

			if item.Action == ActionFail {
				t.PASS = false
			}
		default:
			//fmt.Println("default", text)
		}
	}
}
