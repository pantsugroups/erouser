package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	User struct {
		gorm.Model
		Age    int
		Name   string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
		Num    int    `gorm:"AUTO_INCREMENT"` // 自增
		Email  string
		Verify bool   //判断是否验证成功
		IM     string //QQ号
	}
)

var e *echo.Echo
var db *gorm.DB
var port string
var confuse string
