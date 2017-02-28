package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
	"regexp"
	"strings"
)

func PascalSpaces(name string) string {
	re := regexp.MustCompile("([A-Z])")
	return strings.TrimSpace(re.ReplaceAllStringFunc(name, func(s string) string {
		return " " + strings.ToLower(s)
	}))
}

// func SlashPath(path, child string) (ret string) {
// 	return path + "/" + child
// }

// func LastPath(path string) (ret string) {
// 	parts := strings.Split(path, "/")
// 	if len(parts) > 0 {
// 		ret = parts[len(parts)-1]
// 	}
// 	return
// }

func MakeToken(name string) string {
	return "[" + name + "]"
}

func TokenizePhrase(phrase string) (pre, post, token string) {
	//"[Fragment]" -> "", "", "Fragment"
	//"the [Subject] uses" -> "the", "uses", "Subject"
	// "classes" -> "classes", "", ""
	first := strings.SplitN(phrase, "[", 2)
	pre = strings.TrimSpace(first[0])
	if len(first) > 1 {
		end := strings.SplitN(first[1], "]", 2)
		if len(end) > 1 {
			post = strings.TrimSpace(end[1])
		}
		if t := end[0]; len(t) > 0 {
			token = MakeToken(t)
		}
	}
	return
}
func Tokenize(p *inspect.ParamInfo) (pre, post, token string) {
	if ptr := p.Phrase; ptr != nil {
		pre, post, token = TokenizePhrase(*ptr)
	}
	if len(token) == 0 {
		token = MakeToken(PascalSpaces(p.Name))
	}
	return
}

func FormatString(data interface{}) (ret string, err error) {
	if data == nil {
		ret = "<blank>"
	} else if s, ok := data.(string); ok {
		ret = s
	} else {
		err = errutil.New("not a string")
	}
	return
}

// func Format(data interface{}) (ret string, err error) {
// 	// array of these???
// 	switch val := data.(type) {
// 	case string:
// 		ret = val
// 	case float64:
// 		ret = strconv.FormatFloat(val, 'g', -1, 64)
// 	case bool:
// 		ret = strconv.FormatBool(val)
// 	default:
// 		err = errutil.New("couldnt format data", data)
// 	}
// 	return
// }

//func Tag(tags ...string) string {
//	return strings.Join(tags, " ")
//}
