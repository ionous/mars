package rt

import "strings"

type StreamEnd string

func (reason StreamEnd) Error() string {
	return strings.Join([]string{"stream end", string(reason)}, ":")
}

type StreamExceeded string

func (reason StreamExceeded) Error() string {
	return strings.Join([]string{"stream exceeded", string(reason)}, ":")
}
