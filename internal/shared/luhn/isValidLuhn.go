package luhn

func IsValidLuhn(number string) bool {
	if number == "" {
		return false
	}
	sum := 0
	alt := false

	for i := len(number) - 1; i >= 0; i-- {
		digit := number[i]
		if digit < '0' || digit > '9' {
			return false
		}

		n := int(digit - '0')
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}
	return sum%10 == 0
}
