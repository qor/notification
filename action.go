package notification

type ActionArgument struct {
	Message  *Message
	Argument interface{}
}

type Action struct {
	Name     string
	Method   string
	URL      string
	Resource *Resource
	Handle   func(*ActionArgument) error
}
