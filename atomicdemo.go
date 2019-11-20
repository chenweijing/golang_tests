

import (
	"fmt"
	"sync/atomic"
)

var (
	// 序列号
	seq int64
)

//GenID 序列号生成器
func GenID() int64 {
	//尝试原子的增加序列号
	return atomic.AddInt64(&seq, 1)
}

func main2() {
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())
}
