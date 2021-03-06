package runner

import "strings"
import "github.com/dpastoor/nonmemutils/parser"

//PrepareForExecution parses and prepares strings from a file for execution in a different context
// for example, replacing the $DATA path
func PrepareForExecution(s []string) []string {
	for i, line := range s {
		if strings.Contains(line, "$DATA") {
			s[i] = parser.AddPathLevelToData(line)
		}
	}
	return s
}
