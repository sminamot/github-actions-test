package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		x      int
		y      int
		expect int
	}{
		{
			name:   "1+1",
			x:      1,
			y:      1,
			expect: 2,
		},
		{
			name:   "-5+10",
			x:      -5,
			y:      10,
			expect: 5,
		},
		{
			name:   "-5+(-5)",
			x:      -5,
			y:      -5,
			expect: -10,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if ret := Add(tt.x, tt.y); ret != tt.expect {
				t.Errorf("got:%d, want:%d", ret, tt.expect)
			}
		})
	}
}
