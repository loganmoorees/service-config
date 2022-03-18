package core

import (
	"fmt"
	"service-config/api"
	"sync"
)

type Registry struct {
	Apps map[string]*Application
	Lock sync.RWMutex
}

type Application struct {
	Appid           string
	Instance        map[string]*Instance
	LatestTimestamp int64
	Lock            sync.RWMutex
}

/**
 * Instance加入应用App
 */
func (app *Application) addInstance(request *Instance, latestTime int64) (*Instance, bool) {
	app.Lock.Lock()
	defer app.Lock.Unlock()
	appInstance, ok := app.Instance[request.Hostname]
	if ok {
		request.UpTimestamp = appInstance.UpTimestamp
		if request.DirtyTimestamp < appInstance.DirtyTimestamp {
			request = appInstance
		}
	}
	app.Instance[request.Hostname] = request
	app.upLatestTimestamp(latestTime)
	instanceObj := new(Instance)
	*instanceObj = *request
	return instanceObj, !ok
}

func (app *Application) upLatestTimestamp(latestTime int64) {
	if latestTime <= app.LatestTimestamp {
		latestTime = app.LatestTimestamp
	}
	app.LatestTimestamp = latestTime
}

type Instance struct {
	Env             string   `json:"env"`              // 环境
	AppId           string   `json:"appid"`            // 应用标识
	Hostname        string   `json:"hostname"`         // 实例标识
	Address         []string `json:"address"`          // 地址
	Version         string   `json:"version"`          // 版本
	Status          uint32   `json:"status"`           // 状态
	RegTimestamp    int64    `json:"reg_timestamp"`    // 注册时间
	UpTimestamp     int64    `json:"up_timestamp"`     // 上线时间
	RenewTimestamp  int64    `json:"renew_timestamp"`  // 最近续约时间
	DirtyTimestamp  int64    `json:"dirty_timestamp"`  // 脏时间
	LatestTimestamp int64    `json:"latest_timestamp"` // 最后更新时间
}

func (r *Registry) Register(request *Instance, latestTime int64) {
	r.Lock.RLock()
	app, ok := r.Apps[getKey(request.AppId, request.Env)]
	r.Lock.Unlock()
	if !ok {
		app = api.InitApplication(request.AppId)
	}
	app.addInstance(request, latestTime)

}

func getKey(appId string, env string) string {
	return fmt.Sprintf("%s-%s", appId, env)
}
