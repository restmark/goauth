package mock

type mock struct{}

func New() *mock {
	return &mock{}
}
