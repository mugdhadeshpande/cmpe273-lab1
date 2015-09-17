package main

import "testing"

func TestSleep(t *testing.T) {
  cache.nextPrune = time.Now().Add(-1 * time.Second)
cache.Get("fail get triggers a prune")
time.Sleep(5 * time.Millisecond)
}
