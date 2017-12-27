1. 依据文档图6-1，用中文描述 Reactive 动机

个人理解：Motivation for Reactive Client即是一个异步数据流，相比于同步请求需要一个一个顺序返回请求，reactive编程提供了用异步的方式达到可以同时处理多个来自客户端请求，从而缩短响应时间。另外为了不把内部服务暴露给外界，客户端对内部服务的多个请求的回应将被合并成单个响应再返回给客户端。

2. 使用 go HTTPClient 实现图 6-2 的 Naive Approach
