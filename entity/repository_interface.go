package entity

type OrderRepositoryInterface interface {
	Save(user *User) error
	GetTotal() (int, error)
}
