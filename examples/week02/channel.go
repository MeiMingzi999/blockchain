package main

import "fmt"

func sendNum() {

	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("通道关闭")
				return
			}
			fmt.Printf("接收到数字: %d\n", msg)
		}

	}

}

func bufferChannel() {
	ch := make(chan int, 10)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("通道关闭")
				return
			}
			fmt.Printf("接收到数字: %d\n", msg)

		}

	}

}

func main() {
	//sendNum()
	bufferChannel()
}
