# llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll

反射优化

考虑外边在封装一层调用

#### 介绍
拼接sql语句

#### 软件架构
软件架构说明


#### 安装教程



#### 使用说明
field 字段名： 不写默认转换为 小写下划线分隔
table 所属表： 不写默认和主键key表相同
join  连接:    连接表，会全部拼接到查询语句中
 type Issues struct {
 	Id           int    `Key:"id" table:"issues"` //主键
 	IssuesName   string //`table:"issues" field:"issues_name"`
 	Label        string //`table:"issues" field:"label"`
 	Description  string //`table:"issues" field:"description"`
 	PriorityName string `join:"LEFT JOIN priority ON issues.pid=priority.id" table:"priority" ` // 优先级名字
  	Uid          int    //`table:"issues" field:"uid"`
  }

 结果分析调用 查询结果的一行进行赋值创建issues对象
 顺序和字段issues对象字段顺序一样
  func (issues Issues) ArrayOf(b ...string) interface{} {
  	var ret Issues
  	ret.Id, _ = strconv.Atoi(b[0])
  	ret.IssuesName = b[1]
  	ret.Label = b[2]
  	ret.Description = b[3]
  	ret.Uid, _ = strconv.Atoi(b[4])
  	ret.PriorityName = b[5]
  	return ret
  }

 issues转为map
  func (i Issues) ToMap() map[string]interface{} {
  	m := make(map[string]interface{})
  	m["id"] = i.Id
  	m["uid"] = i.Uid
  	m["description"] = i.Description
  	m["issues_name"] = i.IssuesName
  	m["label"] = i.Label
  	m["priority_name"] = i.PriorityName
  	return m
  }

 开始使用
  repository, err := lksql.DefaultFacory(Issues{})
  if err != nil {
  	fmt.Printf("err: %v\n", err)
  }

 删除
  repository.DeleteByTable("issues")("id=?").Go(DB, 1)

 修改
  var issues Issues
  issues.Id = 1
  issues.Description = "--_--"
  issues.IssuesName = "name"
  issues.Label = "lable"
  issues.Uid = 1
  issues.PriorityName = "pname"
  repository.UpdateByTable("issues", issues)("id=?").Go(DB, 1)

 增加
  list := make([]lksql.Type, 0)
  for i := 0; i < 5; i++ {
  	var issues Issues
  	issues.Id = 1
  	issues.Description = "--_--"
  	issues.IssuesName = "name"
  	issues.Label = "lable"
  	issues.Uid = 1
  	issues.PriorityName = "pname"
  	list = append(list, issues)
  }
  repository.InsertByTable("issues", list...).Go(DB)

 建议使用save保存不用多次拼接
 查询
 // SELECT  issues.id , issues.issues_name , issues.label , issues.description , priority.priority_name , issues.uid
 //	FROM  issues
 //        LEFT JOIN priority ON issues.pid=priority.id

 // SelectAll := repository.Select().Save()
  i, _ := SelectAll.Go(DB)
  fmt.Printf("i: %v\n", i)

 	// SELECT  issues.id , issues.issues_name , issues.label , issues.description , priority.priority_name , issues.uid
 	// FROM  issues
  //        LEFT JOIN priority ON issues.pid=priority.id
  // WHERE lable=?
  // ORDER BY id DESC
  // LIMIT ?,?
  i, _ = repository.Select().Where("label=?").Order("id DESC").Limit("?,?").Go(DB, "lable", 3, 10)
  fmt.Printf("i: %v\n", i)

 分组查询
  // SELECT  issues.id,label,priority_name,Min(uid)
  // FROM  issues
  // 		LEFT JOIN priority ON issues.pid=priority.id
  // GROUP BY issues.id,label,priority_name
  i, _ = repository.Select().

 分组之后返回字段需要修改，括号依次是 分组，修改返回字段，一行返回结果的处理
  Group("issues.id,label,priority_name")("issues.id,label,priority_name,Min(uid)")(
  	func(s ...string) interface{} {
  		return nil
  	}).Go(DB)
  fmt.Printf("i: %v\n", i)

 子查询
  sq := repository.Select().Where("issues.id=?").SetSelectField("uid")(nil).Save()
 	// SELECT  issues.id , issues.issues_name , issues.label , issues.description , priority.priority_name , issues.uid
 // FROM  issues
 //         LEFT JOIN priority ON issues.pid=priority.id
 // WHERE uid=(SELECT  uid
 // 	FROM  issues
 //         	LEFT JOIN priority ON issues.pid=priority.id
 // WHERE issues.id=?)
 repository.Select().Where(fmt.Sprint("uid<", "(", sq.GetSql(), ")")).Go(DB, 1)
#### 参与贡献




#### 特技

