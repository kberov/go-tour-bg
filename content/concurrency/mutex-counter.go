// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter може да се ползва едновременно от много го-нишки.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc ужеличава стойността, към която сочи даден ключ.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Заключваме, така че само една нишка може да достъпва речника c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value връща текущата стойност за даден ключ.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Заключваме, така че само една нишка може да достъпва речника c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
