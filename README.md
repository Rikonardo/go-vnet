# Go Virtual Network

Tiny library for making virtual listeners in Go.

Example:

```go
package main

import (
    "github.com/Rikonardo/go-vnet"
    "github.com/gin-gonic/gin"
    "io"
    "net"
)

func main() {
    realListener, _ := net.Listen("tcp", ":8080")
    listener := vnet.Builder().
        AddrFunc(func() net.Addr {
            return realListener.Addr()
        }).
        Build()
	
    go func() {
        for {
            conn, _ := realListener.Accept()
            vconn := listener.Connect()
            go io.Copy(vconn, conn)
            go io.Copy(conn, vconn)
        }
    }()
	
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello World!")
    })
    r.RunListener(listener)
}
```