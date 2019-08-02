package dao

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"text/template"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbCfg struct {
	User     string
	Password string
	DbName   string
	Host     string
	MaxOpen  int
	MaxIdle  int
	Enabled  bool
	DebugMode bool
}
var conn *gorm.DB
func InitDB(){
	var  dbInfo dbCfg

	if _, err := toml.DecodeFile("./configs/mysql.toml", &dbInfo); err != nil {
		panic(err)
	}
	dbNet, _ := template.New("dbconfig").Parse("{{.User}}:{{.Password}}@tcp({{.Host}})/{{.DbName}}?charset=utf8&parseTime=true&loc=Local")
	dbBuffer := new(bytes.Buffer)
	dbNet.Execute(dbBuffer, dbInfo)

	rdb,err := gorm.Open("mysql",dbBuffer.String())
	if err !=nil {
		fmt.Println("数据库链接错误")
		panic(err)
	}
	rdb.DB().SetMaxOpenConns(dbInfo.MaxOpen)
	rdb.DB().SetMaxIdleConns(dbInfo.MaxIdle)
	if dbInfo.DebugMode {
		conn = rdb.Debug()
	} else {
		conn = rdb
	}
	fmt.Println(conn)
}

type BaseDao struct {
	Conn *gorm.DB
}
//添加
func(this *BaseDao) Create(entity interface{}) error{
	err := this.Conn.Create(entity).Error
	return err
}
//修改
func(this *BaseDao) Update(condition interface{},entity interface{}) error{

	err := this.Conn.Model(condition).Where(condition).Update(entity).Error
	return err
}
//删除
func(this *BaseDao) Delete(condition interface{}) error{
	err := this.Conn.Where(condition).Delete(condition).Error
	return err
}
//查询
func(this *BaseDao) List(condition interface{},rows interface{}) error{
	err := this.Conn.Where(condition).Find(rows).Error
	return err
}
//查询某一条
func(this *BaseDao) First(rows interface{}) (interface{},error){
	err := this.Conn.First(rows).Error
	return rows,err
}
//分页查询
func(this *BaseDao) ListByPage (condition interface{},rows interface{},page *Page) error{
	err := this.Conn.Where(condition).Limit(page.PageSize).Offset((page.PageIndex - 1) * page.PageSize).Find(rows).Error
	return err
}
//条数
func(this *BaseDao) Count (condition interface{},rows interface{}) (int,error){
	var count int = 0
	err := this.Conn.Where(condition).Model(rows).Count(&count).Error
	return count,err
}


//获取事务
func GetTransaction() *gorm.DB{
	return conn.Begin()
}

func NewBaseDao(con *gorm.DB) BaseDao{
	if con == nil {
		return BaseDao{
			Conn:conn,
		}
	} else {
		return BaseDao{
			Conn:con,
		}
	}
}

type Page struct {
	PageSize  int
	PageIndex int
}



