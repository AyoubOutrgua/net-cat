package functions

func IsValidUsername(username string) bool {
	if len(username) > 15 {
		return false
	}
	mutexClient.Lock()
	for _, client := range clients {
		if username == client {
			mutexClient.Unlock()
			return false
		}
	}
	mutexClient.Unlock()
	return true
}
