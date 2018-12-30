package gedcom5

var IsAt = IsExact(0x40)
var IsSpace = IsExact(0x20)

func IsEither(checkers ...ByteTypeChecker) ByteTypeChecker {
	return func(b byte) bool {
		for _, check := range checkers {
			if check(b) {
				return true
			}
		}
		return false
	}
}

func IsDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func IsAlpha(b byte) bool {
	return (b >= 0x41 && b <= 0x5A) || (b >= 0x61 && b <= 0x7A) || b == 0x5F
}

func IsExact(b byte) ByteTypeChecker {
	return func(other byte) bool {
		return other == b
	}
}

func IsAlphaNum(b byte) bool {
	return IsDigit(b) || IsAlpha(b)
}
