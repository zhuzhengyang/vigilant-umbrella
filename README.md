# go plugin热更新总结

* open plugin时会执行plugin的init
* 不能open两次同一个BuildID的plugin，即使plugin代码完全相同，也要装作两个package才能open（[BuildID](https://github.com/golang/go/blob/be04da8f0c5cb57e2736cfac8c76971d7d8cfc6f/src/cmd/go/internal/load/pkg.go#L1621) 由文件名等计算而来）
* 新老plugin的内部变量是独立的
* 可以在主程序里创建data，plugin里定义func，将&data传到plugin的func里
* 改变data结构，然后编译出来的plugin,**不能**被主进程open，会报 [plugin was built with a different version of package](https://github.com/zhuzhengyang/vigilant-umbrella/tree/master/change_struct)
* plugin需要注意维护所引用data的生命周期，在资源被完全释放后，老plugin才会被gc回收
* [网上的老例子](https://github.com/scgywx/myplugin/issues/1)，在1.14上跑不起来

#### 热更新只适合耦合度低相对独立的模块. 
#### 如果plugin切的太细，会难以管理，可能带来性能问题?

## 一些对比
微服务: 主进程保有网络和数据的状态，算法部分提出为单独一个或几个服务  
脚本调用：主进程保有网络和数据的状态，算法部分调用脚本语言  
plugin: 算法部分做为plugin加载到主进程中  

|       | 微服务  | 脚本调用  | plugin |
|:----  |:------ |:-------- |:------ |
| 优点   | 可以用go开发 | 部署使用简单安全 | 可以用go开发，开销小 |
| 缺点   | rpc调用网络开销大；复杂场景需要把整个状态输入输出，服务膨胀，容易使主进程变成网关，微服务取代主进程 | 脚本语言运行效率低；复杂场景输入输出大;  | 管理困难，不知道老plugin是否被释放; plugin和主进程共同依赖的go pkg版本必须相同 |

