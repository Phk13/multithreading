package deadlock

import (
	"sort"
	"time"

	"github.com/phk13/multithreading/deadlocks_train/common"
)

func lockInsersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*common.Crossing) {
	var intersectionsToLock []*common.Intersection
	for _, crossing := range crossings {
		if crossing.Intersection.LockedBy != id && reserveEnd >= crossing.Position && reserveStart <= crossing.Position {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].Id < intersectionsToLock[j].Id
	})

	for _, it := range intersectionsToLock {
		it.Mutex.Lock()
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
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
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
			time.Sleep(30 * time.Millisecond)
		}
	}
}
