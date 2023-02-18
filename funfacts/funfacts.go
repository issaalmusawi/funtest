package funfacts

type FunFacts struct {
	Sun, Luna, Terra []Facts
}
type Facts struct {
	Interesting     string
	Value           float64
	temperatureType string
}

func GetFunFacts(about string) []Facts {
	fFacts := FunFacts{
		Sun: []Facts{
			{Interesting: "temperatur i solens kjerne", Value: 15000000, temperatureType: "C"},
			{Interesting: "temperatur på ytre lag av solen", Value: 5778, temperatureType: "K"},
		},

		Terra: []Facts{
			{Interesting: "høyeste temperatur målt på jordens overflate", Value: 56.7, temperatureType: "C"},
			{Interesting: "laveste temperatur målt på jordens overflate", Value: -69.7, temperatureType: "C"},
			{Interesting: "temperatur i jordens indre kjerne", Value: 9118.85, temperatureType: "C"},
		},

		Luna: []Facts{
			{Interesting: "temperatur på månens overflate om natten", Value: -183, temperatureType: "C"},
			{Interesting: "temperatur på månens overflate om dagen", Value: 106, temperatureType: "C"},
		},
	}

	if about == "Sun" {
		return fFacts.Sun
	} else if about == "Terra" {
		return fFacts.Terra
	} else if about == "Luna" {
		return fFacts.Luna
	} else {
		return []Facts{}
	}
}
