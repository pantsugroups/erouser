package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

//外部调用劫持的过程
func Hijack(db2 *gorm.DB, e2 *echo.Echo) {
	e = e2
	db = db2
}
