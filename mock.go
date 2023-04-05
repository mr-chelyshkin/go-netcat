package go_netcat

type logger struct{}

func newMockLogger() *logger {
	return &logger{}
}

func (l *logger) Println(...any) {}
