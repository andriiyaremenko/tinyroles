# TinyRoles

Simple Role-Manager library for Golang

This library allows you to create roles, assign, withdraw and check permissions.

## Key Features

* Assign permissions to a role
* Withdraw permissions from a role
* Check if role has a permission

## Under the hood
This library uses [Bit Flags](https://en.wikipedia.org/wiki/Bit_field) under the hood which makes it simple and fast.

## Installing
````bash
go get github.com/andriiyaremenko/tinyroles
````

## Create Permissions
````go
package main

import (
	"fmt"
	"github.com/andriiyaremenko/tinyroles"
)

var (
	permission0 = tinyroles.NewPermission(1)
	permission1 = tinyroles.NewPermission(2)
	permission2 = tinyroles.NewPermission(3)
	permission3 = tinyroles.NewPermission(4)
)

func main() {
	fmt.Println(permission0)
	// Output: 2
	fmt.Println(permission1)
	// Output: 4
	fmt.Println(permission2)
	// Output: 8
	fmt.Println(permission3)
	// Output: 16
}
````

## Assign a Permissions to a Role
````go
package main

import (
	"fmt"
	"github.com/andriiyaremenko/tinyroles"
)

const (
	Role1 tinyroles.Role = "MyUser"
)

var (
	permission0 = tinyroles.NewPermission(1)
	permission1 = tinyroles.NewPermission(2)
	permission2 = tinyroles.NewPermission(3)
	permission3 = tinyroles.NewPermission(4)
)

func main() {
	roles := new(tinyroles.Roles)
	roles.AssignPermissions(Role1, permission0, permission1, permission2)

	fmt.Println(roles.HasPermission(Role1, permission0))
	// Output: true
	fmt.Println(roles.HasPermission(Role1, permission1))
	// Output: true
	fmt.Println(roles.HasPermission(Role1, permission2))
	// Output: true
	fmt.Println(roles.HasPermission(Role1, permission3))
	// Output: false
	fmt.Println(roles.GetRoleValue(Role1))
	// Output: 14
}
````

## Withdraw Permissions from a Role
````go
package main

import (
	"fmt"
	"github.com/andriiyaremenko/tinyroles"
)

const (
	tinyroles.Role = "MyUser"
)

var (
	permission0 = tinyroles.NewPermission(1)
	permission1 = tinyroles.NewPermission(2)
	permission2 = tinyroles.NewPermission(3)
	permission3 = tinyroles.NewPermission(4)
)

func main() {
	roles := new(tinyroles.Roles)
	roles.AssignPermissions(Role1, permission0, permission1, permission2)
	roles.WithdrawPermissions(Role1, permission0)

	fmt.Println(roles.HasPermission(Role1, permission0))
	// Output: false
	fmt.Println(roles.HasPermission(Role1, permission1))
	// Output: true
	fmt.Println(roles.HasPermission(Role1, permission2))
	// Output: true
	fmt.Println(roles.HasPermission(Role1, permission3))
	// Output: false
	fmt.Println(roles.GetRoleValue(Role1))
	// Output: 12
}
````

