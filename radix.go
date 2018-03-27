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
	var pos, i int
	var num byte
	counters = make([]int, numCounters)
	for shift = 0; shift < (intSize * 8); shift = shift + 8 {
		copy(counters, emptyCounters)
		for i = 0; i < size; i++ {
			num = byte(source[i] >> shift)
			counters[num]++
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

func IntsEx(source []int) {
	var buffer []int
	var size int
	size = len(source)
	buffer = make([]int, size)
	Ints(source, buffer, size)
}
