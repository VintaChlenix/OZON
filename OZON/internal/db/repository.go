package db

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type DB interface {
	AddURL(url, key string) error
	GetURL(key string) (string, error)
}
