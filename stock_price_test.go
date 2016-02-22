package stock_price

import "testing"

func Testget_max_profit(t *testing.T)
{
	cases := []struct {
		in int[]
		want int
	}{
		{[10, 7, 5, 8, 11, 9], 6}
	}
	for _, c := range cases {
		got := get_max_profit(c.in)
		if got != c.want {
			t.Errorf("get_max_profit(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}