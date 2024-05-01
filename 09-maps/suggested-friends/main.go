package main

import "fmt"

func main() {
	friendships := map[string][]string{
		"Dhalinar": {"Kaladin", "Pattern"},
		"Kaladin":  {"Dhalinar", "Syl", "Teft"},
		"Pattern":  {"Dhalinar", "Teft"},
		"Syl":      {"Kaladin"},
		"Teft":     {"Kaladin", "Pattern"},
		"Moash":    {},
		"Shallan":  {"Pattern", "Kaladin"},
	}

	// findSuggestedFriends("Pattern", friendships)
	fmt.Println(findSuggestedFriends("Pattern", friendships))
}

func findSuggestedFriends(username string, friednships map[string][]string) []string {
	directFriends := friednships[username]
	filterFriends := make(map[string]bool)
	suggestedFriends := make([]string, 0)

	for user, friends := range friednships {
		if user == username {
			continue
		}

		for _, d := range directFriends {
			if d == user {
				for _, f := range friends {
					if f != username {
						filterFriends[f] = true
					}
				}
			}
		}
	}

	for friend := range filterFriends {
		suggestedFriends = append(suggestedFriends, friend)
	}

	if len(suggestedFriends) == 0 {
		return nil
	}

	return suggestedFriends
}
