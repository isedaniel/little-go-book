package main

func NewMessages(error, message string) *messages {
	return &messages{
		error:   error,
		message: message,
	}
}

type messages struct {
	error   string
	message string
}

func (m *messages) Bar() {
	m.message += " (this passed through Bar)"
}

type newerror struct {
	m     *messages
	error string
}

func (n *newerror) Baz() {
	n.m.Bar()
	n.error = n.m.message
}
