package rtm

type Repository interface {
	FetchChannel(channel string) (*Channel, error)
}

type TestRepository struct{}

func (repo *TestRepository) FetchChannel(channel string) (*Channel, error) {
	return NewChannel(channel), nil
}
