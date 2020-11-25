package tinyroles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	permission0 = NewPermission(1)
	permission1 = NewPermission(2)
	permission2 = NewPermission(3)
	permission3 = NewPermission(4)
	permission4 = NewPermission(5)
	permission5 = NewPermission(6)
	permission6 = NewPermission(7)
	permission7 = NewPermission(8)
)

const (
	Role1 Role = "Role1"
	Role2 Role = "Role2"
)

func TestHasPermissionsWorks(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)
	roles.AssignPermissions(Role1, permission0, permission2, permission4, permission6)
	roles.AssignPermissions(Role2, permission0, permission1, permission3, permission5, permission7)

	assert.True(roles.HasPermission(Role1, permission0))
	assert.True(roles.HasPermission(Role1, permission2))
	assert.True(roles.HasPermission(Role1, permission4))
	assert.True(roles.HasPermission(Role1, permission6))

	assert.True(roles.HasPermission(Role2, permission0))
	assert.True(roles.HasPermission(Role2, permission1))
	assert.True(roles.HasPermission(Role2, permission3))
	assert.True(roles.HasPermission(Role2, permission5))
	assert.True(roles.HasPermission(Role2, permission7))

	assert.False(roles.HasPermission(Role2, permission2))
	assert.False(roles.HasPermission(Role2, permission4))
	assert.False(roles.HasPermission(Role2, permission6))

	assert.False(roles.HasPermission(Role1, permission1))
	assert.False(roles.HasPermission(Role1, permission3))
	assert.False(roles.HasPermission(Role1, permission5))
	assert.False(roles.HasPermission(Role1, permission7))
}

func TestHasPermissionReturnsFalseForRoleWithoutPermissions(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)

	assert.False(roles.HasPermission(Role1, permission0))
	assert.False(roles.HasPermission(Role1, permission2))
	assert.False(roles.HasPermission(Role1, permission4))
	assert.False(roles.HasPermission(Role1, permission6))

	assert.False(roles.HasPermission(Role2, permission0))
	assert.False(roles.HasPermission(Role2, permission1))
	assert.False(roles.HasPermission(Role2, permission3))
	assert.False(roles.HasPermission(Role2, permission5))
	assert.False(roles.HasPermission(Role2, permission7))

	assert.False(roles.HasPermission(Role2, permission2))
	assert.False(roles.HasPermission(Role2, permission4))
	assert.False(roles.HasPermission(Role2, permission6))

	assert.False(roles.HasPermission(Role1, permission1))
	assert.False(roles.HasPermission(Role1, permission3))
	assert.False(roles.HasPermission(Role1, permission5))
	assert.False(roles.HasPermission(Role1, permission7))
}

func TestAssignPermissionsInSubsequentCalls(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)

	roles.AssignPermissions(Role1, permission4, permission6)
	roles.AssignPermissions(Role1, permission0, permission2)

	roles.AssignPermissions(Role2, permission0, permission1)
	roles.AssignPermissions(Role2, permission3)
	roles.AssignPermissions(Role2, permission5, permission7)

	assert.True(roles.HasPermission(Role1, permission0))
	assert.True(roles.HasPermission(Role1, permission2))
	assert.True(roles.HasPermission(Role1, permission4))
	assert.True(roles.HasPermission(Role1, permission6))

	assert.True(roles.HasPermission(Role2, permission0))
	assert.True(roles.HasPermission(Role2, permission1))
	assert.True(roles.HasPermission(Role2, permission3))
	assert.True(roles.HasPermission(Role2, permission5))
	assert.True(roles.HasPermission(Role2, permission7))

	assert.False(roles.HasPermission(Role2, permission2))
	assert.False(roles.HasPermission(Role2, permission4))
	assert.False(roles.HasPermission(Role2, permission6))

	assert.False(roles.HasPermission(Role1, permission1))
	assert.False(roles.HasPermission(Role1, permission3))
	assert.False(roles.HasPermission(Role1, permission5))
	assert.False(roles.HasPermission(Role1, permission7))
}

func TestWithdrawPermissionsWorks(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)

	roles.AssignPermissions(Role1, permission0, permission2, permission4, permission6)
	roles.AssignPermissions(Role1, permission1, permission3, permission5, permission7)

	assert.True(roles.HasPermission(Role1, permission0))
	assert.True(roles.HasPermission(Role1, permission2))
	assert.True(roles.HasPermission(Role1, permission4))
	assert.True(roles.HasPermission(Role1, permission6))

	assert.True(roles.HasPermission(Role1, permission1))
	assert.True(roles.HasPermission(Role1, permission3))
	assert.True(roles.HasPermission(Role1, permission5))
	assert.True(roles.HasPermission(Role1, permission7))

	roles.WithdrawPermissions(Role1, permission1, permission3, permission5, permission7)

	assert.True(roles.HasPermission(Role1, permission0))
	assert.True(roles.HasPermission(Role1, permission2))
	assert.True(roles.HasPermission(Role1, permission4))
	assert.True(roles.HasPermission(Role1, permission6))

	assert.False(roles.HasPermission(Role1, permission1))
	assert.False(roles.HasPermission(Role1, permission3))
	assert.False(roles.HasPermission(Role1, permission5))
	assert.False(roles.HasPermission(Role1, permission7))
}

func TestHasPermissionWorksWithoutRace(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)
	roles.AssignPermissions(Role1, permission0, permission2, permission4, permission6)
	roles.AssignPermissions(Role2, permission0, permission1, permission3, permission5, permission7)

	go func() {
		assert.True(roles.HasPermission(Role1, permission0))
		assert.True(roles.HasPermission(Role1, permission2))
		assert.True(roles.HasPermission(Role1, permission4))
		assert.True(roles.HasPermission(Role1, permission6))
	}()

	go func() {
		assert.True(roles.HasPermission(Role2, permission0))
		assert.True(roles.HasPermission(Role2, permission1))
		assert.True(roles.HasPermission(Role2, permission3))
		assert.True(roles.HasPermission(Role2, permission5))
		assert.True(roles.HasPermission(Role2, permission7))
	}()

	go func() {
		assert.False(roles.HasPermission(Role2, permission2))
		assert.False(roles.HasPermission(Role2, permission4))
		assert.False(roles.HasPermission(Role2, permission6))
	}()

	go func() {
		assert.False(roles.HasPermission(Role1, permission1))
		assert.False(roles.HasPermission(Role1, permission3))
		assert.False(roles.HasPermission(Role1, permission5))
		assert.False(roles.HasPermission(Role1, permission7))
	}()
}
