#goice
goice 是一个 纯 go 实现的ICE(RFC5245)P2P 通信库,不依赖任何 c 语言库,这样可以保证 goice 可以运行在所有支持 go 的平台上.

## 如何使用
请参考ice/example/example.go, 这是一个简单的不依赖信令服务器,演示两个节点如何建立 P2P 连接.
他依赖一个 turnserver. 我不能保证这个 turnserver 一直都能正常运行.如果turnserver 不能正常工作,请搭建自己的 turnserver.

## 在 ubuntu上搭建自己的 turnserver
```bash
apt install turnserver
turnserver -u bai:bai
```
turnserver 默认端口3478,请保证外部可以使用.
