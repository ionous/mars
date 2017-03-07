package blocks

import (
	"fmt"
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

func DataItemOrItem(data interface{}) (ret string) {
	v := r.ValueOf(data)
	if v.Kind() != r.Array && v.Kind() != r.Slice {
		ret = fmt.Sprint("unexpected data", v.Kind(), data)
	} else if cnt := v.Len(); cnt > 0 {
		strs := make([]string, cnt)
		for i := 0; i < cnt; i++ {
			el := v.Index(i)
			strs[i] = DataToString(el.Interface())
		}
		ret = strings.Join(strs, " or ")
	}
	return ret
}

func DataToString(data interface{}) (ret string) {
	// FIX: arrays of these???
	switch val := data.(type) {
	case string:
		ret = val
	case float64:
		ret = strconv.FormatFloat(val, 'g', -1, 64)
	case bool:
		ret = strconv.FormatBool(val)
	default:
		ret = fmt.Sprint(val)
	}
	return
}
