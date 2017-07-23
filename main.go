package main

import (
	"fmt"
	"skiplist"
)

func calcScore(x int) float64 {
	return float64(x) * 1.1
}

func main() {
	sp := skiplist.Skiplist{}
	sp.Init()
	datas := []skiplist.Data{
		skiplist.Data{X: 1},
		skiplist.Data{X: 4},
		skiplist.Data{X: 8},
		skiplist.Data{X: 12},
		skiplist.Data{X: 19},
		skiplist.Data{X: 33},
		skiplist.Data{X: 45},
		skiplist.Data{X: 62},
		skiplist.Data{X: 69},
		skiplist.Data{X: 123},
		skiplist.Data{X: 150},
		skiplist.Data{X: 189},
		skiplist.Data{X: 216},
		skiplist.Data{X: 246},
		skiplist.Data{X: 286},
		skiplist.Data{X: 323},
		skiplist.Data{X: 366},
		skiplist.Data{X: 423},
		skiplist.Data{X: 532},
		skiplist.Data{X: 587},
	}
	scores := make([]float64, 0, len(datas))
	for _, dt := range datas {
		scores = append(scores, calcScore(dt.X))
	}
	for i := range datas {
		sp.Insert(scores[i], &datas[i])
	}
	var x *skiplist.SkiplistNode
	for i := 0; i < skiplist.SKIPLIST_MAX_LEVEL; i++ {
		if sp.Head.Level[i].Forward != nil {
			fmt.Printf("%d: ", i+1)
			x = sp.Head.Level[i].Forward
			for x != nil {
				fmt.Printf("%d --- ", x.Obj.X)
				x = x.Level[i].Forward
			}
			fmt.Println()
		}
	}
	qRange := skiplist.ScoreRange{Left: 234, Right: 444}
	res := sp.FirstInRange(&qRange)
	fmt.Println(res.Obj, res.Score)
	res = sp.LastInRange(&qRange)
	fmt.Println(res.Obj, res.Score)
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
