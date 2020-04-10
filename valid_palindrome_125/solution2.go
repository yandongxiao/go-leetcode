package valid_palindrome_125

func isPalindrome(s string) bool {
	begin := 0
	end := len(s) - 1
	for begin < end {
		var b bool
		var x byte

		for begin <= end {
			if x, b = valid(s[begin]); b {
				break
			}
			begin++
		}

		var y byte
		for begin <= end {
			if y, b = valid(s[end]); b {
				break
			}
			end--
		}

		if x != y {
			return false
		}
		begin++
		end--
	}
	return true
}

func valid(b byte) (byte, bool) {
	if b >= '0' && b <= '9' {
		return b, true
	}
	if b >= 'a' && b <= 'z' {
		return b, true
	}

	if b >= 'A' && b <= 'Z' {
		return 'a' + (b - 'A'), true
	}

	return 0, false
}
