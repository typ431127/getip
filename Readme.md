一个简单的运维调试工具，此工具旨在帮助运维定位真实ip，在一些复杂的场景中，一个域名要经过很多层转发才会到达后端服务，这个过程中ip地址因为转发会出现后端服务获取到的是内网ip的情况。

浏览器直接访问ip:port，会返回如下json ,根据返回信息开发运维人员能够更好的调试获取的Header信息，以此来获取真实IP
```json
{
    "Client-Ip": "127.0.0.1",
    "Header": {
        "Accept": [
            "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
        ],
        "Accept-Encoding": [
            "gzip, deflate, br"
        ],
        "Accept-Language": [
            "zh,zh-CN;q=0.9"
        ],
        "Connection": [
            "keep-alive"
        ],
        "Cookie": [
            "csrftoken=;sessionid="
        ],
        "Dnt": [
            "1"
        ],
        "Sec-Ch-Ua": [
            "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\""
        ],
        "Sec-Ch-Ua-Mobile": [
            "?0"
        ],
        "Sec-Ch-Ua-Platform": [
            "\"Windows\""
        ],
        "Sec-Fetch-Dest": [
            "document"
        ],
        "Sec-Fetch-Mode": [
            "navigate"
        ],
        "Sec-Fetch-Site": [
            "none"
        ],
        "Sec-Fetch-User": [
            "?1"
        ],
        "Upgrade-Insecure-Requests": [
            "1"
        ],
        "User-Agent": [
            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"
        ]
    },
    "Method": "GET",
    "RealIp": "127.0.0.1",
    "RequestURI": "/"
}
```
- Client-Ip: 程序获取到的IP
- Header: 获取到的Header 头信息
- RealIp: 通过X-Forwarded-For,X-Yestea-Forwarded-For 由程序获取到真实IP
- Method: 请求方法
- RequestURI: 请求url

#### 启动
```shell
./app -port :8081
```

#### docker运行
```shell
docker run -itd -p 8087:8080 typ431127/getip:0.1.5
```
realip库参考: https://github.com/tomasen/realip