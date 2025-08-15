package functions

func Atoi(s string) int {
	number := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			number = number*10 + (int(s[i] - '0'))
		} else {
			return 0
		}
	}
	return number
}
