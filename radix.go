package radix

const (
	numCounters = 256
)

var (
	emptyCounters []int
)

func init() {
	emptyCounters = make([]int, numCounters)
}

func Ints(source, buffer []int, size int) {
	var shift byte
	var counters []int
	var pos, i, curr, prev int
	var num byte
	var sorted bool
	counters = make([]int, numCounters)
	for shift = 0; shift < (intSize * 8); shift += 8 {
		copy(counters, emptyCounters)
		sorted = true
		prev = source[0]
		for i = 0; i < size; i++ {
			curr = source[i]
			sorted = sorted && (prev <= curr)
			num = byte(curr >> shift)
			counters[num]++
			prev = curr
		}
		if sorted {
			return
		}
		pos = 0
		for i = 0; i < numCounters; i++ {
			pos, counters[i] = (pos + counters[i]), pos
		}
		for i = 0; i < size; i++ {
			num = byte(source[i] >> shift)
			pos = counters[num]
			buffer[pos] = source[i]
			counters[num]++
		}
		source, buffer = buffer, source
	}
}
