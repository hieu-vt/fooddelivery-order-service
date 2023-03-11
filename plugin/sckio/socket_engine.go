package sckio

import (
	"log"
	"sync"
)

type socketEngine struct {
	storage map[int]Conn
	locker  *sync.RWMutex
	prefix  string
}

func NewSocketEngine(prefix string) *socketEngine {
	return &socketEngine{
		storage: make(map[int]Conn),
		locker:  new(sync.RWMutex),
		prefix:  prefix,
	}
}

func (engine *socketEngine) GetPrefix() string {
	return engine.prefix
}

func (engine *socketEngine) Get() interface{} {
	return engine
}

func (engine *socketEngine) Name() string {
	return engine.prefix
}

func (engine *socketEngine) InitFlags() {
}

func (engine *socketEngine) Configure() error {
	return nil
}

func (engine *socketEngine) Run() error {
	return nil
}

func (engine *socketEngine) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()
	return c
}

func (engine *socketEngine) SaveAppSocket(userId int, appSck Conn) {
	engine.locker.Lock()

	//appSck.Join("order-{ordID}")

	if engine.storage[userId] != nil {
		engine.storage[userId].Close()
	}

	engine.storage[userId] = appSck

	engine.locker.Unlock()
}

func (engine *socketEngine) GetAppSocket(userId int) Conn {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	log.Println(engine.storage)

	return engine.storage[userId]
}

func (engine *socketEngine) RemoveAppSocket(userId int) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	//if v, ok := engine.storage[userId]; ok {
	//	for i := range v {
	//		if v[i] == appSck {
	//			engine.storage[userId] = append(v[:i], v[i+1:]...)
	//			break
	//		}
	//	}
	//}
}
