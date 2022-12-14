# Http Server

## 路由注册

#### 1. 规范路由

除了我们自己定义的业务路由外，`Server` 自动帮我们注册了两个路由：
* `/api.json`：基于 OpenAPIv3 协议自动生成的接口文档
* `swaggerPath`：自动生成的 SwaggerUI 页面

这两个功能默认开启，方便开发者查看文档和调试代码，可以通过 `openapiPath` 和 `swaggerPath` 两个配置项控制是否开启。

#### 2. 路由方法定义

其中输入参数和输出参数都是两个，并且都是必须的一个都不能少。简单介绍下：

| 参数                  | 说明     | 注意事项                                                     |
| :-------------------- | :------- | :----------------------------------------------------------- |
| `ctx context.Context` | 上下文   | `Server`组件会自动从请求中获取并传递给接口方法               |
| `req *Request`        | 请求对象 | 就算没有接收参数也要定义，因为请求结构体中不仅仅包含请求参数的定义，也包含了接口的请求定义。 |
| `res *Response`       | 返回对象 | 就算没有返回参数也要定义，因为返回结构体中不仅仅包含返回参数的定义，也可以包含接口返回定义。 |
| `err error`           | 错误对象 | `Server`通过该参数判断接口执行成功或失败。                   |

#### 3. 请求和返回结构体

