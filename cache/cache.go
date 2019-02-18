package cache

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)


type DataCache struct {
	CacheData *cache.Cache
	lock    sync.RWMutex
}

//NewCaheData cache all realtime data
func NewCaheData()*DataCache{
	return &DataCache{CacheData:cache.New(2*time.Minute, 3*time.Minute)}
}


func (ch *DataCache) Set(ID string, value interface{}, exp time.Duration) {
	ch.CacheData.Set(ID, value, exp)
}

func (ch *DataCache) Get(ID string) (interface{}, bool) {
	return ch.CacheData.Get(ID)
}
