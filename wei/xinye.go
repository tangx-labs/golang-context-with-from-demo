package wei

import (
	"context"
)

// 定义一个新类型
type enemy int

// 定义曹军的 key ， 用作 context 传递
var keyCaojun = enemy(0)

func WithEnemyContext(ctx context.Context, number int) context.Context {
	return context.WithValue(ctx, keyCaojun, number)
}

func FromEnemyContext(ctx context.Context) int {
	val := ctx.Value(keyCaojun)

	n, ok := val.(int) // 类型断言
	if !ok {
		return 0
	}
	return n
}
