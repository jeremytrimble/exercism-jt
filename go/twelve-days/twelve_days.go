package twelve

import "fmt"

const testVersion = 1

type infoType struct {
	day_ordinal string
	gift        string
}

var dayInfo = []infoType{
	{ "first",     "a Partridge in a Pear Tree" },
	{ "second",    "two Turtle Doves"           },
	{ "third",     "three French Hens"          },
	{ "fourth",    "four Calling Birds"         },
	{ "fifth",     "five Gold Rings"            },
	{ "sixth",     "six Geese-a-Laying"         },
	{ "seventh",   "seven Swans-a-Swimming"     },
	{ "eighth",    "eight Maids-a-Milking"      },
	{ "ninth",     "nine Ladies Dancing"        },
	{ "tenth",     "ten Lords-a-Leaping"        },
	{ "eleventh",  "eleven Pipers Piping"       },
	{ "twelfth",   "twelve Drummers Drumming"   },
}

func gifts_for_verse(verse_num int) string {

	rv := ""

	for idx := verse_num - 1; idx >= 0; idx-- {
		rv += dayInfo[idx].gift

		if idx != 0 {
			rv += ", "
		} else {
			rv += "."
		}

		if idx == 1 && verse_num > 1 {
			rv += "and "
		}
	}

	return rv
}

func Verse(verse_num int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s",
		dayInfo[verse_num-1].day_ordinal, gifts_for_verse(verse_num))
}

func Song() string {
	rv := ""

	for i := 1; i <= 12; i++ {
		rv += Verse(i) + "\n"
	}

	return rv
}
