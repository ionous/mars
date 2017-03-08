package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	r "reflect"
	"regexp"
	"strconv"
	"strings"
)

func Spaces(s ...string) string {
	return strings.Join(s, " ")
}

func Lines(s ...string) []string {
	return s
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

// DataItemOrItem, join a list with "or".
func DataItemOrItem(data interface{}) string {
	return JoinList(r.ValueOf(data), " or ")
}

func DataItemAndItem(data interface{}) string {
	return JoinList(r.ValueOf(data), " and ")
}

func JoinList(list r.Value, sep string) (ret string) {
	if k := list.Kind(); k != r.Slice {
		ret = Spaces("<unexpected", k.String(), "items>")
	} else if cnt := list.Len(); cnt == 0 {
		ret = "<empty>"
	} else {
		strs := make([]string, cnt)
		for i := 0; i < cnt; i++ {
			v := list.Index(i)
			strs[i] = DataToString(v.Interface())
		}
		ret = strings.Join(strs, sep)
	}
	return
}

// DataToString, format script values in a simple way.
func DataToString(data interface{}) (ret string) {
	if data == nil {
		ret = "<blank>"
	} else {
		v := r.ValueOf(data)
		switch k := v.Kind(); k {
		case r.String:
			ret = v.String()
		case r.Float64:
			val := v.Float()
			ret = strconv.FormatFloat(val, 'g', -1, 64)
		case r.Bool:
			val := v.Bool()
			ret = strconv.FormatBool(val)
		case r.Slice:
			ret = JoinList(v, ", ")
		default:
			ret = Spaces("<unknown", k.String(), "data>")
		}
	}
	return
}
