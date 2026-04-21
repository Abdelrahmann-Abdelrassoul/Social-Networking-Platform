package postgres

import "social-networking-platform/posts-service/internal/domain"

type PostRepository interface {
    Save(post domain.Post) error
    GetByID(id string) (*domain.Post, error)
    Update(post domain.Post) error
    Delete(id string) error
}

type StubPostRepository struct{}

func NewStubPostRepository() *StubPostRepository {
    return &StubPostRepository{}
}

func (r *StubPostRepository) Save(post domain.Post) error            { return nil }
func (r *StubPostRepository) GetByID(id string) (*domain.Post, error) { return nil, nil }
func (r *StubPostRepository) Update(post domain.Post) error          { return nil }
func (r *StubPostRepository) Delete(id string) error                 { return nil }
