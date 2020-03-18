package common

type Error struct {
	info string
}

func (e *Error) Error() string {
	return e.info
}

func (e *Error) Base(err error) *Error {
	e.info += " | " + err.Error()
	return e
}

func NewError(info string) *Error {
	return &Error{
		info: info,
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
