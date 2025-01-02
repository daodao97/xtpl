# xtpl

application template for [github.com/daodao97/xgo](https://github.com/daodao97/xgo)

## 特性

- 自带管理后台-无需前端开发
- 自动生成接口文档

![preview-admin](https://github.com/user-attachments/assets/c340095a-e14b-499f-8638-ec0dc872b6da)

![preview-docs](https://github.com/user-attachments/assets/fc85257e-5e12-4ebe-baf0-b38a816d64ad)

## 项目结构

- cmd/main.go 主程序
- internal/admin 后台管理
- internal/api 前端接口
- internal/conf 配置
- internal/dao 数据库
- internal/model 模型
- internal/service 服务
- internal/utils 工具
- internal/pkg 公共包

## 数据库

`docs/operator.sql` 用户表结构, 自行导入, 默认用户名: test, 密码: test

## run

```bash
go run ./cmd/... --enable-openapi true
```

## 文档

http://localhost:4001/docs
