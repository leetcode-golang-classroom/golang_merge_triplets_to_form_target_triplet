package sol

func mergeTriplets(triplets [][]int, target []int) bool {
	possibles := make([]bool, 3)
	for _, triplet := range triplets {
		if triplet[0] <= target[0] && triplet[1] <= target[1] && triplet[2] <= target[2] {
			for idx := range possibles {
				if !possibles[idx] && triplet[idx] == target[idx] {
					possibles[idx] = true
				}
			}
			if possibles[0] && possibles[1] && possibles[2] {
				return true
			}
		}
	}
	return false
}
