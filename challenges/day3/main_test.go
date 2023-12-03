package main

import "testing"

func Test_solvePart1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example input should be 4765",
			args{
				lines: []string{
					"467..114..",
					"...*......",
					"..35..633.",
					"......#...",
					"617*......",
					".....+.58.",
					"..592.....",
					"......755.",
					"...$.*....",
					".664.598..",
					"..........",
					"......$404",
				},
			}, 4765,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.lines); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
