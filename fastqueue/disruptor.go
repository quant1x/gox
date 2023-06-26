package fastqueue

type Disruptor struct {
	Writer
	Reader
}

func createDisruptor(writer Writer, reader Reader) Disruptor {
	return Disruptor{
		Writer: writer,
		Reader: reader,
	}
}
