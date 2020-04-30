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

