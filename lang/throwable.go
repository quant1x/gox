package lang

type Exception interface {
	error
	Code() int
}

