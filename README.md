### 此项目来自gin mail 打算一步步跟着视频做
### 第一个视频的待办是 没有成功运行getmoney
* 原因应该是jwt鉴权不懂 怎么发送url
* 第四个视频 跨域请求第一次见
* TODO：图片上传到七牛云 可以看他的CSDN的博客
* TODO：这一块jwt鉴权不懂（需要再学习）
* 还有一个要学习的是明文密文的互相转换 还有前面的加密操作也不是很懂


### 项目开发流程
1. 想好架构图 创建所有项目所需文件
2. 读取配置文件 
3. 写数据库的模型 然后使用ORM软件建表
4. 用户注册和登录 登录会返回一个token 用户拿到这个token（相当于用于认证的东西） 后续发送请求才能成功
5. 用户更新信息和上传头像
6. 用户进行邮箱的发信操作和校验
7. 用户获取金额


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