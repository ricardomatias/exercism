package bob

import (
	"fmt"
	"strings"
	"regexp"
)

func IsYell(remark string) bool {
	reg, _ := regexp.Compile(`[^a-zA-Z]`)

	justWords := reg.ReplaceAllString(remark, "")

	if justWords == "" {
		return false
	}

	return strings.ToUpper(remark) == remark
}

func Hey(remark string) string {
	reg, _ := regexp.Compile(`\?$`)

	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	var isQuestion bool = reg.MatchString(remark);

	if IsYell(remark) {
		if isQuestion == true {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	}

	if isQuestion == true {
		return "Sure."
	}

	return "Whatever."
}
