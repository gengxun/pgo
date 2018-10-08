package Redis

import (
    "time"

    "github.com/pinguo/pgo"
    "github.com/pinguo/pgo/Util"
)

// Adapter of Redis Client, add context support.
// usage: redis := this.GetObject("@pgo/Client/Redis/Adapter").(*Adapter)
type Adapter struct {
    pgo.Object
    client       *Client
    panicRecover bool
}

func (a *Adapter) Construct(componentId ...string) {
    id := defaultComponentId
    if len(componentId) > 0 {
        id = componentId[0]
    }

    a.client = pgo.App.Get(id).(*Client)
    a.panicRecover = true
}

func (a *Adapter) SetPanicRecover(v bool) {
    a.panicRecover = v
}

func (a *Adapter) GetClient() *Client {
    return a.client
}

func (a *Adapter) handlePanic() {
    if a.panicRecover {
        if v := recover(); v != nil {
            a.GetContext().Error(Util.ToString(v))
        }
    }
}

func (a *Adapter) Get(key string) *pgo.Value {
    profile := "Redis.Get"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    res, hit := a.client.Get(key), 0
    if res != nil {
        hit = 1
    }

    a.GetContext().Counting(profile, hit, 1)
    return res
}

func (a *Adapter) MGet(keys []string) map[string]*pgo.Value {
    profile := "Redis.MGet"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    res, hit := a.client.MGet(keys), 0
    for _, v := range res {
        if v != nil {
            hit += 1
        }
    }

    a.GetContext().Counting(profile, hit, len(keys))
    return res
}

func (a *Adapter) Set(key string, value interface{}, expire ...time.Duration) bool {
    profile := "Redis.Set"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    return a.client.Set(key, value, expire...)
}

func (a *Adapter) MSet(items map[string]interface{}, expire ...time.Duration) bool {
    profile := "Redis.MSet"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    return a.client.MSet(items, expire...)
}

func (a *Adapter) Add(key string, value interface{}, expire ...time.Duration) bool {
    profile := "Redis.Add"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    return a.client.Add(key, value, expire...)
}

func (a *Adapter) MAdd(items map[string]interface{}, expire ...time.Duration) bool {
    profile := "Redis.MAdd"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStart(profile)
    defer a.handlePanic()

    return a.client.MAdd(items, expire...)
}

func (a *Adapter) Del(key string) bool {
    profile := "Redis.Del"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    return a.client.Del(key)
}

func (a *Adapter) MDel(keys []string) bool {
    profile := "Redis.MDel"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    return a.client.MDel(keys)
}

func (a *Adapter) Exists(key string) bool {
    profile := "Redis.Exists"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    return a.client.Exists(key)
}

func (a *Adapter) Incr(key string, delta int) int {
    profile := "Redis.Incr"
    a.GetContext().ProfileStart(profile)
    defer a.GetContext().ProfileStop(profile)
    defer a.handlePanic()

    return a.client.Incr(key, delta)
}
