package main

func main() {
	//传入channal的指针才能改变channal的cap
	//var r chan int
	//fmt.Printf("len:%d,cap:%d", len(r), cap(r))
	//changeR(&r)
	//fmt.Printf("len:%d,cap:%d", len(r), cap(r))

	//说明切片只要不是copy，都是复制的引用
	//a := []int{1,2,3,4,5,6,7,8,9}
	//var c [][]int
	//var b [][]int
	//c = append(c, a[0:3])
	//c = append(c, a[3:6])
	//c = append(c, a[6:])
	//b = append(b, append(c[0], c[1]...))
	//a[3:] = b[0]
	//for _, v := range c{
	//	go func(v1 []int) {
	//		v1[0] = 9
	//	}(v)
	//}
	//time.Sleep(2*time.Second)
	//fmt.Println(a)

	//一定要传入wg的指针才行
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go testWait2(&wg)
	//wg.Wait()
	//fmt.Println("done")

}

//func testWait(wg sync.WaitGroup, a int) {
//	var wg2 sync.WaitGroup
//	wg2.Add(1)
//	testWait(wg2, a-1)
//	if a != 0 {
//		wg2.Wait()
//	}
//
//	wg.Done()
//}
//
//func testWait2(wg *sync.WaitGroup) {
//	wg.Done()
//}

//func changeR(r *chan int) {
//	*r = make(chan int, 10)
//}