package repository

import (
	"fmt"
	"reflect"

	"github.com/lkfll/lksql/analyze"
	basesql "github.com/lkfll/lksql/baseSql"
	"github.com/lkfll/lksql/exec"
)

// 实体类仓库
type EntityRepository struct {
	*Repository
	EntitySql *basesql.EntitySql // 生成的基础sql语句
}

//
func EntityStart(repository *Repository, entitySql *basesql.EntitySql) *EntityRepository {
	var ret EntityRepository
	ret.Repository = repository
	ret.EntitySql = entitySql
	return &ret
}

// 按表格删除
// DELETE FROM user WHERE 1=1;
func (Repository *EntityRepository) Delete() func(string) *exec.Delete {
	return exec.NewDelete(Repository.EntitySql.Delete.Sql)
}

// 按表格插入
// values 传入参数使用map 使用map代替对象
// INSERT INTO user( user_name  , password  , role ) VALUES  ( '小明' , '16191411qin' , 0  ) , ( '小明' , '16191411qin' , 0  )
func (Repository *EntityRepository) Insert(params ...analyze.Type) *exec.Insert {
	sql := Repository.EntitySql.Insert // 表对应的sql

	value := "(" // 拼接增加的一行value
	for i := 0; i < len(Repository.EntitySql.Insert.InsertFields); i++ {
		value = fmt.Sprintf("%s %s,", value, "?")
	}
	value = fmt.Sprintf("%s %s", value[:len(value)-1], ")") // 一行value

	var valuesSql string = ""
	ParamList := make([]interface{}, 0) // 参数切片
	for i := 0; i < len(params); i++ {  // 传入参数的长度
		valuesSql = fmt.Sprintf("%s %s,", valuesSql, value) // 拼接sql
		param := reflect.ValueOf(params[i])                 // 获取一行参数value反射
		for _, v := range Repository.EntitySql.Insert.InsertFields {
			// 通过反射获取参数列表
			ParamList = append(ParamList, param.FieldByName(v.Name).Interface())
		}
	}
	tempsql := fmt.Sprintf(sql.Sql, valuesSql[:len(valuesSql)-1])

	return exec.NewInsert(tempsql, ParamList...)
}

// 查询全部
// SELECT  issues.id , issues.issues_name , issues.label , issues.description , issues.uid , priority.priority_name
// FROM  issues
func (repository *EntityRepository) Select() *exec.Query {
	return exec.NewQuery(repository.EntitySql.SelectSql, exec.NewScan(repository.Type, repository.SelectField...))
}

// 按表格修改
// UPDATE user SET  user_name='迪迦奥特曼' , password='16191411qin' , role=0  WHERE id = 1 and id = 2;
func (Repository *EntityRepository) Update(param analyze.Type) func(string) *exec.Update {
	sql := Repository.EntitySql.Update // 表对应的sql

	fields := Repository.EntitySql.Update.UpdateFields // 表的字段

	ParamList := make([]interface{}, 0)  // 参数切片
	paramValue := reflect.ValueOf(param) // 获取一行参数value反射
	for _, v := range fields {
		// 通过反射获取参数列表
		ParamList = append(ParamList, paramValue.FieldByName(v.Name).Interface())
	}

	return exec.NewUpdate(sql.Sql, ParamList...)
}
