package boardInfo

var buildCosts = map[string]int{
	"A30": 40,
	"B27": 30,
	"B29": 30,
	"C22": 10,
	"C24": 10,
	"C26": 40,
	"C28": 20,
	"C30": 20,
	"D19": 10,
	"D21": 10,
	"D23": 10,
	"D25": 40,
	"D27": 40,
	"D29": 10,
	"E4":  20,
	"E8":  10,
	"E10": 10,
	"E12": 10,
	"E16": 10,
	"E18": 10,
	"E20": 40,
	"E22": 40,
	"E24": 40,
	"E26": 40,
	"E28": 10,
	"E30": 20,
	"F3":  10,
	"F5":  10,
	"F7":  10,
	"F9":  10,
	"F11": 10,
	"F13": 10,
	"F15": 10,
	"F17": 80,
	"F19": 80,
	"F21": 60,
	"F23": 40,
	"F25": 60,
	"F27": 10,
	"G2":  10,
	"G4":  10,
	"G6":  10,
	"G8":  10,
	"G10": 10,
	"G12": 10,
	"G14": 20,
	"G16": 60,
	"G18": 80,
	"G20": 20,
	"G22": 20,
	"G24": 20,
	"H1":  10,
	"H3":  10,
	"H5":  10,
	"H7":  10,
	"H9":  10,
	"H11": 10,
	"H13": 20,
	"H15": 60,
	"H17": 60,
	"H19": 20,
	"H21": 10,
	"H23": 10,
	"I0":  40,
	"I2":  10,
	"I4":  10,
	"I6":  10,
	"I8":  10,
	"I10": 20,
	"I12": 40,
	"I14": 80,
	"I16": 100,
	"I18": 80,
	"I20": 10,
	"I22": 20,
	"I24": 10,
	"J1":  10,
	"J3":  10,
	"J5":  10,
	"J7":  20,
	"J9":  10,
	"J11": 40,
	"J13": 40,
	"J15": 100,
	"J17": 80,
	"J19": 20,
	"J21": 20,
	"K2":  20,
	"K4":  10,
	"K6":  10,
	"K8":  10,
	"K10": 20,
	"K12": 100,
	"K14": 80,
	"K16": 20,
	"K18": 10,
	"K20": 10,
	"K22": 20,
}

// BuildCost returns the cost for a company to lay down track on the specified hex tile.
// If the provided hex coordinate is not a valid part of the map it will return 0.
func BuildCost(hexCoord string) int {
	return buildCosts[hexCoord]
}

// TrainCost calculates the cost of the number-th train that can be bought in the game. The
// pattern for the train cost is that the first train of a given tech level is the most expensive,
// with each subsequent one decreasing in price from the previous by 5*tech level. The cheapest
// (last) train of a given tech level is also the same price as the most expensive (first) train
// of the previous tech level.
func TrainCost(number int) int {
	// First find out what tech level we are in, then figure out the first of the first train
	// of that tech level. This can be calculated by using the formula for triangle numbers
	// adjusted for the number of trains in each level and the base change between each train,
	// added to the cost of the cheapest train (last one of first tech level)
	techLvl := TechLevel(number)
	firstCost := 80 + 20*techLvl*(techLvl+1)/2
	return firstCost - 5*techLvl*((number-1)%5)
}

func AllTrainCosts() (result [6][5]int) {
	for lvl := range result {
		for num := range result[lvl] {
			result[lvl][num] = TrainCost(5*lvl + num + 1)
		}
	}
	return
}
