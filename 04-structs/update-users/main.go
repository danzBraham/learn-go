package main

import "fmt"

func main() {
	userZidan := newUser("Zidan", TypePremium)
	fmt.Println(userZidan)
}

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, memmembershipType membershipType) User {
	membership := Membership{Type: memmembershipType}

	if membership.Type == TypeStandard {
		membership.MessageCharLimit = 100
	} else if membership.Type == TypePremium {
		membership.MessageCharLimit = 1000
	}

	return User{Name: name, Membership: membership}
}
