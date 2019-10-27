##Algorithm

[wk3.go](wk3.go)\

##Review
###docker 源码分析 前五章
[main structure](docker-image.png)
主要介绍了docker的价格 docker client, docker daemon. Docker Daemon 中 主要核心功能为 docker server 和 New Daemon.\
New Daemon 中创建了daemon所需要的网络和镜像的配置，以及驱动。
server 则是基于api结构的高并发设计（得利于golang的高并发支持）。

##Tips
https://stackoverflow.com/questions/7669906/how-to-execute-php-block-from-terminal-without-saving-to-a-file
use php -a to write testing script.

##Share
[Deduplicating Amazon SQS Messages](https://medium.com/avmconsulting-blog/deduplicating-amazon-sqs-messages-dc114d1e6545)

