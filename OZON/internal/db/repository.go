package db

type DB interface {
	AddURL(url, key string) error
	GetURL(key string) (string, error)
}
