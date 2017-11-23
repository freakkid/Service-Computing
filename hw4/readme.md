# 
基本要求

支持静态文件服务
支持简单 js 访问
提交表单，并输出一个表格
对 /unknown 给出开发中的提示，返回码 5xx

基本要求

## 支持静态文件服务

* get static file:
> http://localhost:8080/static/
返回一个有关使用方法的静态网页
> 


支持简单 js 访问
* api
> http://localhost:8080/api/register?username=hnx&password=hnx 
> api/registe?username=XXX&password=XXX
返回username和成败success/fail和具体反馈

> http://localhost:8080/api/add?username=hnx&password=hnx&item=do home work
>  api/add?username=XXX&password=XXX&item=XXX
返回username和成败success/fail和具体反馈

> http://localhost:8080/api/show?username=hnx&password=hnx
> api/show?username=XXX&password=XXX
返回username和成败success/fail和具体反馈

> http://localhost:8080/delete?username=hnx&password=hnx&itemIndex=2
> api/delete?username=XXX&password=XXX&itemIndex=

返回username和成败success/fail和具体反馈

提交表单，并输出一个表格
对 /unknown 给出开发中的提示，返回码 5xx
* form

> http://localhost:8080/register
输入username  password=hnx

> http://localhost:8080/addItem
输入username  password=hnx

> http://localhost:8080/deleteItem
输入username  password=hnx

> http://localhost:8080/showItems
输入username  password=hnx

数据库UUID