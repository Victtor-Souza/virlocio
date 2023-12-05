package virloc

import (
	"fmt"
	"math/rand"
	"strings"
)

func CalculateChecksum(data string) string {

	var calc byte
	var calculatedChecksum string
	for r := 0; r < len(data); r++ {
		if r > 0 && data[r] == '*' && data[r-1] == ';' {
			break
		}
		caracter := byte(data[r])
		calc = calc ^ caracter
	}
	calculatedChecksum = fmt.Sprintf("%X", calc)
	return calculatedChecksum
}

func GetMessageNumber() string {
	min := 32768
	max := 65535
	messageNumber := rand.Intn(max-min) + min

	return fmt.Sprintf("%X", messageNumber)
}

func removeSpecialCharsAndSpaces(message string) string {
	rmvSp := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(message, " ", ""), ">", " "), "<", ""))
	return rmvSp
}

func removeDeviceData(message string) string {
	arr := strings.Split(message, ";")
	return arr[0]
}
