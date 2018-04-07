package hello

import (
    "testing"
)

func TestHello(t *testing.T) {
    result := hello("fujita")

    if result != "hello, fujita" {
        t.Fatal("failed test")
    }
}
