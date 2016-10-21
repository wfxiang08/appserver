# 文件上传服务

## 环境搭建:
* 工作目录: 约定使用 ~/gowork
* 创建项目
	* beego框架:
		* 参考文档: http://beego.me/quickstart
		* go get github.com/astaxie/beego
		* go get github.com/beego/bee
		* bee new fileupload
	* 测试:
		* go get -u github.com/stretchr/testify/assert
		* 运行TestCase:
			*  go test service -v -run "TestAudioOperation"
			* `如何跑TestCase?`

## 下载代码?

```bash
go get git.chunyu.me/feiwang/appserver
cd appserver
source start_env.sh
cd src
gpm install
```

## 运维:
* 参考: http://beego.me/docs/install/bee.md
* bee pack 打包代码和编译结果
* 参考: http://beego.me/docs/deploy/
	* conf/app.conf
	* 这个部分如何定制呢?
