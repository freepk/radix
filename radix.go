package radix

const (
	radix       = 8
	numCounters = 1 << radix
)

func Ints(source, buffer []int, size int) {
	var shift byte
	var counters [numCounters]int
	var pos, i, curr, prev int
	var num byte
	var sorted bool
	for shift = 0; shift < (intSize * radix); shift += radix {
		for i = 0; i < numCounters; i++ {
			counters[i] = 0
		}
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
			if (shift / radix % 2) == 1 {
				copy(buffer, source)
			}
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
