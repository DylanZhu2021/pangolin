package sort

import (
	"math/bits"
	"math/rand"
	"strconv"
)

// 一个满足sort.Interface的类型，通常是一个集合，可以被这个包中的程序排序。
// 通过这个包中的例程进行排序。这些方法要求
// 集合的元素被一个整数索引所列举。
type Interface interface {
	// Len 是集合中元素的数量。
	Len() int

	// Less 报告索引为i的元素是否应该排序在索引为j的元素之前。
	// 索引i的元素是否应该排序在索引j的元素之前。
	Less(i, j int) bool

	// Swap 交换索引为i和j的元素。
	Swap(i, j int)

	ShiftTail(i, j int)
	ShiftHead(i, j int)
	CyclicSwaps(is, js []int)
}

const (
	SHORTEST_MEDIAN_OF_MEDIANS = 50
	MAX_SWAPS                  = 4 * 3
	BLOCK                      = 128
	MAX_STEPS                  = 5
	SHORTEST_SHIFTING          = 50
	MAX_INSERTION              = 20
)

var offsetsL [BLOCK]int
var offsetsR [BLOCK]int

func partialInsertionSort(data Interface, a, b int) bool {

	len := b - a
	i := 1
	for k := 0; k < MAX_STEPS; k += 1 {
		for i < len && !data.Less(a+i, a+i-1) {
			i += 1
		}

		if i == len {
			return true
		}

		if len < SHORTEST_SHIFTING {
			return false
		}

		data.Swap(a+i-1, a+i)

		data.ShiftTail(a, a+i)
		data.ShiftHead(a+i, b)
	}

	return false
}

func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		data.ShiftTail(0, i+1)
	}
}

func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	//建立堆，最大的元素在顶部。
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	//将最大的元素弹出，放到数据的末端。
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}

func partitionInBlock(data Interface, a, b, pivot int) int {
	l := a
	blockL := BLOCK
	startL := 0
	endL := 0

	r := b
	blockR := BLOCK
	startR := 0
	endR := 0

	for {
		isDone := (r - l) <= 2*BLOCK

		if isDone {
			rem := r - l
			if startL < endL || startR < endR {
				rem -= BLOCK
			}

			if startL < endL {
				blockR = rem
			} else if startR < endR {
				blockL = rem
			} else {
				blockL = rem / 2
				blockR = rem - blockL
			}
		}

		if startL == endL {
			startL = 0
			endL = 0
			elem := l

			for i := 0; i < blockL; i += 1 {
				offsetsL[endL] = l + i

				if !data.Less(elem, pivot) {
					endL += 1
				}

				elem += 1
			}
		}

		if startR == endR {
			startR = 0
			endR = 0
			elem := r

			for i := 0; i < blockR; i += 1 {
				elem -= 1
				offsetsR[endR] = r - i - 1

				if data.Less(elem, pivot) {
					endR += 1
				}
			}
		}

		count := min(endL-startL, endR-startR)
		if count > 0 {
			data.CyclicSwaps(offsetsL[startL:(startL+count)], offsetsR[startR:(startR+count)])
			startL += count
			startR += count
		}

		if startL == endL {
			l += blockL
		}

		if startR == endR {
			r -= blockR
		}

		if isDone {
			break
		}
	}

	if startL < endL {
		for startL < endL {
			endL -= 1
			data.Swap(offsetsL[endL], r-1)
			r -= 1
		}
		return (r - a)
	} else if startR < endR {
		for startR < endR {
			endR -= 1
			data.Swap(l, offsetsR[endR])
			l += 1
		}
		return (l - a)
	} else {
		return (l - a)
	}
}

func partition(data Interface, a, b, pivot int) (int, bool) {
	data.Swap(a, pivot)
	pivot = a

	l := a + 1
	r := b
	for l < r && data.Less(l, pivot) {
		l += 1
	}

	for l < r && !data.Less(r-1, pivot) {
		r -= 1
	}

	numOfSmallerThanPivotElems := partitionInBlock(data, l, r, pivot)
	mid := (l - 1 + numOfSmallerThanPivotElems)
	wasPartitioned := (l >= r)

	data.Swap(a, mid)
	return mid, wasPartitioned
}

func partitionEqual(data Interface, a, b, pivot int) int {
	data.Swap(a, pivot)
	pivot = a

	l := a + 1
	r := b
	for {
		for l < r && !data.Less(pivot, l) {
			l += 1
		}

		for l < r && data.Less(pivot, r-1) {
			r -= 1
		}

		if l >= r {
			break
		}

		r -= 1
		data.Swap(l, r)
		l += 1
	}

	return l
}

func breakPatterns(data Interface, a, b int) {
	len := b - a
	if len >= 8 {
		var shift uint = uint(strconv.IntSize - bits.LeadingZeros(uint(len)))
		var nextPowerOfTwo uint = 1 << shift

		modulus := nextPowerOfTwo
		pos := a + (len / 4 * 2)

		for i := 0; i < 3; i += 1 {
			var gen uint = uint(rand.Int())

			other := int(gen & (modulus - 1))
			if other >= len {
				other -= len
			}
			other += a

			data.Swap(pos-1+i, other)
		}
	}
}

func reverseRange(data Interface, a, b int) {
	i := a
	j := b - 1
	for i < j {
		data.Swap(i, j)
		i += 1
		j -= 1
	}
}

func sort2(data Interface, a, b, swaps *int) {
	if data.Less(*b, *a) {
		t := *b
		*b = *a
		*a = t

		*swaps += 1
	}
}

func sort3(data Interface, a, b, c, swaps *int) {
	sort2(data, a, b, swaps)
	sort2(data, b, c, swaps)
	sort2(data, a, b, swaps)
}

func sortAdajacent(data Interface, a, swaps *int) {
	t := *a
	t_minus_one := t - 1
	t_plus_one := t + 1
	sort3(data, &t_minus_one, a, &t_plus_one, swaps)
}

func choosePivot(data Interface, x, y int) (pivot int, likelySorted bool) {
	len := y - x

	a := len / 4 * 1
	b := len / 4 * 2
	c := len / 4 * 3

	swaps := 0

	if len >= 8 {
		if len >= SHORTEST_MEDIAN_OF_MEDIANS {
			sortAdajacent(data, &a, &swaps)
			sortAdajacent(data, &b, &swaps)
			sortAdajacent(data, &c, &swaps)
		}

		sort3(data, &a, &b, &c, &swaps)
	}

	if swaps < MAX_SWAPS {
		return x + b, (swaps == 0)
	} else {
		reverseRange(data, a, b)
		return x + (len - 1 - b), true
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func recurse(data Interface, a, b, pred, limit int) {
	wasBalanced := true
	wasPartitioned := true

	for {
		len := b - a
		if len <= MAX_INSERTION {
			insertionSort(data, a, b)
			return
		}

		if limit == 0 {
			heapSort(data, a, b)
			return
		}

		if !wasBalanced {
			breakPatterns(data, a, b)
			limit -= 1
		}

		pivot, likelySorted := choosePivot(data, a, b)
		if wasBalanced && wasPartitioned && likelySorted {
			if partialInsertionSort(data, a, b) {
				return
			}
		}

		if pred > 0 {
			if !data.Less(pred, pivot) {
				mid := partitionEqual(data, a, b, pivot)
				a = mid
				continue
			}
		}

		mid, wasP := partition(data, a, b, pivot)
		wasBalanced = min(mid-a, len-(mid-a)) >= (len / 8)
		wasPartitioned = wasP

		left_len := mid - a
		right_len := len - (mid - a) - 1
		if left_len < right_len {
			recurse(data, a, mid, pred, limit)
			a = mid + 1
			pred = mid
		} else {
			recurse(data, mid+1, b, mid, limit)
			b = mid
		}
	}
}

func quickSort(data Interface, a, b int) {
	n := data.Len()
	if n == 0 {
		return
	}

	limit := strconv.IntSize - bits.LeadingZeros(uint(n))
	pred := -1
	recurse(data, a, b, pred, limit)
}

// Sort 排序数据.
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n)
}

// 适用于常见情况的便利类型

// IntSlice 将Interface的方法附加到[]int上，按照递增的顺序进行排序。
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p IntSlice) ShiftTail(a, b int) {
	len := b - a
	if len >= 2 && p[b-1] < p[b-2] {
		tmp := p[b-1]
		p[b-1] = p[b-2]

		i := b - 3
		for i >= 0 {
			if !(tmp < p[i]) {
				break
			}

			p[i+1] = p[i]
			i -= 1
		}

		p[i+1] = tmp
	}
}

func (p IntSlice) ShiftHead(a, b int) {
	len := b - a
	if len >= 2 && p[a+1] < p[a] {
		tmp := p[a]
		p[a] = p[a+1]

		i := a + 2
		for i < len {
			if !(p[i] < tmp) {
				break
			}

			p[i-1] = p[i]
			i += 1
		}

		p[i-1] = tmp
	}
}

func (p IntSlice) CyclicSwaps(is, js []int) {
	count := len(is)
	tmp := p[is[0]]
	p[is[0]] = p[js[0]]

	for i := 1; i < count; i += 1 {
		p[js[i-1]] = p[is[i]]
		p[is[i]] = p[js[i]]
	}

	p[js[count-1]] = tmp
}

func (p IntSlice) Sort() { Sort(p) }

// Float64Slice 将Interface的方法附加到[]float64上，按照递增的顺序排序。
// (非数字的数值被视为小于其他数值)。
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Float64Slice) ShiftTail(a, b int) {
	len := b - a
	if len >= 2 {
		for i := len - 1; i >= 1; i -= 1 {
			if !p.Less(i, i-1) {
				break
			}

			p.Swap(i, i-1)
		}
	}
}

func (p Float64Slice) ShiftHead(a, b int) {
	len := b - a
	if len >= 2 {
		for i := 1; i < len; i += 1 {
			if !p.Less(i, i-1) {
				break
			}

			p.Swap(i, i-1)
		}
	}
}

func (p Float64Slice) CyclicSwaps(is, js []int) {
	count := len(is)
	tmp := p[is[0]]
	p[is[0]] = p[js[0]]

	for i := 1; i < count; i += 1 {
		p[js[i-1]] = p[is[i]]
		p[is[i]] = p[js[i]]
	}

	p[js[count-1]] = tmp
}

// isNaN 是math.IsNaN的一个副本，以避免对数学包的依赖。
func isNaN(f float64) bool {
	return f != f
}

func (p Float64Slice) Sort() { Sort(p) }

// StringSlice 将Interface的方法附加到[]string上，按照递增的顺序进行排序。
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StringSlice) ShiftTail(a, b int) {
	len := b - a
	if len >= 2 {
		for i := len - 1; i >= 1; i -= 1 {
			if !p.Less(i, i-1) {
				break
			}

			p.Swap(i, i-1)
		}
	}
}

func (p StringSlice) ShiftHead(a, b int) {
	len := b - a
	if len >= 2 {
		for i := 1; i < len; i += 1 {
			if !p.Less(i, i-1) {
				break
			}

			p.Swap(i, i-1)
		}
	}
}

func (p StringSlice) CyclicSwaps(is, js []int) {
	count := len(is)
	tmp := p[is[0]]
	p[is[0]] = p[js[0]]

	for i := 1; i < count; i += 1 {
		p[js[i-1]] = p[is[i]]
		p[is[i]] = p[js[i]]
	}

	p[js[count-1]] = tmp
}

func (p StringSlice) Sort() { Sort(p) }

// 常见情况下的便利包装器

// Ints 按照递增的顺序对一个片状图进行排序。
func Ints(a []int) { Sort(IntSlice(a)) }

// Float64s 按照递增的顺序对一个浮动64的片断进行排序
// (非数字的数值被视为小于其他数值)。
func Float64s(a []float64) { Sort(Float64Slice(a)) }

// Strings 按照递增的顺序对一个字符串切片进行排序。
func Strings(a []string) { Sort(StringSlice(a)) }

// IntsAreSorted 测试一个整数片是否按递增顺序排序。
func IntsAreSorted(a []int) bool { return IsSorted(IntSlice(a)) }

// Float64sAreSorted 测试一个float64s的片断是否按递增顺序排序
// (非数字的数值被视为小于其他数值)。
func Float64sAreSorted(a []float64) bool { return IsSorted(Float64Slice(a)) }

// StringsAreSorted 测试一个字符串片断是否按递增顺序排序。
func StringsAreSorted(a []string) bool { return IsSorted(StringSlice(a)) }

// IsSorted 报告数据是否被排序。
func IsSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
