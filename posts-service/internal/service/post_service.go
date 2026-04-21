package service

type PostService interface {
    CreatePost() error
    GetPost() error
    UpdatePost() error
    DeletePost() error
}

type StubPostService struct{}

func NewStubPostService() *StubPostService {
    return &StubPostService{}
}

func (s *StubPostService) CreatePost() error { return nil }
func (s *StubPostService) GetPost() error    { return nil }
func (s *StubPostService) UpdatePost() error { return nil }
func (s *StubPostService) DeletePost() error { return nil }
