package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"regexp"
	"strings"
)

func Spaces(s ...string) string {
	return strings.Join(s, " ")
}

func PascalSpaces(name string) string {
	re := regexp.MustCompile("([A-Z])")
	return strings.TrimSpace(re.ReplaceAllStringFunc(name, func(s string) string {
		return " " + strings.ToLower(s)
	}))
}

func MakeToken(name string) string {
	return "[" + name + "]"
}

const NewLineString = string('\n')

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

// }

//func Tag(tags ...string) string {
//	return strings.Join(tags, " ")
//}
