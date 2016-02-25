package mocks

import (
	"github.com/v2ray/v2ray-core/common/protocol"
)

type MockUserSet struct {
	Users      []*protocol.User
	UserHashes map[string]int
	Timestamps map[string]protocol.Timestamp
}

func (us *MockUserSet) Add(user *protocol.User) error {
	us.Users = append(us.Users, user)
	return nil
}

func (us *MockUserSet) Get(userhash []byte) (*protocol.User, protocol.Timestamp, bool) {
	idx, found := us.UserHashes[string(userhash)]
	if found {
		return us.Users[idx], us.Timestamps[string(userhash)], true
	}
	return nil, 0, false
}
