package main

import (
    "fmt"
    "time"
    "golang.org/x/net/context"
)

// 向外部提供的阻塞接口
// 计算 a + b, 注意 a, b 均不能为负
// 如果计算被中断, 则返回 -1
func Loop(ctx context.Context, count int) int {
    res := 0
    value := ctx.Value("extra")
    if value != nil {
        switch value.(type) {
        case int:
            count += value.(int)
        default:
            fmt.Println("extra is expected to be an integer")
        }
    }
    for i := 0; i < count; i++ {
        res += 1
        time.Sleep(1 * time.Second) // 但是由于我的机器指令集中没有这条指令,

        select {
        case <-ctx.Done():
            return -1
        default:
        }
        deadline, ok := ctx.Deadline()
        if ok {
            remainedTime := deadline.Sub(time.Now())
            fmt.Println("remained seconds", remainedTime)
            if remainedTime < 1 * time.Second {
                fmt.Println("pre-complete")
                return 99
            }
        }
        fmt.Println("loop", res, )
    }
    return res
}

func main() {
    {
        timeout := 5 * time.Second
        ctx, _ := context.WithTimeout(context.Background(), timeout)
        res := Loop(context.WithValue(ctx, "extra", 1), 2)
        if res < 0 {
            fmt.Println("not done yet, reason:", ctx.Err())
        } else {
            fmt.Printf("loop done %d\n", res)
        }
    }
    {
        ctx, cancel := context.WithCancel(context.Background())
        go func() {
            time.Sleep(3 * time.Second)
            cancel() // 在调用处主动取消
        }()
        res := Loop(ctx, 4)
        if res < 0 {
            fmt.Println("not done yet, reason:", ctx.Err())
        } else {
            fmt.Printf("loop done %d\n", res)
        }
    }
}

