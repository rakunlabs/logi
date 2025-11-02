package logi

import "strconv"

// UnescapeJSONInLine processes a line of output, unescaping JSON strings
// while preserving ANSI color codes
func UnescapeJSONInLine(line []byte) []byte {
	if len(line) == 0 {
		return line
	}

	var result []byte
	i := 0

	for i < len(line) {
		// Check for ANSI escape sequence
		if line[i] == '\x1b' && i+1 < len(line) && line[i+1] == '[' {
			// Find the end of the ANSI sequence
			j := i + 2
			for j < len(line) && line[j] != 'm' {
				j++
			}
			if j < len(line) {
				j++ // include the 'm'
			}
			// Copy the ANSI sequence as-is
			result = append(result, line[i:j]...)
			i = j
			continue
		}

		// Check for quoted string that might contain escaped JSON: "{...}" or "[...]"
		if line[i] == '"' && i+2 < len(line) && line[i+1] == '{' ||
			(line[i] == '"' && i+2 < len(line) && line[i+1] == '[') {
			// This looks like it might be a JSON string
			// Try to find the closing quote
			j := i + 1
			escaped := false
			foundJSON := false

			// Check if this string contains escaped quotes (indicating escaped JSON)
			for j < len(line) {
				if escaped {
					escaped = false
					j++
					continue
				}

				if line[j] == '\\' {
					// Check if next char is a quote
					if j+1 < len(line) && line[j+1] == '"' {
						foundJSON = true
					}
					escaped = true
					j++
					continue
				}

				if line[j] == '"' {
					// Found closing quote
					j++ // include the closing quote
					break
				}
				j++
			}

			// If we found escaped quotes, try to unescape this as a JSON string
			if foundJSON && j > i+1 {
				quotedStr := string(line[i:j])
				// Try to unquote it
				unquoted, err := strconv.Unquote(quotedStr)
				if err == nil {
					// Successfully unescaped
					result = append(result, []byte(unquoted)...)
					i = j
					continue
				}
			}

			// If unescaping failed or no escaped JSON found, fall through
		}

		// Default: copy the current byte
		result = append(result, line[i])
		i++
	}

	return result
}
