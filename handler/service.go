package handler

type Service interface {

}

type TestService struct {

}

func NewTestService()Service{
	return &TestService{}
}