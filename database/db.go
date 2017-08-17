package database

type SKUDataAccess interface {
	Insert(int, string, string, int) error
	Get(int) string
	Print() string
	Update(int, int) error
	Delete(int) error
}

//New initializes a pointer to a sql database
func New() SKUDataAccess {
	m := MyDb{db: openDatabaseConnection()}
	return &m
}
