# 开发 web 服务程序

## Task requirements

1. 熟悉 go 服务器工作原理

2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo

3. 使用 curl 工具访问 web 程序

4. 对 web 执行压力测试

The task requirements is from [ex-cloudgo-start](doc/ex-cloudgo-start.html)

In order to know golang web better, I **didn't use any frame** in this homework

## About Todos

This program easy to use can help people to record what they need to do in a simple way. Users just type URL into their browser or enter command at the command line. There are only four command for using totally. The data of users are saved in database stored on the server so users should be online when using this tools. User should register by username and password. User can add, delete and show the todolist with username and password.

## Install
Before running this program, please make sure that you have installed the following tools.

+ golang

+ MySQL

    [Install MySQL on ubuntu](https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-14-04)

+ a database driver

    [mysql](https://github.com/go-sql-driver/mysql) is a database driver of the program. 

    > $ go get github.com/go-sql-driver/mysql

Run the program

+ Get the program

    > $ go get github.com/freakkid/Service-Computing/hw3

+ Reset some paraments

    To work on your computer, the code of the program should be reset according to the conditions of your computer. 
    
    You should reset the **username**, **password**, **addrs**, **port** on the [server_data.go](). The **username**, **password** are the username and password of your mysql databases and **addrs**, **port** are the ip address and port of tcp connection to your databases.

+ Install again after reseting
    
    > $ go install github.com/freakkid/Service-Computing/hw3
    
    Had better make sure that have added GOPATH in PATH.

    The default port is 8080:

    > $ hw3

    ```
    2017/11/08 20:58:25 [Todos]  port 8080 is listening
    
    ...
    ```

    Or the port can be seted by:

    > $ hw3 -p 9090

    ```
    2017/11/08 20:58:25 [Todos]  port 9090 is listening
    
     ...
    ```

    End the program by _Ctrl + C_

## Usage and Curl Testing

+ Register as a user

    * The register URL format is: user/registe?username=XXX&password=XXX

    * You can type URL into browser:

    >  http://localhost:8080/user/register?username=hnx&password=hnx

    ![images/register.png](images/register.png)

    Or you can enter on the command line:

    > $ curl -v  http://localhost:8080/user/register?username=hnx\&password=hnx

    ```    
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /user/register?username=hnx&password=hnx HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 08 Nov 2017 13:17:44 GMT
    < Content-Length: 134
    < Content-Type: text/html; charset=utf-8
    < 
    <!DOCTYPE html>
    <html>

    <head>
        <title>Todos</title>
    </head>

    <body>
        <h1>hnx</h1>

        <p>register success</p>
    </body>

    * Connection #0 to host localhost left intact

    ```

    * If you type empty username or password, you will receive a error message:

    > $ curl -v  httalhost:8080/user/register?username=\&password=hnx

    ```
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /user/register?username=&password=hnx HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 08 Nov 2017 13:22:34 GMT
    < Content-Length: 156
    < Content-Type: text/html; charset=utf-8
    < 
    <!DOCTYPE html>
    <html>

    <head>
        <title>Todos</title>
    </head>

    <body>
        <h1></h1>

        <p>username and password should be non-empty</p>
    </body>

    * Connection #0 to host localhost left intact
    </html>
    ```

    * If you register twice with the same username, you will receive errors:

    > $ curl -v  http://localhost:8080/user/register?username=hnx\&password=hnx

    ```
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /user/register?username=hnx&password=hnx HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 08 Nov 2017 13:22:12 GMT
    < Content-Length: 164
    < Content-Type: text/html; charset=utf-8
    < 
    <!DOCTYPE html>
    <html>

    <head>
        <title>Todos</title>
    </head>

    <body>
        <h1>hnx</h1>

        <p>register fail: the username may have been used</p>
    </body>

    * Connection #0 to host localhost left intact
    </html>
    ```

    * Or, if you type invalid URL, you may get a 404 page:

    > $ curl -v  httalhost:8080/user/regis

    ```    
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /user/regis HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 404 Not Found
    < Content-Type: text/plain; charset=utf-8
    < X-Content-Type-Options: nosniff
    < Date: Wed, 08 Nov 2017 13:27:58 GMT
    < Content-Length: 19
    < 
    404 page not found
    * Connection #0 to host localhost left intact
    ```

+ Add todo item

    * The add URL format is: todo/add?username=XXX&password=XXX&item=XXX

    * You can type URL into browser:

        >  http://localhost:8080/todo/add?username=hnx&password=hnx&item=do home work

        ![images/add.png](images/add.png)

        Or you can enter on the command line:

        > $ curl -v  http://localhost:8080/todo/add?username=hnx\&password=hnx\&item=havefun

        ```
        *   Trying 127.0.0.1...
        * Connected to localhost (127.0.0.1) port 8080 (#0)
        > GET /todo/add?username=hnx&password=hnx&item=havefun HTTP/1.1
        > Host: localhost:8080
        > User-Agent: curl/7.47.0
        > Accept: */*
        > 
        < HTTP/1.1 200 OK
        < Date: Wed, 08 Nov 2017 13:44:35 GMT
        < Content-Length: 129
        < Content-Type: text/html; charset=utf-8
        < 
        <!DOCTYPE html>
        <html>

        <head>
            <title>Todos</title>
        </head>

        <body>
            <h1>hnx</h1>

            <p>add success</p>
        </body>

        * Connection #0 to host localhost left intact
        </html>
        ```
    
    * If you type empty item:
        
        > $ curl -v  http://localhost:8080/todo/add?username=hnx\&password=hnx\&item=

        ```
        *   Trying 127.0.0.1...
        * Connected to localhost (127.0.0.1) port 8080 (#0)
        > GET /todo/add?username=hnx&password=hnx&item= HTTP/1.1
        > Host: localhost:8080
        > User-Agent: curl/7.47.0
        > Accept: */*
        > 
        < HTTP/1.1 200 OK
        < Date: Wed, 08 Nov 2017 13:45:49 GMT
        < Content-Length: 195
        < Content-Type: text/html; charset=utf-8
        < 
        <!DOCTYPE html>
        <html>

        <head>
            <title>Todos</title>
        </head>

        <body>
            <h1>hnx</h1>

            <p>add fail: please check username and password and the item should be non-empty</p>
        </body>

        * Connection #0 to host localhost left intact
        </html>
        ```

+ Show all todo items
    
    * The show URL format is: todo/show?username=XXX&password=XXX

    * You can type URL into browser:

    >  http://localhost:8080/todo/show?username=hnx&password=hnx

    ![images/show.png](images/show.png)

    Or you can enter on the command line:

    > $ curl -v  http://localhost:8080/todo/show?username=hnx\&password=hnx

    ```
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /todo/show?username=hnx&password=hnx HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 08 Nov 2017 13:48:52 GMT
    < Content-Length: 320
    < Content-Type: text/html; charset=utf-8
    < 
    <!DOCTYPE html>
    <html>

    <head>
        <title>Todos</title>
    </head>

    <body>
        <h1>hnx</h1>

        <p>show success： you have 4 todo items</p>

        <div>
            <br /><br />   1 : do homework  <br /><br />   2 : havefun  <br /><br />   3 : eat  <br /><br />   4 : sleep  <br /><br /> 
        </div>
    </body>

    * Connection #0 to host localhost left intact
    </html>
    ```

+ Delete todo item
    
    * The delete URL format is: todo/delete?username=XXX&password=XXX&itemIndex=XXX

    * You can type URL into browser to delete the second item:

    >  http://localhost:8080/todo/delete?username=hnx&password=hnx&itemIndex=2

    ![images/delete.png](images/delete.png)

    And show todolist after deleting:

   ![images/show2.png](images/show2.png) 

    Or you can enter on the command line:

    > $ curl -v  http://localhost:8080/todo/delete?username=hnx\&password=hnx\&itemIndex=3

    ```
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /todo/delete?username=hnx&password=hnx&itemIndex=3 HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 08 Nov 2017 13:54:48 GMT
    < Content-Length: 132
    < Content-Type: text/html; charset=utf-8
    < 
    <!DOCTYPE html>
    <html>

    <head>
        <title>Todos</title>
    </head>

    <body>
        <h1>hnx</h1>

        <p>delete success</p>
    </body>

    * Connection #0 to host localhost left intact
    </html>
    ```
    And show all items:

    > $ curl -v  httalhost:8080/todo/show?username=hnx\&password=hnx

    ```
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /todo/show?username=hnx&password=hnx HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 08 Nov 2017 14:18:41 GMT
    < Content-Length: 257
    < Content-Type: text/html; charset=utf-8
    < 
    <!DOCTYPE html>
    <html>

    <head>
        <title>Todos</title>
    </head>

    <body>
        <h1>hnx</h1>

        <p>show success： you have 2 todo items</p>

        <div>
            <br /><br />   1 : do homework  <br /><br />   2 : eat  <br /><br /> 
        </div>
    </body>

    * Connection #0 to host localhost left intact
    </html>
    ```

    If you type invalid itemIndex such as non-numeric string, less than one or larger than todolist size, you may get error tips:

    > $ curl -v  httalhost:8080/todo/delete?username=hnx\&password=hnx\&itemIndex=0
    
    ```
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 8080 (#0)
    > GET /todo/delete?username=hnx&password=hnx&itemIndex=0 HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 08 Nov 2017 14:03:30 GMT
    < Content-Length: 200
    < Content-Type: text/html; charset=utf-8
    < 
    <!DOCTYPE html>
    <html>

    <head>
        <title>Todos</title>
    </head>

    <body>
        <h1>hnx</h1>

        <p>delete fail: please check username and password and the item index should be valid</p>
    </body>

    * Connection #0 to host localhost left intact
    </html>
    ```

+ Server will show some information about the requests:

    ```
    2017/11/08 21:12:08 [Todos]  port 8080 is listening
    2017/11/08 21:12:10 [Todos]  GET | 200 | register
    2017/11/08 21:17:44 [Todos]  GET | 200 | register
    2017/11/08 21:22:12 [Todos]  GET | 200 | register
    2017/11/08 21:22:34 [Todos]  GET | 200 | register
    2017/11/08 21:27:58 [Todos]  GET | 400
    2017/11/08 21:32:52 [Todos]  GET | 200 | register
    2017/11/08 21:33:49 [Todos]  GET | 200 | register
    2017/11/08 21:34:08 [Todos]  GET | 200 | add
    2017/11/08 21:34:14 [Todos]  GET | 200 | add
    2017/11/08 21:34:15 [Todos]  GET | 200 | register
    2017/11/08 21:34:24 [Todos]  GET | 200 | register
    2017/11/08 21:42:56 [Todos]  GET | 200 | show
    2017/11/08 21:44:24 [Todos]  GET | 200 | add
    2017/11/08 21:44:35 [Todos]  GET | 200 | add
    2017/11/08 21:45:49 [Todos]  GET | 200 | add
    2017/11/08 21:46:48 [Todos]  GET | 400
    2017/11/08 21:48:40 [Todos]  GET | 200 | show
    2017/11/08 21:48:52 [Todos]  GET | 200 | show
    2017/11/08 21:52:30 [Todos]  GET | 200 | delete
    2017/11/08 21:53:10 [Todos]  GET | 200 | show
    2017/11/08 21:54:48 [Todos]  GET | 200 | delete
    2017/11/08 22:03:16 [Todos]  GET | 200 | show
    2017/11/08 22:03:21 [Todos]  GET | 200 | show
    2017/11/08 22:03:30 [Todos]  GET | 200 | delete
    ...
    ```
## ab Testing

For the database connection has the max connection number, the server will report errors if requests are too much. So I test my program using URL that has nothing about SQL.

commonly used parameters:
* -n means the numbers of all requests
* -c means the numbers of requests in one time and the default number is 1

I test my web program by 10000 requests and 1000 requests one time(that is 1000 users).

> $ ab -n 10000 -c 1000 http://localhost:8080/user/register?username=

This is the ab-test result, I write some Chinese comments in it by "#":

```
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost   # 服务器主机
Server Port:            8080        # 服务器端口

Document Path:          /user/register?username=
Document Length:        156 bytes   # 测试的页面文档长度

Concurrency Level:      1000        # 并发用户数，即-c参数指定的数量
Time taken for tests:   0.848 seconds   # 测试总用时
Complete requests:      10000       # 测试完成的请求数量
Failed requests:        0           # 测试失败的请求数量
Total transferred:      2730000 bytes   # 响应数据长度总和
HTML transferred:       1560000 bytes   # html内容长度
Requests per second:    11792.94 [#/sec] (mean) # 每秒请求数量
Time per request:       84.796 [ms] (mean)      # 平均请求响应时间
Time per request:       0.085 [ms] (mean, across all concurrent requests)   # 每个请求实际运行的平均时间
Transfer rate:          3144.02 [Kbytes/sec] received   # 平均每秒网络流量，帮助排除网络流量过大导致响应时间延长的可能

Connection Times (ms)   # 网络消耗时间的组成
              min  mean[+/-sd] median   max
Connect:        0   13   8.7     11      39
Processing:     3   38  18.6     33     253
Waiting:        0   34  18.4     29     248
Total:          4   51  21.3     48     261

Percentage of the requests served within a certain time (ms)
  50%     48    # 50%的用户请求在48ms内完成，后面的类似
  66%     59
  75%     63
  80%     66
  90%     80
  95%     96
  98%    104
  99%    109
 100%    261 (longest request)
```



## At last, thanks for reviewing!
