package loader

// InitLoader 初始化so加载器
// 传入参数：无
// 返回参数：so加载器
func InitLoader() *SoLoader {
	soLoader := new(SoLoader)
	soLoader.Sos = make(map[string]map[int]*SoPackage)
	soLoader.soCounter = make(map[string]int)
	return soLoader
}
