package main

import (
	"context"
	"fmt"

	"github.com/tangx-labs/golang-context-with-from-demo/info"
)

func main() {
	ctx := context.Background()

	Liubei(ctx, 9)

	// Output:
	// 刘备: 曹操来了 9 万人
	// 关羽(1) <-: 曹操来了 9 万人
	// 关羽(2) ->: 曹操来了 90 万人
	// 张飞: 曹操来了 90 万人
}

func Liubei(ctx context.Context, n int) {

	fmt.Printf("刘备: 曹操来了 %d 万人\n", n)

	// 使用 context 重新包裹
	ctx = info.WithEnemyContext(ctx, n)

	Guanyu(ctx)

}

func Guanyu(ctx context.Context) {
	n := info.FromEnemyContext(ctx)

	fmt.Printf("关羽(1) <-: 曹操来了 %d 万人\n", n)
	// 扩大数量
	if n%2 == 1 {
		n = n * 10
	}

	fmt.Printf("关羽(2) ->: 曹操来了 %d 万人\n", n)

	// 使用 context 重新包裹
	ctx = info.WithEnemyContext(ctx, n)

	Zhangfei(ctx)
}

func Zhangfei(ctx context.Context) {
	n := info.FromEnemyContext(ctx)
	fmt.Printf("张飞: 曹操来了 %d 万人\n", n)
}
