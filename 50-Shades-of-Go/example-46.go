package main

import "sync"

type myLocker sync.Locker

func main() {
	var lock myLocker = new(sync.Mutex)
	lock.Lock()   //ok
	lock.Unlock() //ok
}
