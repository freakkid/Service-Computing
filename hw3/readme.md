# 开发 web 服务程序

## Task requirements

1. 熟悉 go 服务器工作原理

2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo

3. 使用 curl 工具访问 web 程序

4. 对 web 执行压力测试

The task requirements is from [ex-cloudgo-start](doc/ex-cloudgo-start.html)

## About Todos

This program easy to use can help people to record what they need to do in a simple way. Users just type URL into their browser or enter command at the command line. There are only four command for using totally. The data of users are saved in database stored on the server so users should be online when using this tools.

## Install
Before running this program, please make sure that you have installed the following tools.

+ golang

+ MySQL

    [Install MySQL on ubuntu](https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-14-04)

+ a database driver

    [mysql](https://github.com/go-sql-driver/mysql) is a database driver of the program. 

    > $ go get github.com/go-sql-driver/mysql

## Usage

+ Reset some paraments

    To work on your computer, the code of the program should be reset according to the conditions of your computer. 
    
    You should reset the **username**, **password**, **addrs**, **port** on the [server_data.go](). The **username**, **password** are the username and password of your mysql databases and **addrs**, **port** are the ip address and port of tcp connection to your databases.

+ get the program



// user/registe?username=XXX&password=XXX
// todo/add?username=XXX&&password=XXX&item=XXX
// todo/delete?username=XXX&password=XXX&itemIndex=XXX
// todo/show?username=XXX&&password=XXX

/* This is a server to save datas that about the things need to be done.
 * User should register by username and password. User can add or delete
 * the todolist by loginnng with username and password because for
 * "IsPrivate" defaults to TURE. However, user can set "IsPrivate" to
 * FALSE for convince that he can edit his todolist only with username.
 * Of course, other users can edit the unprotected todolist at will.
 */
