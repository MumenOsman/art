package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func SingleDecode(input string) (string, error) {

	var result strings.Builder
	i := 0
	n := len(input)

	for i < n {
		if input[i] == '[' {
			start := i

			// 1. Find the closing bracket ']' (relative index)
			j := strings.Index(input[i+1:], "]")

			// --- MANDATORY FIX FOR UNBALANCED BRACKETS ---
			if j == -1 {
				return "", fmt.Errorf("unbalanced brackets: missing closing ']'")
			}
			// ---------------------------------------------

			j += i + 1 // j is now the absolute index of ']'

			content := input[start+1 : j]
			spaceIndex := strings.Index(content, " ")

			// 2. Error checks for space and content
			if spaceIndex == -1 || spaceIndex == 0 || spaceIndex == len(content)-1 {
				return "", fmt.Errorf("malformed arguments (space missing or pattern empty)")
			}

			countStr := content[:spaceIndex]
			charsToRepeat := content[spaceIndex+1:]

			// --- NEW MANDATORY CHECK ---
			// 4. Error: The pattern must not contain square brackets
			if strings.ContainsAny(charsToRepeat, "[]") {
				return "", fmt.Errorf("malformed arguments: pattern cannot contain brackets")
			}

			// 3. Error: Count is not a number
			count, err := strconv.Atoi(countStr)
			if err != nil || count < 0 {
				return "", fmt.Errorf("count is not a valid non-negative number")
			}

			// 4. Expansion
			result.WriteString(strings.Repeat(charsToRepeat, count))

			// Move index past the closing ']'
			i = j + 1
		} else if input[i] == ']' {
			// --- MANDATORY FIX: Check for an unexpected closing bracket ---
			// If we encounter a ']' character outside of a parsing block,
			// it means a prior starting '[' was missed, or it's simply unbalanced.
			return "", fmt.Errorf("unbalanced brackets: unexpected closing ']'")
			// -----------------------------------------------------------
		} else {
			// Normal character
			result.WriteByte(input[i])
			i++
		}
	}

	return result.String(), nil
}
