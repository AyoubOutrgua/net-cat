package functions

func IsValidUsername(username string) bool {
	if len(username) > 15 {
		return false
	}
	mutex.Lock()
	for _, client := range clients {
		if username == client {
			return false
		}
	}
	mutex.Unlock()
	return true
}
