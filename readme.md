# Ero User
一款使用RESTful API规范的，可作为execute直接编译使用的，也可以作为package倒入使用的用户信息管理API

## 使用第三方包
如果想要作为包引用进行二次开发，请详细阅读以下项目的文档
```
"github.com/urfave/cli"
"github.com/labstack/echo"
"github.com/jinzhu/gorm"
```

## 作为程序使用
编译
```bash
./install.bat | ./install.sh
```
运行
```bash
cd bin/
./erouser r --port 80
```
文档
```bash
./erouser -h
```

## 作为包引用
引用了本项目之后，请使用`Hijack`函数，接管本程序

然后使用`LoadRoutes`函数，加载本项目的路由。