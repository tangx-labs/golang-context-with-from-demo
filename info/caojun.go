package info

import (
	"context"
)

// 定义曹军的 key ， 用于 context 的值传递
type enemy int // 定义一个新类型
const keyCaojun = enemy(0)

// 将曹军数量注入到 context 中
func WithEnemyContext(ctx context.Context, number int) context.Context {
	return context.WithValue(ctx, keyCaojun, number)
}

// 从 context 中取出曹军数量
func FromEnemyContext(ctx context.Context) int {
	val := ctx.Value(keyCaojun)

	n, ok := val.(int) // 类型断言
	if !ok {
		return 0
	}
	return n
}
