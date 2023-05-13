package domain

type Role int

const (
	GUEST Role = iota
	HOST
	ADMIN
)

func (role Role) String() string {
	switch role {
	case GUEST:
		return "Guest"
	case HOST:
		return "Host"
	case ADMIN:
		return "Admin"
	}
	return "Unknown"
}

type JwtClaims struct {
	UserId   string `json:"id"`
	Role     Role   `json:"role"`
	Username string `json:"string"`
}
