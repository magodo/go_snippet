package main

func QSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	l, r := 0, len(arr)-1
	var sp int
	for {
		for l < r && arr[l] < arr[len(arr)-1] {
			l++
			sp = l
		}
		for l < r && arr[r] >= arr[len(arr)-1] {
			r--
		}
		if l == r {
			break
		}
		swap(arr, l, r)
		sp = r
	}
	swap(arr, sp, len(arr)-1)

	QSort(arr[:sp])
	QSort(arr[sp+1 : len(arr)])
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
