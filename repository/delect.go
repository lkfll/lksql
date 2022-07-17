package repository

import "github.com/lkfll/lksql/exec"

// // 按表格删除
// // DELETE FROM user WHERE 1=1;
// func (Repository *Repository) DeleteByTable(tableName string) func(string) *exec.Delete {
// 	return exec.NewDelete(Repository.BaseSQL.Delete[tableName].Sql)
// }

// 按表格删除
// DELETE FROM user WHERE 1=1;
func (Repository *EntityRepository) Delete() func(string) *exec.Delete {
	return exec.NewDelete(Repository.EntitySql.Delete.Sql)
}
