## 在linux下配置beego环境
## 配置项目和git关联
1.在本地任意目录执行  beego new beego_helloworld ,会在$GOPATH/src目录下新建beego_helloworld目录
2 在github中创建一个仓库(仓库的名称一定也是beego_helloworld) ,不要勾选README.md
3 将本地的beego_helloworld文件夹 和远程 git仓库相关联
	git init //git初始化,创建.git隐藏文件夹,就将本地的代码变成git仓库
	git remote add origin evenio@163.com:wbcg00/https://github.com/wbcg00/beego_helloworld.git
4 添加README.md
	touch README.md 添加内容
	git add README.md
	git commit README.md -m"日志"
	git push -u origin master 
5 提交beegoi_helloworld项目的源码
	git status 查看需要添加的文件和文件夹
	git add (重复) 添加 
	git commit -a -m"日志"
        git push -u origin master
6 至此完成本地项目和git的同步关联
