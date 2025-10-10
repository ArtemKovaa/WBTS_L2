package main

import "testing"

func TestUnpackRLE(t *testing.T) {
    t.Run("error case", func(t *testing.T) {
		inputs := []string{"45", "\\2", "abcd\\"} 
		
		for _, input := range inputs {
			_, err := UnpackRLE(input)

			if err == nil {
				t.Fatalf("UnpackRLE(%s) did not ", err)
			}
		}

        result, err := DoSomething(1)
        if err != nil {
            t.Fatalf("Ожидалась ошибка nil, получили %v", err)
        }
        if result != "ok" {
            t.Errorf("Ожидалось 'ok', получили %s", result)
        }
    })

    t.Run("happy path", func(t *testing.T) {
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
				t.Errorf("UnpackRLE(%s) = %v; want %s", tt.input, err, tt.want)
			}

			if got != tt.want {
				t.Errorf("UnpackRLE(%s) = %s; want %s", tt.input, got, tt.want)
        	}
   		}
    })
}