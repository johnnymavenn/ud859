package ud859

import (
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

const keyNoFilters = "CACHE_NO_FILTERS"

var deleteCacheNoFilters = delay.Func("delete_no_filters",
	func(c context.Context) {
		err := memcache.Delete(c, keyNoFilters)
		if err != nil {
			log.Errorf(c, "unable to delete cache: %v", err)
		}
	})

var setCacheNoFilters = delay.Func("set_no_filters",
	func(c context.Context, conferences *Conferences) {
		item := &memcache.Item{
			Key:        keyNoFilters,
			Object:     conferences,
			Expiration: 10 * time.Minute,
		}
		err := memcache.Gob.Set(c, item)
		if err != nil {
			log.Errorf(c, "unable to set cache: %v", err)
		}
	})

func getCacheNoFilters(c context.Context) *Conferences {
	conferences := new(Conferences)
	_, err := memcache.Gob.Get(c, keyNoFilters, conferences)
	if err == memcache.ErrCacheMiss {
		return nil
	} else if err != nil {
		log.Errorf(c, "unable to get cache: %v", err)
		return nil
	}
	return conferences
}
