## Algorithm

[wk4.go](wk4.go)

## Review
### docker 源码分析 前六，七章
这两章介绍了 daemon 和 container 网络的实现。daemon一章节介绍通过veth把外部网络如何带入到container中。容器网络一章则具体介绍了\
1. 四种网络模式。
2. 从client怎么配置
3. execdrive如何调用/创建 libcontainer
4. libcontainer 是如和用exec.cmd直接调用linux内核的。

其中，execdriver 主要用native为例（linux）。而linux的namespace和cgroups的架构让人印象深刻，一个用于隔离进程，而一个则是用于限制资源。docker则依赖这两个设定完成了虚拟化。
关于container的网络共享问题，kubernates的pod是个更好的答案，值得研究一下。

## Tips
using v8 as a back end javascript engine when building back end lanuage could help to avoid some trouble. (resource share and limitation) shall read the doc to learn more details.

## Share
[v8](https://v8.dev/docs)的文档。
