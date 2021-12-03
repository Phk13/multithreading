package deadlock

import (
	"sync"
	"time"

	"github.com/phk13/multithreading/deadlocks_train/common"
)

var (
	controller = sync.Mutex{}
	cond       = sync.NewCond(&controller)
)

func allFree(intersectionsToLock []*common.Intersection) bool {
	for _, it := range intersectionsToLock {
		if it.LockedBy >= 0 {
			return false
		}
	}
	return true
}

func lockInsersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*common.Crossing) {
	var intersectionsToLock []*common.Intersection
	for _, crossing := range crossings {
		if crossing.Intersection.LockedBy != id && reserveEnd >= crossing.Position && reserveStart <= crossing.Position {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	controller.Lock()
	for !allFree(intersectionsToLock) {
		cond.Wait()
	}

	for _, it := range intersectionsToLock {
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
	controller.Unlock()
}

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockInsersectionsInDistance(train.Id, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				controller.Lock()
				crossing.Intersection.LockedBy = -1
				cond.Broadcast()
				controller.Unlock()
			}
			time.Sleep(30 * time.Millisecond)
		}
	}
}
