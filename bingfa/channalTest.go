package main

import (
	"fmt"
	"time"
)

func recv(a chan int)  {
	ret :=  <- a
	fmt.Print("接收成功",ret,"\n")
}

func main ()  {
	//a := make(chan int,1)
	//// 启用goroutine从通道接收值
	//an := make(chan int)
	//go recv(an)
	//an <- 20
	//a <- 10  //往通道里面发送值，如果没有缓存，一只在阻塞死锁
	////所以没有缓存区的就必须有一个接收者直接接收,所以无缓存的通道就是同步通道，有缓存的就是异步通道
	////fmt.Println(len(a))
	//fmt.Println("发送成功\n")
	//b := <- a
	//fmt.Println(b)
	//close(a)

	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("把%d放入ch1\n",i)
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			fmt.Printf("从ch1中取出%d\n" ,i)
			ch2 <- i * i
		}
		close(ch2)
		fmt.Println("关闭ch1输入\n" )
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Printf("最后打印从ch1中取出的值%d\n" ,i)
		fmt.Println(i)
	}
	time.Sleep(time.Second*10)
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}

}
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}