package db

type InMemoryDB struct{
	data map[string]string
}

func (db *InMemoryDB) AddURL(url, key string) error{
	db.data[key]=url
	return nil
}

func (db *InMemoryDB) GetURL(key string) (string, error){
	return db.data[key], nil
}

func NewInMemory() *InMemoryDB{
	db := InMemoryDB{data: map[string]string{}}
	return &db
}
