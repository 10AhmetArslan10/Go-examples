package piscine

func StrRev(s string) string {
	liste := []rune(s)
	n := ""
	for i := len(s) - 1; i >= 0; i-- {
		n += string(liste[i])
	}
	return string(n)
}
