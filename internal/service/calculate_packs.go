package service

import (
	"context"
	"errors"
	"math"
)

func (s *PackCalculatorService) CalculatePacks(ctx context.Context, itemCount int) (map[int]int, int, error) {
	if itemCount <= 0 {
		return nil, 0, errors.New("itemCount must be greater than 0")
	}

	packSizes, err := s.packRepo.GetAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	bestPacks, minTotal := findOptimalPackCombination(itemCount, packSizes)
	if bestPacks == nil {
		return nil, 0, errors.New("no valid pack combination found")
	}

	return bestPacks, minTotal, nil
}

func findOptimalPackCombination(target int, packSizes []int) (map[int]int, int) {
	// This function finds the optimal combination of pack sizes to meet or exceed the target item count.
	// It uses a depth-first search (DFS) approach with memoization to explore all combinations of pack sizes.
	//
	// Returns:
	// - A map of pack sizes to their counts if a valid combination is found, or nil if not.

	// result is a struct to hold the total count and the packs used.
	// // It uses a map to store the count of each pack size used in the combination.
	type result struct {
		total int
		packs map[int]int
	}

	// memoization map to store already computed results for specific totals.
	memo := make(map[int]result)

	// dfs is a recursive function that explores all combinations of pack sizes.
	// It takes the current total as an argument and returns the best result found.
	var dfs func(currentTotal int) result
	dfs = func(currentTotal int) result {
		// If we have already computed this total, return the cached result.
		if val, ok := memo[currentTotal]; ok {
			return val
		}

		// If we've already passed the target, this is a potential candidate
		if currentTotal >= target {
			return result{
				total: currentTotal,
				packs: make(map[int]int),
			}
		}

		// best is initialized to a result with a very high total and nil packs.
		// This will be updated with the best valid combination found during the search.
		// It represents the best combination of packs found so far.
		// It starts with a total of math.MaxInt to ensure any valid combination will be better.
		// The packs map is initialized to nil, indicating no packs have been used yet.
		best := result{total: math.MaxInt, packs: nil}

		// Iterate through each pack size and recursively call dfs to explore further combinations.
		// For each pack size, we add it to the current total and check the result of that path.
		// If the result is valid (not nil), we check if it is better than the current best.
		for _, size := range packSizes {
			sub := dfs(currentTotal + size)
			if sub.packs == nil {
				continue
			}

			// Build new combination
			currentPacks := copyMap(sub.packs)
			currentPacks[size]++
			currentTotalWithThis := sub.total

			// Prefer the one with the smallest total ≥ target,
			// and among those, fewest packs
			if currentTotalWithThis < best.total ||
				(currentTotalWithThis == best.total && totalPacks(currentPacks) < totalPacks(best.packs)) {
				best.total = currentTotalWithThis
				best.packs = currentPacks
			}
		}

		// memoize the result for the current total
		// This helps avoid redundant calculations in future calls.
		memo[currentTotal] = best
		return best
	}

	// Start the DFS from a total of 0.
	// This will explore all combinations of pack sizes starting from zero.
	bestResult := dfs(0)

	// If the best result's packs are nil, it means no valid combination was found.
	if bestResult.packs == nil {
		return nil, 0
	}

	return bestResult.packs, bestResult.total
}

func totalPacks(packs map[int]int) int {
	// This function calculates the total number of packs used in the given map.

	count := 0
	for _, qty := range packs {
		count += qty
	}
	return count
}

func copyMap(original map[int]int) map[int]int {
	// This function creates a copy of the given map.
	// If the original map is nil, it returns an empty map.

	if original == nil {
		return map[int]int{}
	}
	copy := make(map[int]int)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}
