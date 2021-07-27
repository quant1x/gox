package lang

type Throwable interface {
	error
	Code() int
}

type Exception struct {
	Throwable
	code    int
	message string
}

func New(code int, message string) *Exception {
	return &Exception{
		code:    code,
		message: message,
	}
}

func (this Exception) Error() string {
	return this.message
}

func (this Exception) Code() int {
	return this.code
}
