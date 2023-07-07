package times

import (
	"regexp"
)

// if validation is fail return to true
func TimeFormatValidator(str string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	submatchall := re.FindAllString(str, -1)
	if len(submatchall) == 0 {
		return true
	}
	return false
}
