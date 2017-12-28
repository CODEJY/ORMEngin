ORM反射
===
[pml博客](http://m.blog.csdn.net/pmlpml/article/details/78850516)  
引用了pml老师的[sqlt](https://github.com/pmlpml/sqlt)

任务
---
1. orm规则
2. 实现自动插入数据
3. 实现查询结果自动映射

文件树
---
- dao: 包含实现自动插入以及查询映射的文件夹
    - helpfulFunc.go: 辅助函数集合，包括insert和find两个操作的所有辅助函数
    - orm-engine.go: 实现创建新的数据库驱动，实现自动插入，实现查询映射
- entity: 定义用户信息实体struct的文件夹
    - UserInfo.go: 定义用户信息struct，orm映射规则
- sqlt: pml老师的数据库开发模板文件，方便开发
- main.go: 测试程序功能 

测试
---
1. 首先创建数据库UserDB以及数据表userinfo  
![](https://github.com/CODEJY/ORMEngine/blob/master/ScreenShot/create%20database.png)
2. 查看userinfo表的信息  
![](https://github.com/CODEJY/ORMEngine/blob/master/ScreenShot/before.png)
3. 运行程序   
命令行输出：  
![](https://github.com/CODEJY/ORMEngine/blob/master/ScreenShot/cmd.png)
查看userinfo表是否写入成功：  
![](https://github.com/CODEJY/ORMEngine/blob/master/ScreenShot/after.png)
