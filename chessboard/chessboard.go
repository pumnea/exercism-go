package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if f, exists := cb[file]; exists {
		count := 0
		for _, occupied := range f {
			if occupied {
				count++
			}
		}
		return count
	}
	return 0
}

// C untInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank > len(cb) || rank < 1 {
		return 0
	}
	count := 0
	for _, v := range cb {
		if v[rank-1] {
			count++
		}
	}
	return count
}

// C untAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// C untOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for _, occupied := range v {
			if occupied {
				count++
			}
		}
	}
	return count
}
