1. 依据文档图6-1，用中文描述 Reactive 动机

    个人理解：Motivation for Reactive Client即是一个异步数据流，相比于同步请求需要一个一个顺序返回请求，reactive编程提供了用异步的方式达到可以同时处理多个来自客户端请求，从而缩短响应时间。另外为了不把内部服务暴露给外界，客户端对内部服务的多个请求的回应将被合并成单个响应再返回给客户端。

2. 使用 go HTTPClient 实现图 6-2 的 Naive Approach

3. 为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3

    题目2和3代码运行：
    > go get 


4. 对比两种实现，用数据说明 go 异步 REST 服务协作的优势
5. 思考： 是否存在一般性的解决方案？
<!-- q=apple

from=en

to=zh

appid=2015063000000001

salt=1435660288

平台分配的密钥: 12345678

生成sign：

>拼接字符串1

拼接appid=2015063000000001+q=apple+salt=1435660288+密钥=12345678

得到字符串1 =2015063000000001apple143566028812345678

>计算签名sign（对字符串1做md5加密，注意计算md5之前，串1必须为UTF-8编码）

sign=md5(2015063000000001apple143566028812345678)

sign=f89f9594663708c1605f3d736d01d2d4

完整请求为：

http://api.fanyi.baidu.com/api/trans/vip/translate?q=apple&from=en&to=zh&appid=2015063000000001&salt=1435660288&sign=f89f9594663708c1605f3d736d01d2d4 -->