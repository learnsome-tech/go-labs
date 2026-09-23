# Concurrent Systems Programming with Go — lesson m01l03 — Go Workspaces and Modules
# https://learnsome.tech/courses/go-course/watch?lesson=m01l03
# © LearnSome.tech
package check

import "fmt"

// Target is one thing we are asked to watch.
type Target struct {
	Name string
	Port int
}

// Describe renders a target for a human reader.
func Describe(t Target) string {
	return fmt.Sprintf("%s on port %d", t.Name, t.Port)
}

func internalNote() string {
	return "not visible outside this package"
}
