package main

import (
	"fmt"
	"sort"
	"strings"
)

type Visit struct {
	username  string
	website   string
	timestamp int
}

func generateCombination(websites []string) []string {
	var combos []string
	n := len(websites)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				pattern := websites[i] + " " + websites[j] + " " + websites[k]
				combos = append(combos, pattern)
			}
		}
	}
	return combos
}
func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	var visits []Visit
	for i := range username {
		visits = append(visits, Visit{username[i], website[i], timestamp[i]})
	}
	sort.Slice(visits, func(i, j int) bool {
		return visits[i].timestamp < visits[j].timestamp
	})
	userVisits := make(map[string][]string)
	for _, visit := range visits {
		userVisits[visit.username] = append(userVisits[visit.username], visit.website)
	}
	patternUsers := make(map[string]map[string]bool)
	for user, websites := range userVisits {
		seen := make(map[string]bool)
		combos := generateCombination(websites)
		for _, pattern := range combos {
			if !seen[pattern] {
				if patternUsers[pattern] == nil {
					patternUsers[pattern] = make(map[string]bool)
				}
				patternUsers[pattern][user] = true
				seen[pattern] = true
			}
		}
	}
	maxCount := 0
	result := ""
	for pattern, users := range patternUsers {
		count := len(users)
		if count > maxCount || (count == maxCount && pattern < result) {
			result = pattern
			maxCount = count
		}
	}
	return strings.Split(result, " ")
}
func main() {
	username := []string{"bob", "bob", "bob", "alice", "alice", "alice", "alice", "charlie", "charlie", "charlie"}
	timestamp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	website := []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"}
	result := mostVisitedPattern(username, timestamp, website)
	fmt.Println(result) // Output: ["home", "about", "career"]
}
