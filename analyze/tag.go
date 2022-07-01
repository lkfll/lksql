package analyze

// field tag 成员标签
const TableNameTag string = "table" // 所属表名字
const FieldNameTag string = "field" // 数据库对应字段名字
const JoinTag string = "join"       // 连接 LEFT JOIN issues i on (u.id = i.uid)
const KeyTag string = "Key"         // 主键 主表

// TODO 扩展
