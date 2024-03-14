package core

import "guowei.com/mmo_game_actor/core/module"

type mod struct {
	mi module.Module
}

type App struct {
	mods []*mod // App 这个臭宝有哪些模块
}

func DefaultApp() *App {
	return &App{}
}

func (a *App) Start(mods ...module.Module) {
	if len(mods) == 0 {
		return
	}

	// 模块注册
	for _, item := range mods {
		m := new(mod)
		m.mi = item
		a.mods = append(a.mods, m)
	}

	// 模块初始化
	for _, item := range a.mods {
		mi := item.mi
		if err := mi.OnInit(); err != nil {
		}
	}

	// 模块启动
	for _, item := range a.mods {
		go run(item)
	}
}

// 原神, 启动
func run(m *mod) {
	m.mi.Run()
}
