### 此项目来自gin mail 打算一步步跟着视频做
### 第一个视频的待办是 没有成功运行getmoney
* 原因应该是jwt鉴权不懂 怎么发送url
* 第四个视频 跨域请求第一次见
* TODO：图片上传到七牛云 可以看他的CSDN的博客


### 项目开发流程
1. 想好架构图 创建所有项目所需文件
2. 读取配置文件 
3. 写数据库的模型 然后使用ORM软件建表
4. 用户注册和登录 使用gin


### 各种文档链接
1. [GORM](https://gorm.io/zh_CN/docs/create.html#%E9%BB%98%E8%AE%A4%E5%80%BC)
2. [gin文档](https://gin-gonic.com/zh-cn/docs/)


### HTTP协议
Method URL HTTP版本
Header1: Value1
Header2: Value2
...
请求体（可选）

```http
POST /api/users HTTP/1.1
Host: example.com
User-Agent: Win64
{
  "name": "John Doe",
  "email": "johndoe@example.com"
}
```