package infrastructure

import (
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/interfaces/database"
	"github.com/Kazuki-Tohyama/go-nuxt-blog/api/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := gorm.Open("mysql", "root:root@tcp(mysql:3306)/sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	conn.AutoMigrate(&domain.Article{})
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
  return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
  return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
  return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value ...interface{}) *gorm.DB {
  return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value ...interface{}) *gorm.DB {
  return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
  return handler.Conn.Where(query, args...)
}
