package kafka

type FollowConsumer interface {
    Start() error
}

type StubFollowConsumer struct{}

func NewStubFollowConsumer() *StubFollowConsumer {
    return &StubFollowConsumer{}
}

func (c *StubFollowConsumer) Start() error { return nil }
