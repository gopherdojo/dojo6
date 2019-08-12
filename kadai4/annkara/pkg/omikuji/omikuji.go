package omikuji

import "math/rand"

// Draw a omikuji
func Draw() (int, string) {

	me := rand.Intn(7)
	var unsei string
	switch me {
	case 6:
		unsei = "大吉"
	case 5, 4:
		unsei = "中吉"
	case 3, 2:
		unsei = "小吉"
	case 1:
		unsei = "凶"
	default:
		return Draw()
	}
	return me, unsei
}
