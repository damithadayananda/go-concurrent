package range_thr_channel

import "fmt"

func RangeTroughChannel()  {
	c:=make(chan string,1)

	t:=concurrentTest{

	}
	go t.Print(c,"job 1")
	go t.Print(c,"job 2")
	go t.Print(c,"job 3")
	go t.Print(c,"job 4")
	go t.Print(c,"job 5")

	for i:=0;i<50;i++{

		fmt.Println(<-c)
	}
	//close(c)
}

type concurrentTest struct {
//hkhkhkhkhk
}

type concurrent interface {
	Print(c1 chan string,s string )
}

var t =0
func(co * concurrentTest) Print(ch chan string,s string){
	for i:=0;i<10;i++{
		v:=fmt.Sprintf("|%v|%v|",s,i)
		ch <- v
		//fmt.Println(v)
	}

}
