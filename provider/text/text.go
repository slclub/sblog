package text

import "regexp"
import "strings"

func DeScript(str string) (ret string) {
	str = de(str)
	re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	ret = re.ReplaceAllString(str, "")
	return
}

func DeHtml(src string) string {
	src = de(src)
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	return src
}

func de(str string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllStringFunc(str, strings.ToLower)

	return str
}
