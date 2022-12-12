package cache

import (
    "Login-Backend/src/config"
    "go.etcd.io/bbolt"
)

var tokenCache *bbolt.DB

func Initialize() error {
    cacheCfg := config.CacheConfig
    var err error
    tokenCache, err = bbolt.Open(cacheCfg.Path, 0600, nil)
    if err != nil {
        return err
    }
    err = tokenCache.Update(func(tx *bbolt.Tx) error {
        _, err = tx.CreateBucketIfNotExists([]byte("token"))
        if err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        return err
    }
    err = tokenCache.Update(func(tx *bbolt.Tx) error {
        _, err = tx.CreateBucketIfNotExists([]byte("email"))
        if err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        return err
    }
    return nil
}
