# goice
go implementation of ICE( RFC 5245),it doesn't depends any c library,so it can runs on all the platforms which go support.

## how to use
please reference ice/exmaple/example.go, this a simple example to present how to setup a p2p connection between two nodes without any signal server.

## how to setup a turnserver on ubutu

```bash
apt install turnserver
turnserver -u bai:bai
```
turnserver's default port is 3478,make sure your firewall doesn't block it.
