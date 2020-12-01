学习笔记

### 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

应该 Wrap 这个 error

原因是：

1. 产生错误的地方，是可以提供有效的上下文，因此要把上下文抛给上层
2. 为了上层能够感知该错误（方便上层进行等值判断），同时又不受携带上下文参数的影响，因此在抛出错误时候需要 Wrap

#### 代码示例

```text

├── README.md
├── api
│   └── handler.go
├── dao
│   ├── implement.go
│   ├── interface.go
│   └── mocks
│       └── implement_mock.go
├── errs
│   └── errors.go
├── main.go
├── model
│   └── user.go
└── service
    ├── implement.go
    └── interface.go

```

api handler 是处理请求，返回响应的地方，它依赖 service，service 依赖 dao

dao 抛出的 error，经由 service，传到 api handler，统一在这里记录错误，处理错误码


#### 在 api Handler 处理错误日志

```go

user, err := h.service.GetUserById(id)
if err != nil {
    logrus.Errorf("get user by id error: %+v", err)
}

```

#### service 直接返回 dao 抛出的错误

```go

func (i *implService) GetUserById(id uint) (user model.User, err error) {
	return i.dao.GetUserById(id)
}

```

#### 如何解决依赖

依赖 interface 而非实现

#### 如何测试

根据 interface 生成 mock 实现