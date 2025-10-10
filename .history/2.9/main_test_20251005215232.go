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
			input string
			want string
		}{
			{"a4bc2d5e", "aaaabccddddde"},
			{"abcd", "abcd"},
			{"", ""},
			{"qwe\\4\\5", "qwe45"},
			{"qwe\\45", "qwe44444"},
		}

		for _, tt := range tests {
			got, err := UnpackRLE(tt.input)
			if err != nil {
				
			}

			if got != tt.want || err != nil {
				t.Errorf("UnpackRLE(%s) = %s; want %s", tt.input, got, tt.want)
        }
    }

        _, err := UnpackRLE(-1)
        if err == nil {
            t.Fatal("Ожидалась ошибка, получили nil")
        }
        expected := "недопустимое значение"
        if err.Error() != expected {
            t.Errorf("Ожидалась ошибка %q, получили %q", expected, err.Error())
        }
    })




}