package client

import "testing"

func TestStore(t *testing.T) {

    var cl ClientStore
    cl.Storage = make(map[string][]byte)
    want := "Banana"

    key, _ := cl.Store([]byte("123"), []byte(want))
    content, _ := cl.Retrieve([]byte("123"), key)

    if string(content) != want {
    	t.Errorf("content was %q, want %q", content, want)
    }
}