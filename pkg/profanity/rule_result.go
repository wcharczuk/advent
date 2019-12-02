package profanity

import (
	"fmt"
	"strings"
)

// RuleResult is a result from a rule.
type RuleResult struct {
	OK      bool
	File    string
	Line    int
	Message string
	Err     error
}

// Failure returns a failure error message for a given file and error.
func (r RuleResult) Failure(rule Rule) error {
	var tokens []string
	tokens = append(tokens, fmt.Sprintf("%s:%d", r.File, r.Line))
	if rule.ID != "" {
		tokens = append(tokens, fmt.Sprintf("\t%s: %s", AnsiLightBlack("id"), rule.ID))
	}
	if rule.Description != "" {
		tokens = append(tokens, fmt.Sprintf("\t%s: %s", AnsiLightBlack("description"), rule.Description))
	}
	tokens = append(tokens, fmt.Sprintf("\t%s: %s", AnsiLightBlack("status"), AnsiRed("failed")))
	tokens = append(tokens, fmt.Sprintf("\t%s: %s", AnsiLightBlack("rule"), r.Message))
	return fmt.Errorf(strings.Join(tokens, "\n"))
}
