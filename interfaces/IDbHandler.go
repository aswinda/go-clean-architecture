package interfaces

type IDbHandler interface {
	Execute(statement string) (int64, error)
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
