package tinyroles

// Creates Permission with Value of 2^`base`
func NewPermission(base uint64) Permission {
	return permission(1 << base)
}

type Permission interface {
	sealed()
	Value() uint64
}

type permission uint64

func (p permission) Value() uint64 {
	return uint64(p)
}

func (p permission) sealed() {}
