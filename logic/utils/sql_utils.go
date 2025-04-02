package utils

import "strings"

// SplitSQLStatements splits a SQL script into individual statements.
func SplitSQLStatements(script string) []string {
	statements := []string{}
	current := ""
	for _, line := range strings.Split(script, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "--") {
			// Skip empty lines and comments
			continue
		}
		current += trimmed + " "
		if strings.HasSuffix(trimmed, ";") {
			// End of a statement
			statements = append(statements, current)
			current = ""
		}
	}
	return statements
}
