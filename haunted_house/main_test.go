package main

import "testing"

func Test_getMaxGroup(t *testing.T) {
	type args struct {
		n      int
		people [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Initial case",
			args: args{
				n: 6,
				people: [][]int{
					{1, 2},
					{1, 4},
					{0, 3},
					{0, 1},
					{3, 4},
					{0, 2},
				},
			},
			want: 3,
		},
		{
			name: "All can go",
			args: args{
				n: 2,
				people: [][]int{
					{1, 2},
					{1, 4},
				},
			},
			want: 2,
		},
		{
			name: "No groups formed",
			args: args{
				n: 2,
				people: [][]int{
					{3, 5},
					{4, 4},
				},
			},
			want: 0,
		},
		{
			name: "No groups formed",
			args: args{
				n: 2,
				people: [][]int{
					{0, 1},
					{0, 0},
				},
			},
			want: 1,
		},
		{
			name: "Not all can go",
			args: args{
				n: 2,
				people: [][]int{
					{0, 1},
					{0, 1},
					{0, 1},
					{0, 1},
				},
			},
			want: 2,
		},
		{
			name: "Nobody want to go",
			args: args{
				n:      0,
				people: [][]int{},
			},
			want: 0,
		},
		{
			name: "Single person can go",
			args: args{
				n: 1,
				people: [][]int{
					{0, 10},
				},
			},
			want: 1,
		},
		{
			name: "Single person can't go",
			args: args{
				n: 1,
				people: [][]int{
					{2, 10},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxGroup(tt.args.n, tt.args.people); got != tt.want {
				t.Errorf("getMaxGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
