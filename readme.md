# goHostStat
## 监控服务器状态，离线报警，收集服务器信息 
~~~
Usage of ./goHostSata:
  -l string
        listen addr (default ":12345")
  -s    server
  -t int
        Lost Time for alert /s (default 60)
  -u    getdata

~~~

## featureList
### V2
* 连接方式支持 tcp/udp/quic
* 替换提取主机参数方式
* 报警消息分级