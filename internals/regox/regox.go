package regox

const (
	MATCHED      = "Matched"
	UNMATCHED    = "Unmatched"
	SYNTAX_ERROR = "Syntax_error"
)

func Match(i, j int, pattern, text string) string {
	n, m := len(pattern), len(text)

	for i < n && j < m {
		switch pattern[i] {
		case '?':
			i++
			j++

		case '*':
			res := Match(i+1, j, pattern, text)
			if res != UNMATCHED {
				return res
			}

			j++

		case '[':
			matched := false
			negate := false

			i++ // skipping [

			if pattern[i] == text[j] {
				if pattern[i] == '!' {
					negate = true
				}
				matched = true
			}
			i++

			for i < n && pattern[i] != ']' {
				if pattern[i] == text[j] {
					matched = true
				}
				i++
			}

			if pattern[i] != ']' {
				return SYNTAX_ERROR
			}

			if negate {
				matched = !matched
			}

			if !matched {
				return UNMATCHED
			}

			i++
			j++

		default:
			if pattern[i] == text[j] {
				i++
				j++
			} else {
				return UNMATCHED
			}
		}
	}

	if i >= n && j < m {
		return UNMATCHED
	} else if i < n && j >= m {
		for i < n {
			if pattern[i] != '*' {
				return UNMATCHED
			}
			i++
		}
	}

	return MATCHED
}
