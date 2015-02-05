package main

import (
    "fmt"
    "github.com/satori/go.uuid"
)


func main() {
    //使用UUID V4创建
    u1 := uuid.NewV4()
    fmt.Printf("UUID字符串: %s\n", u1)

    // Parsing UUID from string input
    u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    if err != nil {
        fmt.Printf("出错了: %s", err)
    }
    fmt.Printf("解析成功: %s", u2)
}

/***Read Me***
1.install uuid package:go get github.com/satori/go.uuid
2.run:go run uuidtest.go

run result:
UUIDv4: 7ad519c8-f855-4afc-8c3f-621334d07a22
Successfully parsed: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
***********/