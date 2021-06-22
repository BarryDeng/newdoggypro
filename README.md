# MyDoggyPro

基于ReactJS和Gin的一个简单猜图游戏。

## 构建方式

需要事先安装好NodeJS环境、NPM和yarn以生成前端，并安装有Golang 1.16工具链，go mod处于auto或者on状态（`go env -w GO111MODULE=of`）。

```
$ cd web
$ yarn
$ npm run build
$ cd ../
$ go mod download
$ go build
```

安装MySQL数据库，root密码设置为`123456`，并创建一个名为`demo`的Schema。

# 运行方式

Windows下
```powershell
.\newdoggypro.exe
```

Linux、MacOS下
```bash
./newdoggypro
```