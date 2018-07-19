package infrastructures

import (
	"database/sql"
	"fmt"

	"github.com/aswinda/loket-backend-test/interfaces"
)

type MysqlHandler struct {
	Conn *sql.DB
}

type MysqlRow struct {
	Rows *sql.Rows
}

func (handler *MysqlHandler) Execute(statement string) (int64, error) {
	row, err := handler.Conn.Exec(statement)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id, err := row.LastInsertId()

	return id, err
}

func (handler *MysqlHandler) Query(statement string) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(MysqlRow), err
	}

	row := new(MysqlRow)
	row.Rows = rows

	return row, nil
}

func (r MysqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MysqlRow) Next() bool {
	return r.Rows.Next()
}
