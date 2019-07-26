package counter

const chunkSize = 60

type CounterService interface {
	Increment()
	Add()
	Count() int
	Get() []int
}

// c -> requests in seconds
type counter struct {
	c []int
}

func (c *counter) Increment() {
	c.c[len(c.c)-1] += 1
	return
}

func (c *counter) Add() {
	c.c = c.c[1:]
	c.c = append(c.c, 0)
}

// sum all int in slice
func (c counter) Count() int {
	sum := 0
	for _, val := range c.c {
		sum += val
	}
	return sum
}

func (c counter) Get() []int {
	return c.c
}

func NewCounter(arr []int) CounterService {
	ctr := &counter{
		c: arr[0:chunkSize],
	}
	return ctr
}
