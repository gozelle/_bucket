package bucket

type Message interface {
	String() string
}

type String string

func (p String) String() string {
	return string(p)
}
