package Utilities

import (
	"fmt"
	"strconv"
)

func Convert(path string) string {
	var result string
	//var windowsResult string
	for i := 0; i < len(path); i++ {
		if path[i] == '%' {
			result += mapValues(path[i+1 : i+3])
			i += 2
		} else {
			result += string(path[i])
		}
		//windowsResult += string(result[i])
		//if result[i] == ' ' {
		//	windowsResult += "` "
		//} else {
		//	windowsResult += string(result[i])
		//}
	}
	return result
}

func mapValues(val string) string {
	var intVal, _ = strconv.Atoi(val)
	switch intVal {
	case 20:
		return " "
	case 21:
		return "!"
	case 22:
		return "\""
	case 23:
		return "#"
	case 24:
		return "$"
	case 25:
		return "%"
	case 26:
		return "&"
	case 27:
		return "'"
	case 28:
		return "("
	case 29:
		return ")"
	}
	return ""
}

func main() {

	fmt.Println(Convert("/testdir/txt1%20-%20Copy%20%283%29.txt"))

}
