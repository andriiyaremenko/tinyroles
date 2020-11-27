package tinyroles

import "sync"

// Role type
type Role string

// Roles manager. Use it to assign, withdraw and check permissions
type Roles struct {
	mu    sync.RWMutex
	roles map[Role]uint64
}

// Checks if a `role` has a `permission`
// Returns `false` for `role` without assigned permissions
func (rf *Roles) HasPermission(role Role, permission Permission) bool {
	rf.mu.RLock()
	defer rf.mu.RUnlock()

	rv, ok := rf.roles[role]

	return ok && ((rv & permission.Value()) == permission.Value())
}

// Assigns set of `permissions` to a `role`
// It is idempotent - subsequent calls with already assigned roles will take no effect
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

// Withdraws set of `permissions` from a `role`
// It takes into account only assigned `permissions`
// Permissions that was not assigned to this `role` would be ignored
func (rf *Roles) WithdrawPermissions(role Role, permissions ...Permission) {
	rf.mu.RLock()

	if _, ok := rf.roles[role]; !ok {
		rf.mu.RUnlock()

		return
	}

	rf.mu.RUnlock()

	rf.mu.Lock()
	for _, p := range permissions {
		rf.roles[role] = rf.roles[role] &^ p.Value()
	}
	rf.mu.Unlock()
}

// Returns value of a `role`
func (rf *Roles) GetRoleValue(role Role) uint64 {
	rf.mu.RLock()
	defer rf.mu.RUnlock()

	if rv, ok := rf.roles[role]; ok {
		return rv
	}

	return 0
}
