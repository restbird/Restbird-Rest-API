package main
import "time"

func main() {
	for {
		select {
		case <-time.Tick(time.Millisecond * 10000):
			fmt.Println("hello, curren time is :", time.Now())
		}
	}

}
