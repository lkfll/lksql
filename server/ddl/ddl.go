package ddl

import (
	"database/sql"
	"fmt"
)

// r, err := ddl.DB.Query(fmt.Sprintf(DeseDDL, table))
// 	if err != nil {
// 		return nil, err
// 	}
// 	ct, _ := r.ColumnTypes()
// 	for _, v := range ct {
// 		fmt.Printf("v.DatabaseTypeName(): %v\n", v.DatabaseTypeName())
// 	}
// 	fmt.Println(r.Columns())

type DDL struct {
	DB sql.DB
}

// 显示数据库中的表
const ShowTableDDL = "show tables;"

func (ddl *DDL) ShowTables() ([]string, error) {
	// lksql.
	r, err := ddl.DB.Query(ShowTableDDL)
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

// 查看表
const DeseDDL = "desc %s;"

func (ddl *DDL) DeseTable(table string) ([]string, error) {

	// defer r.Close()
	return nil, nil
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
