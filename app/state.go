package app

import "git-knowledge/logger"

// State 应用程序状态，状态不同，方法的行为也会不同
type State interface {
	Start(a *App)
	Stop(a *App)
	Restart(a *App)
	Load(a *App)
}

// StateNew 应用程序刚刚调用 NewApp 方法
type StateNew struct{}

func (s *StateNew) Start(a *App) {
	err := a.echo.Start(":8080")
	if err != nil {
		logger.Fatal("启动服务出现错误 %s", err)
	}
}

func (s *StateNew) Stop(a *App) {
	panic("implement me")
}

func (s *StateNew) Restart(a *App) {
	panic("implement me")
}

func (s *StateNew) Load(a *App) {
	panic("implement me")
}
