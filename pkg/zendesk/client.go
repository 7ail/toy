package zendesk

type messenger struct {}

func Messenger() *messenger {
	return &messenger{}
}
