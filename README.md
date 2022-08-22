# behappy-url-shortener
基于go+redis实现的短地址服务

本家: https://github.com/zero-archive/node-url-shortener


先clone项目

## 项目启动方式:
### 二进制文件

1. go build -o app .
2. 执行app run进行默认值启动
启动参数可参考如下[app run {COMMOND}]
```
-p, --port    Port number for the Express application          [default: 3000]
--redis-host  Redis Server hostname                     [default: "localhost"]
--redis-port  Redis Server port number                         [default: 6379]
--redis-pass  Redis Server password                           [default: false]
--redis-db    Redis DB index                                      [default: 0]
```
使用app run -h方式查看帮助

### docker
1. 进入根目录, 执行docker build -t urlshortener:latest .
2. 执行docker run -d --privileged=true --restart=always -p 3000:3000 --name urlshortener urlshortener:latest run -p 3000

## 使用方式: RESTful API
NOTE: You can send the post requests without the date and c_new params

POST /shorten with json to create shorturl hash
```json
long_url=http://google.com, 
start_date="", 
end_date="", 
c_new=false
```


GET /:hash with shorturl hash to redirect longurl
```json
实现跳转,
如果出现错误则跳转至错误页
```
