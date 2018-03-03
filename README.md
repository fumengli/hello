## 运行程序 ##
    * 1 bee run  或者 go run main.go 
## 在linux下配置beego环境 ##
## 配置项目和git关联 ##
    * 1 在本地任意目录执行  beego new hello ,会在$GOPATH/src目录下新建hello目录
    * 2 在github中创建一个仓库(仓n.go
## 在linux下配置beego环境 ##
## 配置项目和git关联 ##
    * 1在本地任意目录执行  beego new hello ,会在$GOPATH/src目录下新建hello目录
    * 2 在github中创建一个仓库(仓库的名称一定也是hello) ,不要勾选README.md
    * 3 将本地的hello文件夹 和远程 git仓库相关联
        git init //git初始化,创建.git隐藏文件夹,就将本地的代码变成git仓库
        git remote add origin https://github.com/wbcg00/hello.git
    * 4 添加README.md
        touch README.md 添加内容
        git add README.md
        git commit README.md -m"日志"
        git push -u origin master
    * 5 提交hello项目的源码
        git status 查看需要添加的文件和文件夹
        git add (重复) 添加
        git commit -a -m"日志"
        git push -u origin master
    * 6 至此完成本地项目和git的同步关联库的名称一定也是hello) ,不要勾选README.md




