package visualService

// Start 启动模块服务
// 传入：启动参数
// 传出：无
func (app *VisualApp) Start() {
	//todo
}

// Stop 中止模块服务
// 传入：无
// 传出：无
func (app *VisualApp) Stop() {
	//todo
}

// GetApp 获取App
// 传入：无
// 传出：该模块App的指针
func (app *VisualApp) GetApp() *interface{} {
	var value interface{}
	value = app
	return &value
}

// IsAlive 是否在服务
// 传入：无
// 传出：是否在服务
func (app *VisualApp) IsAlive() bool {
	//todo
	return false
}
