项目结构
```
├── common                                      // 放全局变量
├── configs                                     // 配置文件
├── internal                                    // 业务代码
    ├── controller                              //控制层
    ├── dao                                     //数据层
    ├── model                                   //数据层的实体
    ├── service                                 //业务层
    ├── router                                  //路由层，包括路由中间件
    ├── vo                                      //控制层的实体
├── log                                         // log日志
├── static                                      // 静态文件
├── utils                                       // 工具类
├── main.go                                     // 程序入口 
├── pkg                                         // 打包，需执行packge.sh脚本生成
├── packge.sh                                   // 打包脚本


```

