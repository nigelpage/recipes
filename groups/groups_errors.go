package groups

import (
	"fmt"
)

type GroupsError struct {
	val string
	msg string
}

func (ge *GroupsError) Error() {
	return fmt.Sprintf("%s - %s", ge.msg, ge.val)
}

func NewGroupNotFoundError(name string) string {
	return &GroupsError {
		val: name
		msg: "Group not found"
	}
}