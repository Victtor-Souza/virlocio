package virloc

import (
	"fmt"
	"math/rand"
	"strconv"
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

func formatDate(datestring string) string {
	var (
		d string
		m string
		a string
	)
	fmt.Sscanf(datestring, "%2s%2s%2s", &d, &m, &a)

	return fmt.Sprintf("20%s-%s-%s", a, m, d)

}

func formatTime(timestring string) string {
	var (
		h string
		m string
		s string
	)
	fmt.Sscanf(timestring, "%2s%2s%2s", &h, &m, &s)

	return fmt.Sprintf("%s:%s:%s", h, m, s)
}

func getonoff(propertyvalue, off, on string) string {
	if propertyvalue == on {
		return "ON"
	}

	return "OFF"
}

func convertStringToFloat64(value string, decimalPlaces int64) float64 {
	newstr := []string{}
	arr := strings.Split(value, "")

	newstr = append(newstr, arr[:decimalPlaces+1]...)

	newstr = append(newstr, ".")
	newstr = append(newstr, arr[decimalPlaces+1:]...)

	result, _ := strconv.ParseFloat(strings.Join(newstr, ""), 64)

	return result
}

func convertStringToInt64(value string) int64 {
	v, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		fmt.Printf("PARSE ERROR - %s\n", err.Error())
		return 0
	}

	return v
}

func asBits(val uint64) []uint64 {
	bits := []uint64{}
	for i := 0; i < 8; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}
