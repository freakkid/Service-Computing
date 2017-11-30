# 

+ 作业要求

    ![10](img/10.png)

    [作业来源](http://blog.csdn.net/pmlpml/article/details/78602290#六进一步工作)

+ 
1. 使用 xorm 或 gorm 实现本文的程序，从编程效率、程序结构、服务性能等角度对比 database/sql 与 orm 实现的异同！ 
    orm 是否就是实现了 dao 的自动化？
使用 ab 测试性能
参考 Java JdbcTemplate 的设计思想，设计 GoSqlTemplate 的原型, 使得 sql 操作对于爱写 sql 的程序猿操作数据库更容易。 
轻量级别的扩展，程序员的最爱
程序猿不怕写 sql ，怕的是线程安全处理和错误处理
sql 的 CRUD 操作 database/sql 具有强烈的模板特征，适当的回调可以让程序员自己编写 sql 语句和处理 RowMapping
建立在本文 SQLExecer 接口之上做包装，直观上是有利的选择
暂时不用考虑占位符等数据库移植问题，方便使用 mysql 或 sqlite3 就可以
参考资源：github.com/jmoiron/sqlx