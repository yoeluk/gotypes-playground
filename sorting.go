package main

func Sorted[O Ordering[S], S any](o O, arr []S) []S {
	var sort func(a []S) []S
	sort = func(a []S) []S {
		if len(a) <= 1 {
			return a
		}
		first := append(make([]S, 0, 0), a[:len(a)/2]...)
		second := append(make([]S, 0, 0), a[len(a)/2:]...)
		return MergeSorted(o, sort(first), sort(second))
	}
	return sort(arr)
}

func MergeSorted[O Ordering[S], S any](o O, arr1 []S, arr2 []S) []S {
	if len(arr1) == 0 {
		return arr2
	} else if len(arr2) == 0 {
		return arr1
	} else {
		a1 := append(make([]S, 0, 0), arr1...)
		a2 := append(make([]S, 0, 0), arr2...)
		if o.Lt(a1[0], a2[0]) || o.Equiv(a1[0], a2[0]) {
			return append(a1[:1], MergeSorted(o, a1[1:], a2)...)
		} else {
			return append(a2[:1], MergeSorted(o, a1, a2[1:])...)
		}
	}
}
