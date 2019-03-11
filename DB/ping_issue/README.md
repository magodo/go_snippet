这个项目是为了测试mysql driver的 Ping() 函数的bug，顺便测试pg。

具体描述可见我的[blog](https://magodo.github.io/golang-sql/#4-ping). 

这里的*mysql_max_conn.sh*用于模拟连接数满的情况。

mysql的不同版本通过go module来选择。
