package redis

import "social-networking-platform/auth-service/internal/domain"

type SessionRepository interface {
    Save(session domain.Session) error
    Delete(token string) error
}

type StubSessionRepository struct{}

func NewStubSessionRepository() *StubSessionRepository {
    return &StubSessionRepository{}
}

func (r *StubSessionRepository) Save(session domain.Session) error { return nil }
func (r *StubSessionRepository) Delete(token string) error         { return nil }
