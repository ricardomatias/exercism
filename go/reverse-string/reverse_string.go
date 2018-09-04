package reverse

import "strings"

func String(str string) (reversed string) {
	strArr := strings.Split(str, "")

	for index := len(strArr); index > 0; index-- {
		reversed += strArr[index-1]
	}

	return reversed
}
