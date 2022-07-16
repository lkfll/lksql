package server

import (
	"database/sql"
	"fmt"
)

type DDL struct {
	DB sql.DB
}

// 扫描成string切片
func scanOfStringlist(r *sql.Rows, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	defer r.Close()
	ret := make([]string, 0)
	for r.Next() {
		var buf sql.RawBytes
		r.Scan(&buf)
		ret = append(ret, string(buf))
	}
	return ret, nil
}

// 显示数据库中的表
const ShowTableDDL = "show tables;"

func (ddl *DDL) ShowTables() ([]string, error) {
	return scanOfStringlist(ddl.DB.Query(ShowTableDDL))
}

// 查看表
const DeseDDL = "desc %s;"

func (ddl *DDL) DeseTable(table string) ([]string, error) {
	return scanOfStringlist(ddl.DB.Query(fmt.Sprintf(DeseDDL, table)))
}

// 创建表
const CreateTableDDL = `CREATE TABLE %s(
	id INT UNSIGNED PRIMARY KEY ,
	delete_time int,
	update_time int,
	create_time int
);`

func (ddl *DDL) CreateTable(table string) (sql.Result, error) {
	return ddl.DB.Exec(fmt.Sprintf(CreateTableDDL, table))
}

// 删除表
const DropTableDDL string = "drop table %s;"

func (ddl *DDL) DropTable(table string) (sql.Result, error) {
	return ddl.DB.Exec(fmt.Sprintf(DropTableDDL, table))
}
