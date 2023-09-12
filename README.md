### 此项目来自gin mail 打算一步步跟着视频做
### 第一个视频的待办是 没有成功运行getmoney
* TODO:第四个视频 跨域请求第一次见 跨域也不懂
* TODO：图片上传到七牛云 可以看他的CSDN的博客
* 哪个对于支付密码的key加密的 没懂

> PS 这个开发流程其实可以参考字节结营的答辩文档 感觉后端架构都差不多
> 可以找到往届（或者本届）的项目的TOP 然后学习它们的文档 COPY一下
### 项目开发流程
1. 想好架构图 创建所有项目所需文件夹
2. 读取配置文件 
3. 写数据库的模型 然后使用ORM建表
4. 用户注册和登录 登录会返回一个token 用户拿到这个token（相当于用于认证的东西） 后续发送请求才能成功
5. 用户更新信息和上传头像
6. 用户进行邮箱的发信操作和校验
7. 用户获取金额
8. 添加轮播图和日志库
9. 商品的一些操作
11. 地址的操作 要验证传入的addr_id去查数据库里的地址是不是token所拥有的 不然不应该操作


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

[REST API](https://poe.com/s/MLHJzVDNryeEIIWjXgpD)

### JWT
1. JSON Web Token 通过数字签名，以JSON对象为载体，用于授权认证
2. 一旦用户登录，后续每个请求都包含JWT，系统在每次处理用户请求之前，都要进行token鉴权，通过再进行处理
3. JWT由三部分组成 用.拼接

Encoded
```txt
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNiwiZW1haWwiOiJ5eGlhbzQ2MEBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiIsIm9wZXJhdGlvbl90eXBlIjoxLCJleHAiOjE2OTQwMjkyNTMsImlzcyI6ImdvLW1haWwgYnkgRFlGIn0.D7q8s39jSgsU8QGjL2DZyhnx95XwFTaWZpdVlqJMxuQ
```
Decoded

```json
Header
{
  "alg": "HS256",
  "typ": "JWT"
}
Payload
{
  "user_id": 16,
  "email": "yxiao460@gmail.com",
  "password": "",
  "operation_type": 1,
  "exp": 1694029253,
  "iss": "go-mail by DYF"
}
Signature     // secret 保存在服务端
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

> 有人认为token是一个反转的session session把东西存在服务端，而token加密后给客户端（用户）

