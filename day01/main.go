package main

func Part1() int {
	var r int

	for i := 1; i < len(Data); i++ {
		if Data[i] > Data[i-1] {
			r++
		}
	}

	return r
}

func Part2() int {
	var r int

	for i := 3; i < len(Data); i++ {
		if Data[i]+Data[i-1]+Data[i-2] > Data[i-1]+Data[i-2]+Data[i-3] {
			r++
		}
	}

	return r
}
