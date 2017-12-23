package ext

var Pi float64 = float64(3.141592)

// yo:= 3.141592 //関数外で暗黙的な

func GetPI() float64 {
	yo := float64(3.14159265358979)
	return yo
}
