package main

import "testing"

func TestUnpackRLE(t *testing.T) {
    t.Run("happy path", func(t *testing.T) {
        result, err := DoSomething(1)
        if err != nil {
            t.Fatalf("Ожидалась ошибка nil, получили %v", err)
        }
        if result != "ok" {
            t.Errorf("Ожидалось 'ok', получили %s", result)
        }
    })

    t.Run("error case", func(t *testing.T) {
		var tests = []struct {
			s string
			want string
		}{
			{"a4bc2d5e", "aaaabccddddde"},
			{"abcd", "abcd"},
			{"", ""},
			{"qwe\\4\\5", "qwe45"},
			{"qwe\45"},
		}

        _, err := DoSomething(-1)
        if err == nil {
            t.Fatal("Ожидалась ошибка, получили nil")
        }
        expected := "недопустимое значение"
        if err.Error() != expected {
            t.Errorf("Ожидалась ошибка %q, получили %q", expected, err.Error())
        }
    })



    for _, tt := range tests {
        got := IntMin(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("IntMin(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}