package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(SIZE) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, v := range data[1:] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	chunkMaxes := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaxes[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()

	return maximum(chunkMaxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
