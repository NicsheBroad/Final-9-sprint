package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
		{"один элемент", 1, 1},
		{"обычный размер", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.want {
				t.Fatalf("длина слайса: получили %d, ожидали %d", len(data), tt.want)
			}
			for _, v := range data {
				if v <= 0 {
					t.Fatalf("ожидали положительное число, получили %d", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{7}, 7},
		{"максимум в начале", []int{9, 2, 3}, 9},
		{"максимум в середине", []int{1, 8, 3}, 8},
		{"максимум в конце", []int{1, 2, 10}, 10},
		{"одинаковые элементы", []int{4, 4, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("получили %d, ожидали %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{7}, 7},
		{"меньше частей, чем CHUNKS", []int{3, 1, 5}, 5},
		{"длина не кратна CHUNKS, максимум в остатке", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 99}, 99},
		{"максимум в первой части", []int{99, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("получили %d, ожидали %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{1, 7, 8, 9, 100, 1001, 100_000} {
		data := generateRandomElements(size)
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("размер %d: maxChunks=%d, maximum=%d", size, got, want)
		}
	}
}
