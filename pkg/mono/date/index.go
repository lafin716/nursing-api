package date

type DateClient struct {
	Parser
	Formatter
}

func New() *DateClient {
	return &DateClient{
		NewParser(),
		NewFormatter(),
	}
}
