package server

import (
	"database/sql"
	"fmt"
)

type DDL struct {
	DB sql.DB
}

// 显示数据库中的表
const ShowTableDDL = "show tables;"

func (ddl *DDL) ShowTables() (sql.Result, error) {
	return ddl.DB.Exec(ShowTableDDL)
}

// 查看表
const DeseDDL = "desc %s;"

func (ddl *DDL) DeseTable(table string) (sql.Result, error) {
	return ddl.DB.Exec(fmt.Sprintf(DeseDDL, table))

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
