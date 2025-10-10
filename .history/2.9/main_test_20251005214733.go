package main

import "testing"

func TestUnpackRLE(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {
        got := IntMin(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("IntMin(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}