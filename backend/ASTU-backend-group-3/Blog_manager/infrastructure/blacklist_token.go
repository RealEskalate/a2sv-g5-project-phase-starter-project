package infrastructure

import "fmt"

type TokenBlacklist struct {
	blacklist map[string]bool
}

var Blacklist = &TokenBlacklist{
	blacklist: make(map[string]bool),
}

// AddToBlacklist adds a token to the blacklist.
func (tb *TokenBlacklist) AddToBlacklist(token string) {
	fmt.Println("adding to blacklist" + token)
	tb.blacklist[token] = true
}

// IsTokenBlacklisted checks if a token is blacklisted.
func (tb *TokenBlacklist) IsTokenBlacklisted(token string) bool {
	fmt.Println("checking blacklist" + token)
	_, exists := tb.blacklist[token]
	return exists
}
