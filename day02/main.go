package main

func Part1() int {
	var depth, pos int

	for i := 0; i < len(Data); i++ {
		switch Data[i][0] {
		case 0:
			pos += Data[i][1]
		case 1, -1:
			depth += Data[i][0] * Data[i][1]
		}
	}

	return depth * pos
}

func Part2() int {
	var depth, pos, aim int

	for i := 0; i < len(Data); i++ {
		switch Data[i][0] {
		case 0:
			pos += Data[i][1]
			depth += Data[i][1] * aim
		case 1, -1:
			aim += Data[i][0] * Data[i][1]
		}
	}

	return depth * pos
}
