package client

import "testing"

func TestEasyStore(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "hello"},
	}

        var cl ClientDesc

	for _, c := range cases {
                cl.Z = c.in
		got := cl.EasyStore()

		if got != c.want {
			t.Errorf("cl.EasyStore(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}