package problem_1386

import (
	"sort"
)

type Bucket struct {
	Row  int
	Low  int
	High int
}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	result := 0

	sort.Slice(reservedSeats, func(p, q int) bool {
		return reservedSeats[p][0] < reservedSeats[q][0]
	})

	if reservedSeats[0][0] > 1 {
		result += 2 * (reservedSeats[0][0] - 1)
	}

	if reservedSeats[len(reservedSeats)-1][0] < n {
		result += 2 * (n - reservedSeats[len(reservedSeats)-1][0])
	}

	rowData := Bucket{
		Row:  reservedSeats[0][0],
		Low:  0,
		High: 11,
	}

	for _, reservedSeat := range reservedSeats {
		if reservedSeat[0] != rowData.Row {
			result += countAvailableVariants(rowData)

			if reservedSeat[0] > rowData.Row+1 {
				result += 2 * (reservedSeat[0] - rowData.Row - 1)
			}

			rowData = Bucket{
				Row:  reservedSeat[0],
				Low:  0,
				High: 11,
			}
		}

		if reservedSeat[1] >= rowData.Low && reservedSeat[1] <= rowData.High {
			if reservedSeat[1] >= 6 {
				rowData.High = reservedSeat[1]
			} else {
				rowData.Low = reservedSeat[1]
			}
		}
	}

	return result + countAvailableVariants(rowData)
}

func countAvailableVariants(bucket Bucket) int {
	if bucket.Row == 0 {
		return 0
	}

	switch {
	case bucket.Low <= 1 && bucket.High >= 10:
		return 2
	case bucket.Low <= 1 && bucket.High >= 6:
		return 1
	case bucket.Low <= 5 && bucket.High >= 10:
		return 1
	case bucket.Low <= 3 && bucket.High >= 8:
		return 1
	default:
		return 0
	}
}
