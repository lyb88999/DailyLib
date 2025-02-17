package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}
	t.wg.Done()
}

func main() {
	var taskFunc = func(data interface{}) {
		task := data.(*Task)
		task.Do()
		fmt.Printf("task: %d sum: %d\n", task.index, task.sum)
	}
	p, _ := ants.NewPoolWithFunc(10, taskFunc)
	defer p.Release()
	const (
		DataSize    = 10000
		DataPerTask = 100
	)
	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}
	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	tasks := make([]*Task, 0, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{
			index: i,
			nums:  nums[i*DataPerTask : (i+1)*DataPerTask],
			wg:    &wg,
		}
		tasks = append(tasks, task)
		err := p.Invoke(task)
		if err != nil {
			return
		}
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())

	// 验证
	sum := 0
	for _, task := range tasks {
		sum += task.sum
	}
	expect := 0
	for _, num := range nums {
		expect += num
	}
	fmt.Printf("expect: %d, actual: %d\n", expect, sum)
}
