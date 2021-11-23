//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync"
	"time"
)

const freemiumUserMaxTime = 10

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	mutex     sync.Mutex
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.exceededAccumulatedTime() {
		return false
	}
	t := time.Now()

	// this should be cancalled automatically with context cancel or channel bool if it runs more than 10s, but for this simple example this will do
	process()
	return u.afterProcessChecker(t)
}

func main() {
	RunMockServer()
}

func (u *User) exceededAccumulatedTime() bool {
	if u.IsPremium {
		return false
	}

	return u.hasUserExceededTheirAllocatedTime()
}

func (u *User) addTimeUsed(time time.Duration) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.TimeUsed += int64(time.Seconds())
}

// afterProcessChecker if true is returned, the process ran and was not cancalled
func (u *User) afterProcessChecker(t time.Time) bool {
	if u.IsPremium {
		return true
	}

	u.addTimeUsed(time.Since(t))

	return !u.hasUserExceededTheirAllocatedTime()
}

func (u *User) hasUserExceededTheirAllocatedTime() bool {
	return u.TimeUsed > freemiumUserMaxTime
}
