# cjuncms

发音“菌CMS”，C 是不发音的。类似 django 的 d 不发音一样。

```bash
# 执行 初始化前端依赖
npm ci

# 启动前端（开发）
npm run dev

# 编译前端
npm run build
```

```bash
# 执行 数据库 初始化 命令
go run ./cmds/db_init

# 启动服务（开发）
go run .

# 编译
go build
```

```bash
# 执行 fake 命令
go run ./cmds/fake_employee
go run ./cmds/fake_project
```

## 目录

- client 前端
- server 后端

## 命令

```bash
# 更新 cjungo
go get github.com/cjungo/cjungo
```

```bash
# 多个表
gentool -db mysql -dsn "root:123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local" -tables "cj_department, cj_permission, cj_employee_permission, cj_employee, cj_operation, cj_demand, cj_demand_employee, cj_demand_project, cj_project, cj_project_employee, cj_pass, cj_script, cj_machine_cpu_time, cj_machine_virtual_memory, cj_machine_disk_usage, cj_machine_process" -modelPkgName="model" -outPath="./entity" -fieldNullable -fieldWithIndexTag -fieldWithTypeTag  -fieldSignable 
```