package deadlock

import (
	"time"
	"github.com/phk13/multithreading/deadlocks_train/common"
)

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.Mutex.LockedBy = train.Id
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
			time.Sleep(30 * time.Millisecond)
		}
	}
}
