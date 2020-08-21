package concat

import "strings"

//Concat
func Build(strs ...interface{}) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(buildStr(str))
	}
	return sb.String()
}
