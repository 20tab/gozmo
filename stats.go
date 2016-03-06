package gozmo

// there are (currently) three kind of stats: per-frame, per-second and global

func checkPerFrameStats(name string) {
	if Engine.perFrameStats == nil {
		Engine.perFrameStats = make(map[string]float64)
	}
	_, ok := Engine.perFrameStats[name]
	if !ok {
		Engine.perFrameStats[name] = 0
	}
}

func IncPerFrameStats(name string, value float64) {
	checkPerFrameStats(name)
	Engine.perFrameStats[name] += value
}

func DecPerFrameStats(name string, value float64) {
	checkPerFrameStats(name)
	Engine.perFrameStats[name] -= value
}

func SetPerFrameStats(name string, value float64) {
	checkPerFrameStats(name)
	Engine.perFrameStats[name] = value
}

func GetPerFrameStats(name string) float64 {
	checkPerFrameStats(name)
	return Engine.perFrameStats[name]
}

func UpdatePerFrameStats() {
	if Engine.perFrameStats == nil {
		return
	}
	for key, _ := range Engine.perFrameStats {
		Engine.perFrameStats[key] = 0
	}
}
