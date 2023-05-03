package domain

type UserType int8

const (
	Guest UserType = iota
	Host
	Admin
)

func (userType UserType) String() string {
	switch userType {
	case Guest:
		return "Guest"
	case Host:
		return "Host"
	case Admin:
		return "Admin"
	}
	return "Unknown"
}
type JwtData struct {
	UserId			string `bson:"_id"`
	UserType	float64 `bson:"user_type"`
	Username	string		`bson:"username"`
}