package sql2struct

import (
	"database/sql" // sql 包提供了保证SQL或类SQL数据库的泛用接口
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 连接数据库的驱动
)

// DBModel 是整个数据库连接的核心对象，也可以认为是数据库
type DBModel struct {
	DBEngine *sql.DB // 连接数据库的引擎
	DBInfo   *DBInfo //
}

// DBInfo 用于存储用于连接 MySQL 的一些基本信息
type DBInfo struct {
	DBType   string // 要连接的数据库类型
	Host     string // 数据库所在主机 socket
	UserName string // 数据库用户名
	PassWord string // 数据库密码
	Charset  string // 数据库的编码格式(例如，utf-8)
}

// TableColumn 用于存储 COLUMNS 表中需要的一些字段
type TableColumn struct {
	ColumnName    string // 列名
	DataType      string // 列所对应的数据类型
	IsNullable    string // 判断某个列是否为空
	ColumnKey     string // 如果某个列是逐渐，则添加 pri
	ColumnType    string // 列类型
	ColumnComment string // 列注释
}

// DBTypeToStructType 数据库类型到 GO 结构体类型的映射
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

// NewDBModel 新建一个连接数据库的模型
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{
		DBInfo: info,
	}
}

// Connect 连接数据库
func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.PassWord,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)

	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}

	return nil
}

// GetColumns 获取表中列的信息
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	// query 数据库查询请求
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName) // 返回的是行的集合
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}

	defer rows.Close() // 最后要关闭当前这条请求

	var columns []*TableColumn

	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable,
			&column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
