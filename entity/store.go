package entity

type Store interface {
	CreateUser(u UserRequest) (UserModel, error)
}

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) CreateUser(u UserRequest) (UserModel, error) {
	return s.Store.CreateUser(u)
}
