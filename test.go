package main

func main() {
	Append()
	AllocatedAppend()
	AllocatedAssign()
}

// Append appends to a slice without pre-allocating the length or capacity.
func Append() {
	s := make([]int, 0)
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
}

// AllocatedAppend appends to a slice with pre-allocated capacity.
func AllocatedAppend() {
	s := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
}

// AllocatedAssign assigns to a slice with pre-allocated length and capacity.
func AllocatedAssign() {
	// here when you specify the length and not the capacity, the capacity is set to be equal to the length.
	s := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		s[i] = i
	}
}
