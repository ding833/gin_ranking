1. 自定义写入日志：使用logger包的logger自定义打印日志方法
    打印日志信息到自己命名的文件：orderlog(没有文件可以自动生成)
2. 自动打印日志，success or error 到runtime\log下的文件中
3. 路由配置：
     //以中间键的方式调用日志中间键
	    r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	    //添加到Gin路由中，确保每个请求都会经过找个中间键，提供全局异常处理
	    r.Use(logger.Recover)
