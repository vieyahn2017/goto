package main
 
import (
    "fmt"
    //"strings"
    "strconv"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func (l *ListNode) print () {
    for l != nil {
        fmt.Println(l.Val)
        l = l.Next
    }
}

func (l *ListNode) show () string {
    var s1 string
    var s2 string  //end
    for l != nil {
        //strings.Join([]string{s, "&ListNode{,", string(l.Val), ","}, "")
        s1 += "&ListNode{" + strconv.Itoa(l.Val) + ", "
        s2 += "}"
        l = l.Next
    }

    // s1 += "nil"
    // s1 += s2
    //fmt.Println(s1)

    return s1 + "nil" + s2
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 进位值, 只可能为0或1
    promotion := 0

    // 结果表的头结点
    var head *ListNode
    // 保存结果表的尾结点
    var rear *ListNode
    for nil != l1 || nil != l2 {
        sum := 0
        if nil != l1 {
            sum += l1.Val
            l1 = l1.Next
        }
        if nil != l2 {
            sum += l2.Val
            l2 = l2.Next
        }

        sum += promotion
        promotion = 0

        if sum >= 10 {
            promotion = 1
            sum = sum % 10
        }

        node := &ListNode{
            sum,
            nil,
        }

        if nil == head {
            head = node
            rear = node
        } else {
            rear.Next = node
            rear = node
        }

    }

    if promotion > 0 {
        rear.Next = &ListNode{
            promotion,
            nil,
        }
    }

    return head
}

func main() {

    a := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
    b := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
    //c := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}

    var c = addTwoNumbers(a, b) 
    c.print()
    fmt.Println(c.show())

}