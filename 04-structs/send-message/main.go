package main

func main() {}

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

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	}

	return "", false
}

func newUser(name string, membershipType membershipType) User {
	membership := Membership{Type: membershipType}

	if membership.Type == TypeStandard {
		membership.MessageCharLimit = 100
	} else if membership.Type == TypePremium {
		membership.MessageCharLimit = 1000
	}

	return User{Name: name, Membership: membership}
}
