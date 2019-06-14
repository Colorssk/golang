package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Number int
	Id     int
}

type Result struct {
	job *Job
	sum int
}

func calc(job *Job, result chan *Result) {
	var sum int
	number := job.Number
	for number != 0 { // 这样相当于while
		tmp := number % 10
		sum += tmp
		number /= 10
	}

	r := &Result{
		job: job,
		sum: sum,
	}

	result <- r
}

func Worker(jobChan chan *Job, resultChan chan *Result) {

	for job := range jobChan {
		calc(job, resultChan)
	}
}

func startWorkerPool(num int, jobChan chan *Job, resultChan chan *Result) {

	for i := 0; i < num; i++ {
		go Worker(jobChan, resultChan)
	}
}

func printResult(resultChan chan *Result) {
	for result := range resultChan {
		fmt.Printf("job id:%v number:%v result:%d\n", result.job.Id, result.job.Number, result.sum)
	}
}

func main() {

	jobChan := make(chan *Job, 1000)
	resultChan := make(chan *Result, 1000)

	startWorkerPool(128, jobChan, resultChan)

	go printResult(resultChan) // 另起的线程只要结构体中有值就会全部打印
	var id int
	for {
		id++
		number := rand.Int()
		job := &Job{
			Id:     id,
			Number: number,
		}

		jobChan <- job
	}
}
