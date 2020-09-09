# delay_queue 延迟队列

此项目只供学习golang与redis使用。

## 大体实现思路：
- 利用redis的zset实现。
- zset中score是TTR（time to return），member是数据的id。
- 后台goroutine每秒钟从zset中利用zrange取出值最小的前5个，判断当前时间戳timestamp如果大于等于该数据TTR，那么将这个数据放入ready queue。
- ready queue是redis的list，lpush加入队列，brpop出队列。



![image](https://qschou.oss-cn-hangzhou.aliyuncs.com/20190218195349703.png)



## 已有功能：
- 提供grpc接口
- 向队列中Push数据，比如序列化后的json数据，以及需要改数据返回的时间戳。
- 从队列中Pop数据，提供一个超时时间timeout，接口会阻塞，直到ready queue中有数据或者超过timeout。
- publish模式，Push时可以传入notify_url，到期后发送post请求通知，如果post请求失败，或者post请求返回的内容不是'SUCCESS'，会发起重试，重试的间隔为2s/8s/30s/2m/5m/30m/60m。
- 删除已经push的消息，根据Push传回的唯一id，来删除延时队列中对应的数据。

## 接口实现：
- Push：为每个数据生成id，存入数据池和zset中。
- Pop：利用redis的list中brpop命令，传入超时时间，接口阻塞直到有lpush数据或者超时，返回第一条数据。
- publish模式：后台goroutine不断从ready queue中阻塞式轮询，如果有数据，那么post请求指定的notify_url地址通知调用者。
- del：将数据池中的key删除，同时将zset中的member也删除。

