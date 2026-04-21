package kafka

import "social-networking-platform/posts-service/internal/domain"

type PostProducer interface {
    PublishCreated(post domain.Post) error
}

type StubPostProducer struct{}

func NewStubPostProducer() *StubPostProducer {
    return &StubPostProducer{}
}

func (p *StubPostProducer) PublishCreated(post domain.Post) error { return nil }
