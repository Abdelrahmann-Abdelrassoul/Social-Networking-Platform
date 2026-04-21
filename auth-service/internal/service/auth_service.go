package service

type AuthService interface {
    Login() error
    Callback() error
    Logout() error
}

type StubAuthService struct{}

func NewStubAuthService() *StubAuthService {
    return &StubAuthService{}
}

func (s *StubAuthService) Login() error    { return nil }
func (s *StubAuthService) Callback() error { return nil }
func (s *StubAuthService) Logout() error   { return nil }
