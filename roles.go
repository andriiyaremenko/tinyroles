package tinyroles

import "sync"

type Role string

type Roles struct {
	mu    sync.RWMutex
	roles map[Role]uint64
}

// Checks if Role has a Permission
// Returns `false` for Role without assigned permissions
func (rf *Roles) HasPermission(role Role, permission Permission) bool {
	rf.mu.RLock()
	defer rf.mu.RUnlock()

	rv, ok := rf.roles[role]

	return ok && ((rv & permission.Value()) == permission.Value())
}

// Assigns set of Permissions to a Role
func (rf *Roles) AssignPermissions(role Role, permissions ...Permission) {
	var v uint64 = 0
	for _, p := range permissions {
		v = v | p.Value()
	}

	rf.mu.RLock()

	if oldV, ok := rf.roles[role]; ok {
		v = v | oldV
	}

	rf.mu.RUnlock()

	rf.mu.Lock()

	if rf.roles == nil {
		rf.roles = make(map[Role]uint64)
	}

	rf.roles[role] = v

	rf.mu.Unlock()
}

// Withdraws set of Permissions from a Role
func (rf *Roles) WithdrawPermissions(role Role, permissions ...Permission) {
	for _, p := range permissions {
		if rf.HasPermission(role, p) {
			rf.mu.Lock()

			rf.roles[role] = rf.roles[role] ^ p.Value()

			rf.mu.Unlock()
		}
	}
}
