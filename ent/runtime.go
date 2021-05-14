// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/lemon-mint/notification-bot/ent/schema"
	"github.com/lemon-mint/notification-bot/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIsAdmin is the schema descriptor for IsAdmin field.
	userDescIsAdmin := userFields[1].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the IsAdmin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}