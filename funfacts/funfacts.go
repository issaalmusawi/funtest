package funfacts

//"github.com/issaalmusawi/funtest/konv"

type object struct {
	Sun, Luna, Terra []Fact
}
type Fact struct {
	Text  string
	Value float64
	Type  string
}

func GetFunFacts(about string) []Fact {
	Facts := object{
		Sun: []Fact{
			{Text: "temperatur i solens kjerne", Value: 15000000, Type: "C"},
			{Text: "temperatur på ytre lag av solen", Value: 5778, Type: "K"},
		},

		Terra: []Fact{
			{Text: "høyeste temperatur målt på jordens overflate", Value: 56.7, Type: "C"},
			{Text: "laveste temperatur målt på jordens overflate", Value: -89.4, Type: "C"},
			{Text: "temperatur i jordens indre kjerne", Value: 9392, Type: "K"},
		},

		Luna: []Fact{
			{Text: "temperatur på månens overflate om natten", Value: -183, Type: "C"},
			{Text: "temperatur på månens overflate om dagen", Value: 106, Type: "C"},
		},
	}

	if about == "Sun" {
		return Facts.Sun
	} else if about == "Terra" {
		return Facts.Terra
	} else if about == "Luna" {
		return Facts.Luna
	} else {
		return []Fact{}
	}
}
