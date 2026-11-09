package main

// failure is an error with an exit status; run turns it into the error text.
type failure struct {
	code int
	msg  string
}

func (f *failure) Error() string { return f.msg }

func fail(code int, msg string) *failure { return &failure{code: code, msg: msg} }
