package functions

func IsPrintableRange(nameORmsg string) bool {
	if nameORmsg == "" {
		return false
	}
	for _, v := range nameORmsg {
		if v < 32 || v > 126 {
			return false
		}
	}
	return true
}
