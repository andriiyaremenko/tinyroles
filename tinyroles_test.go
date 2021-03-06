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

func TestRoles(t *testing.T) {
	t.Run("HasPermission returns true for assigned permissions and false for unassigned",
		testHasPermissionsWorks)
	t.Run("HasPermission always returns false for role without any permissions",
		testHasPermissionReturnsFalseForRoleWithoutPermissions)
	t.Run("HasPermission is safe to call from different gorutines simultaneously",
		testHasPermissionWorksWithoutRace)
	t.Run("AssignPermissions is idempotent", testAssignPermissionsIsIdempotent)
	t.Run("AssignPermission can be used subsequently", testAssignPermissionsInSubsequentCalls)
	t.Run("WithdrawPermissions withdraws only listed permissions from a role",
		testWithdrawPermissionsWorks)
	t.Run("WithdrawPermissions is safe to call with permissions not assigned to a role previously",
		testWithdrawPermissionsIsSafeToCallWithWrongPermissions)
	t.Run("GetRoleValue returns correct numeric representation of the Role", testGetRoleValueWorks)
}

func testHasPermissionsWorks(t *testing.T) {
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

func testAssignPermissionsIsIdempotent(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)
	roles.AssignPermissions(Role1, permission0, permission2, permission4)
	roles.AssignPermissions(Role2, permission0, permission2, permission4, permission4, permission0)

	assert.Equal(roles.GetRoleValue(Role1), roles.GetRoleValue(Role2))
}

func testGetRoleValueWorks(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)
	roles.AssignPermissions(Role1, permission0, permission2, permission4, permission6)

	expected := permission0.Value() |
		permission2.Value() |
		permission4.Value() |
		permission6.Value()

	assert.Equal(expected, roles.GetRoleValue(Role1))
}

func testHasPermissionReturnsFalseForRoleWithoutPermissions(t *testing.T) {
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

func testAssignPermissionsInSubsequentCalls(t *testing.T) {
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

func testWithdrawPermissionsWorks(t *testing.T) {
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

func testWithdrawPermissionsIsSafeToCallWithWrongPermissions(t *testing.T) {
	assert := assert.New(t)
	roles := new(Roles)

	roles.AssignPermissions(Role1, permission0, permission2, permission4)

	roles.WithdrawPermissions(Role1, permission0, permission1, permission3, permission5, permission7)

	assert.True(roles.HasPermission(Role1, permission2))
	assert.True(roles.HasPermission(Role1, permission4))

	assert.False(roles.HasPermission(Role1, permission0))
	assert.False(roles.HasPermission(Role1, permission1))
	assert.False(roles.HasPermission(Role1, permission3))
	assert.False(roles.HasPermission(Role1, permission5))
	assert.False(roles.HasPermission(Role1, permission6))
	assert.False(roles.HasPermission(Role1, permission7))

	expected := permission2.Value() | permission4.Value()
	assert.Equal(expected, roles.GetRoleValue(Role1))
}

func testHasPermissionWorksWithoutRace(t *testing.T) {
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
