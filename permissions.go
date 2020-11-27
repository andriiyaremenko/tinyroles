package tinyroles

// Creates Permission with Value of 2^`base`
func NewPermission(base uint) Permission {
	return permission(1 << base)
}

// Permissions interface
// It is sealed to ensure no calculation rules was violated when `Value()` is calculated
type Permission interface {
	sealed()

	// Returns permissions value
	Value() uint64
}

type permission uint64

func (p permission) Value() uint64 {
	return uint64(p)
}

func (p permission) sealed() {}
