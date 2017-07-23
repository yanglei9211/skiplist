package main
import (
    "fmt"
    "skiplist"
)

func main()  {
    head := skiplist.Skiplist{}
    fmt.Println(head.Level)

    //head := skiplist.Skiplist{}
    //dt := []int{3,2,1}
    //for _, x := range dt {
    //    node := skiplist.Skiplist{}
    //    node.Next = head.Next
    //    node.Score = x
    //    head.Next = &node
    //}
    //tp := head.Next
    //fmt.Println(head.Score)
    //fmt.Println(tp)
    //for tp!= nil {
    //   fmt.Println(tp.Score)
    //   tp = tp.Next
    //}
}
