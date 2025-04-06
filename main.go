package main

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

func PrimeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2

		for i := 3; i < 10; i += 1 {
			l := int(math.Sqrt(float64(i)))
			found := false

			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true

					break
				}
			}

			if !found {
				result <- i
			}
		}

		close(result)
	}()

	return result
}

func main() {
	pn := PrimeNumber()
	for n := range pn {
		fmt.Println(n)
	}

	start := time.Now()
	fav()
	fmt.Println("took: ", time.Since(start))

	context_practice()
}

func fav() {
	post := fetchPost()

	var wg sync.WaitGroup
	wg.Add(2)

	// channel の初期化
	// 2個のバッファを持った channel を作成
	resChan := make(chan any, 2)

	go fetchPostLikes(post, resChan, &wg)
	go fetchPostComments(post, resChan, &wg)

	wg.Wait()

	// resChan channel への送信を終了し channel を閉じる
	close(resChan)

	// channel が閉じられるまでループする
	for res := range resChan {
		fmt.Println("res: ", res)
	}
}

// 投稿を一件取得する関数.
func fetchPost() string {
	time.Sleep(time.Millisecond * 50)

	return "What programming languages do you prefer?"
}

// 投稿に紐づいたいいね数を取得する関数.
func fetchPostLikes(post string, reschan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)

	reschan <- 10
	wg.Done()
}

// 投稿に紐づいたコメントを全て取得する関数.
func fetchPostComments(post string, reschan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	reschan <- []string{"Golang", "Java", "Rust"}
	wg.Done()
}

func context_practice() {
	fmt.Println("start sub()")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
