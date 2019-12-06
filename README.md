port-scanner  
===
a simple hack tool,build with Go.  

## Example  
查看本机开放的端口`scanner -ip 127.0.0.1`  
查看目标机器开放的端口`scanner -ip 目标机器ip地址`  
查看本机3306端口是否开放`scanner -ip 127.0.0.1 -port 3306`  

## 待实现  
- 中英文档   
- 扫描局域网内所有存活主机  
- 用户名密码爆破(like Hydra)  
- 探测目标端口可能的应用程序  
- 指定协议(TCP、UDP)  
- 本地端口及其占用程序  
## 已实现  
- 指定主机全端口扫描  
- 指定主机特定端口扫描  
- [伪]猜测目标端口可能的应用程序    
## 如何安装并使用
**您可以下载源码后自行编译,或使用下方编译好的二进制文件**  
#### 通过源码安装:  
`cd port-scanner/cmd/ && go build scanner.go`  
  
#### 自行下载二进制:  
`将scanner文件夹添加到环境变量Path内,打开cmd输入scanner即可看到帮助信息。
或将scanner.exe放入已经存在Path中的目录内(如C:\Windows\system32\)`

如果您是临时使用，可以打开cmd切换到本文件夹内，输入scanner.exe即可。
## 下载  
[百度网盘](https://pan.baidu.com/s/11kUnKzcTWU3lPBkQhBqDtw)
## 作者  
[BuTn](https://github.com/kimmosc2)
