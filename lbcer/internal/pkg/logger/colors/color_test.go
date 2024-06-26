package colors

import (
	"testing"
)

func TestColorize(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		color Color
		want  string
	}{
		{
			name:  "magenta",
			s:     "test",
			color: Magenta,
			want:  "\033[35mtest\033[0m",
		},
		{
			name:  "blue",
			s:     "test",
			color: Blue,
			want:  "\033[34mtest\033[0m",
		},
		{
			name:  "yellow",
			s:     "test",
			color: Yellow,
			want:  "\033[33mtest\033[0m",
		},
		{
			name:  "red",
			s:     "test",
			color: Red,
			want:  "\033[31mtest\033[0m",
		},
		{
			name:  "cyan",
			s:     "test",
			color: Cyan,
			want:  "\033[36mtest\033[0m",
		},
		{
			name:  "white",
			s:     "test",
			color: White,
			want:  "\033[37mtest\033[0m",
		},
		{
			name:  "default",
			s:     "test",
			color: 10,
			want:  "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Colorize(tt.s, tt.color); got != tt.want {
				t.Errorf("Colorize() = %v, want %v", got, tt.want)
			}
		})
	}
}
