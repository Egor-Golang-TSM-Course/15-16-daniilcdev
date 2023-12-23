package cancelcallstack

import (
	"context"
	"fmt"
	"time"
)

func CallAndCancel(parentCtx context.Context, delay time.Duration) {
	ctx, cancel := context.WithCancel(parentCtx)
	timer := time.NewTimer(delay)
	defer cancel()

	sum := 0
	go recursiveIncrement(ctx, 1, &sum)

	<-timer.C
	fmt.Println("timer fired")
	timer.Stop()
	cancel()

	<-ctx.Done()
	err := ctx.Err()
	if err != nil {
		fmt.Printf("ctx err: %v\n", err)
	}

	fmt.Printf("exit with captured sum=%d\n", sum)
}

func recursiveIncrement(parentCtx context.Context, step int, c *int) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	for i := 0; i < 1000; i++ {
		*c++
	}

	go recursiveIncrement(ctx, step+1, c)

	<-ctx.Done()
	// if ctx.Err() != nil {
	// 	fmt.Printf("[depth=%d] stopped by cancel\n", step)
	// }
}
