package virloc

import (
	"fmt"
	"strings"
)

func calculateChecksum(data string) string {
	msg := data[:len(data)-4]

	var calc byte
	var calculatedChecksum string
	for r := 0; r < len(msg); r++ {
		if r > 0 && msg[r] == '*' && msg[r-1] == ';' {
			break
		}
		caracter := byte(msg[r])
		calc = calc ^ caracter
	}
	calculatedChecksum = fmt.Sprintf("%X", calc)
	return calculatedChecksum
}

func removeSpecialCharsAndSpaces(message string) string {
	rmvSp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(message, " ", ""), ">", " "), "<", "")
	return rmvSp
}
