/*
 * goroutine的应用场景之一: 挂为一个独立的服务。
 * 该例为单独开一个查询消息通知的服务
 */

package main

import "fmt"

func get_notification(user string) chan string{
   /*
    * 此处可以查询数据库获取新消息等等..
    */
    notifications := make(chan string)

    go func() { // 悬挂一个信道出去
        notifications <- fmt.Sprintf("Hi %s, welcome to weibo.com!", user)
    }()

    return notifications
}

func main() {
    jack := get_notification("jack") //  获取jack的消息
    joe := get_notification("joe") // 获取joe的消息

    // 获取消息的返回
    fmt.Println(<-jack)
    fmt.Println(<-joe)
}
