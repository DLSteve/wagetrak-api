package currency

func CalculateExchangeFromBase(sourceRate, targetRate float32) float32 {
	if sourceRate == 0.0 || targetRate == 0.0 {
		return 0.0
	}

	inverseRate1 := 1 / sourceRate
	inverseRate2 := 1 / targetRate

	return inverseRate1 / inverseRate2
}
