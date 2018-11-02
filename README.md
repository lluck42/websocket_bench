# websocket_bench

### why

github上没找到合适的 websocket 测试工具，虽说有js版的，但是却不知道为何需要 python 2.7 的支持，不想卸载 3.0 版本的

FarhadF/websocket_benchmark 的 benchmash 也比较不错，但 websocket.client 依赖中却少写了关键的 origin 地址，和未知原因导致的连接不能维持

http://www.blue-zero.com/WebSocket/ 看着还可以，也有 live demo ，但是呢，但是呢，竟然需要收费才能使用测试工具


所以呢，构思了下，采用 golang.org/x/net/websocket 的包搭建了 client 测试工具模拟实际并发情况

（golang.org/x/net/websocket 这个包在 github上的地址镜像为 github/golang/net/websocket, 需copy 到 golang.org/x/net/websocket 目录才能使用）


### usage
  `git clone https://github.com/lluck42/websocket_bench`
#### vim main.go 
#### edit origin setting and userNum setting
  `var url = "ws://127.0.0.1:9501/001"`  
  `var userNum = 10000`
#### run
  `go run main.go`
### close
  ctrl+c


请给小星星~

lluck42@163.com
2018年11月2日

