package database

//SKUDataAccess is for defining different types of databases
type SKUDataAccess interface {
	Insert(int, string, string, int) error
	Get(int) string
	Print() string
	Update(int, int) error
	Delete(int) error
}

//NewSQL initializes a pointer to a sql database
func NewSQL() SKUDataAccess {
	sql := SQLdb{db: openDatabaseConnection()}
	return &sql
}

//NewInMemoryDB initializes a pointer to a local database
func NewInMemoryDB() SKUDataAccess {
	mem := MemDb{db: newConnection()}
	return &mem
}
