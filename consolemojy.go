package consolemojy

import (
	"errors"
	"html"
	"strconv"
	"strings"
)

func GetEmo(name string) (string, error) {
	var code string
	switch name {
	case "fire":
		code = "\\U0001F525"
	case "thumbs_up":
		code = "\\U0001F44D"
	case "thumbs_down":
		code = "\\U0001F44E"
	case "check_mark":
		code = "\\U00002705"
	case "x_mark":
		code = "\\U0000274C"
	case "warning":
		code = "\\U000026A0"
	default:
		return "", errors.New("enf - emoji not found")

	}
	h := strings.ReplaceAll(code, "\\U", "0x")
	i, _ := strconv.ParseInt(h, 0, 64)
	return html.UnescapeString(string(i)), nil
}
