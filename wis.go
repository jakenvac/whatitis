package whatitis

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

// Comparator defines a way to return boolean based on two or more inputs
type Comparator func(result interface{}, expected interface{}) (success bool, message string)

type WhatContext struct {
	t             *testing.T
	hasPassed     bool
	itCount       int
	itCountPassed int
}

func newWhatContext(t *testing.T) (context *WhatContext) {
	context = &WhatContext{
		t:         t,
		hasPassed: true,
	}
	return
}

type ItContext struct {
	t         *testing.T
	hasPassed bool
}

func newItContext(t *testing.T) (context *ItContext) {
	context = &ItContext{
		t:         t,
		hasPassed: true,
	}
	return
}

// It defines a test on a specific piece of functionality
func (w *WhatContext) It(description string, f func(i *ItContext)) {
	w.t.Helper()
	w.itCount++
	fmt.Print("\t\t" + description + " ")
	context := newItContext(w.t)
	f(context)
	if context.hasPassed {
		w.itCountPassed++
	}
	fmt.Println()
}

// Is asserts that results are expected
func (i *ItContext) Is(result interface{}, comparator Comparator, expected interface{}) {
	i.t.Helper()
	testDidPass, message := comparator(result, expected)
	messageChar := "✓ "
	messageColor := color.FgGreen
	if testDidPass {
		i.hasPassed = i.hasPassed && true
	} else {
		messageChar = "✗ "
		messageColor = color.FgRed
		i.hasPassed = i.hasPassed && false
		i.t.Error(message)
	}
	c := color.New(messageColor, color.Bold)
	c.Print(messageChar)
}

// What defines what is being tested overall
func What(t *testing.T, description string, f func(w *WhatContext)) {
	t.Helper()
	context := newWhatContext(t)
	c := color.New(color.Bold, color.FgCyan)
	c.Println("\n\t" + description)
	f(context)
	var status string
	msgColor := color.FgGreen
	if context.itCount == context.itCountPassed {
		status = "PASS"
	} else {
		msgColor = color.FgRed
		status = "FAIL"
	}
	c = color.New(msgColor)
	c.Printf("\t%s %d/%d\n\n", status, context.itCountPassed, context.itCount)
}

// EqualTo compares the Equality between two items
var EqualTo = func(result interface{}, expected interface{}) (success bool, message string) {
	success = result == expected
	message = fmt.Sprintf("Expected '%v' but Got '%v'", expected, result)
	return
}

// NotEqualTo to compares the inequality between two items
var NotEqualTo = func(result interface{}, expected interface{}) (success bool, message string) {
	success = result != expected
	message = fmt.Sprintf("Expected '%v' to differ from '%v'", expected, result)
	return
}
