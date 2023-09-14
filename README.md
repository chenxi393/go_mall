### 此项目来自gin mall 打算一步步跟着视频做
### 第一个视频的待办是 没有成功运行getmoney
* TODO:第四个视频 跨域请求第一次见 跨域也不懂
* TODO:图片上传到七牛云 可以看他的CSDN的博客
* TODO:对于支付密码的key加密的 没懂
* TODO:这只是一个小demo 其中很多的细节经不起推敲 需要自己优化很多逻辑 看成熟的后端详细的接口 根据接口来开发项目
* TODO:小知识点都写在对应代码块的注释上 记得去看 有时间整理一下
* TODO: MySQL读写分离 一般公司运维负责配置 开发只需要写代码（感觉可以在docker配多个mysql）
* 目前计划两个个版本
* V0: 学习用的
* V1: 完善的单体版本
* V2: 完善的微服务版本 （这个一定要完成）

> PS 这个开发流程其实可以参考字节结营的答辩文档 感觉后端架构都差不多
> 可以找到往届（或者本届）的项目的TOP 然后学习它们的文档 COPY一下

### 项目开发流程
1. 想好架构图 创建所有项目所需文件夹
2. 读取配置文件 
3. 写数据库的模型 然后使用ORM建表
4. 用户注册和登录
5. 用户更新信息和上传头像
6. 用户进行邮箱的发信操作和校验
7. 用户获取金额
8. 添加轮播图和日志库
9. 商品的一些操作
10. 地址的操作
11. 收藏夹的操作
12. 购物车先跳过 都是CRUD 完成订单操作（有关事务）就去看github终版源码

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
// SHA256是一种常用的加密协议 区块链中hash值的生成也使用了SHA256
```

> 有人认为token是一个反转的session session把东西存在服务端，而token加密后给客户端（用户）


### rebase merge
写过很多次笔记了 而自己实战老是搞出莫须有的冲突
* 我在自己的分支上提交了两个记录 main分支也有两个提交记录 如下图演示
```shell
        E---F  feature
       /    
---A---B---C---D  main

当前在feature 分支
git fetch              #获取最新的提交 也可以使用git pull main 指定分支分支
git rebase origin/main #变基 也就是当前分支基于origin/main 有冲突详见git.md
git push

# main为保护分支
pull request 走pr流程（目前pr的三种合并方式也没用完全理清 第一种方式老是不可用）

# main非保护分支
git checkout main
git merge feature      #这一步就是将刚才变基了的合并到main 合并后main是一条完成的直线
git push

# 结果
---A---B---C---D---E---F  feature       
---A---B---C---D---E---F  main

# 善后工作
如果还想继续利用feature分支 就得强制提交 因为rebase之后部分提交ID（hash值会改变）
这里就是当时起冲突的原因 rebase之后还在原分支上开发
# 解决办法一
git push -f feature    #本地feature强制提交到远程仓库
# 解决办法二
删除云端feature分支 重新建立一个feature分支 （在一个开发者只能拥有一个分支的前提下）
```
**rebase的优缺点**
1. 保持提交历史的线性 且不会像merge一样创建一个合并提交记录 
2. 缺点 解决冲突频繁 需要从左往右一个个commit来解决冲突 而merge是按最终的结果解决一次代码冲突 没有合并的提交记录其实也是一个缺点 出了问题不好溯源

>rebase的准则：不要在公共分支上使用rebase 

>在公共分支上使用merge 在个人分支上使用rebase

**TODO: 具体还需要学习多人开发 git使用的规则**