package rt

type StreamEnd string

func (reason StreamEnd) Error() string {
	return "stream end" + string(reason)
}

type StreamExceeded string

func (reason StreamExceeded) Error() string {
	return "stream exceeded" + string(reason)
}
