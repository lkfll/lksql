package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lkfll/lksql/server/dto"
	"github.com/lkfll/lksql/server/entity"
	"github.com/lkfll/lksql/server/result"
	serverutils "github.com/lkfll/lksql/server/serverUtils"
)

type DDL struct {
	DB *sql.DB
}

// 显示数据库中的表
const ShowTableDDL = "show tables;"

func (ddl *DDL) ShowTables(resp http.ResponseWriter, req *http.Request) {
	r, err := ddl.DB.Query(ShowTableDDL)
	if err != nil {
		fmt.Fprint(resp, result.Err(err))
		return
	}
	defer r.Close()
	ret := make([]string, 0)
	for r.Next() {
		var buf sql.RawBytes
		r.Scan(&buf)
		ret = append(ret, string(buf))
	}
	fmt.Fprint(resp, result.Ok(ret))
}

// ! 考虑废弃 二维数组前端不好渲染
// 查看多个表信息
func (ddl *DDL) DeseTables(resp http.ResponseWriter, req *http.Request) {
	var tables []string
	serverutils.BodyScanJson(req.Body, &tables)
	ret := make([]interface{}, 0)
	for _, v := range tables {
		s := entity.FieldConfigRepository.Select()
		r, err := s.GetResultBySql(ddl.DB, fmt.Sprintf(DeseDDL, v))
		if err != nil {
			fmt.Fprint(resp, result.Err(err))
			return
		}
		defer r.Close()
		i, err := s.Scan(r)
		if err != nil {
			fmt.Fprint(resp, result.Err(err))
			return
		}
		ret = append(ret, i)
	}
	fmt.Fprint(resp, result.Ok(ret))
}

// 查看表信息
const DeseDDL = "desc %s;"

func (ddl *DDL) DeseTable(resp http.ResponseWriter, req *http.Request) {
	var t dto.Table
	serverutils.BodyScanJson(req.Body, &t)
	s := entity.FieldConfigRepository.Select()
	r, err := s.GetResultBySql(ddl.DB, fmt.Sprintf(DeseDDL, t.Table))
	if err != nil {
		fmt.Fprint(resp, result.Err(err))
		return
	}
	defer r.Close()
	i, err := s.Scan(r)
	if err != nil {
		fmt.Fprint(resp, result.Err(err))
		return
	}
	fmt.Fprint(resp, result.Ok(i))
}

// 创建表
const CreateTableDDL = `CREATE TABLE %s(
	id INT UNSIGNED PRIMARY KEY ,
	delete_time int,
	update_time int,
	create_time int
);`

func (ddl *DDL) CreateTable(resp http.ResponseWriter, req *http.Request) {
	var t dto.Table
	serverutils.BodyScanJson(req.Body, &t)
	_, err := ddl.DB.Exec(fmt.Sprintf(CreateTableDDL, t.Table))
	if err != nil {
		fmt.Fprint(resp, result.Err(err))
		return
	}
	fmt.Fprint(resp, result.Ok("ok"))
}

// 删除表
const DropTableDDL string = "drop table %s;"

func (ddl *DDL) DropTable(resp http.ResponseWriter, req *http.Request) {
	var t dto.Table
	serverutils.BodyScanJson(req.Body, &t)
	_, err := ddl.DB.Exec(fmt.Sprintf(DropTableDDL, t.Table))
	if err != nil {
		fmt.Fprint(resp, result.Err(err))
		return
	}
	fmt.Fprint(resp, result.Ok("ok"))
}
