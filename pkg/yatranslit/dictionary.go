package yatranslit

type TranslitDict map[string]map[string]string

const (
	anyOther = "*"
	special  = ":"
)

var Dictionary = TranslitDict{
	"а": {anyOther: "a"},
	"б": {anyOther: "b"},
	"в": {anyOther: "v"},
	"г": {anyOther: "g"},
	"д": {anyOther: "d"},
	"ж": {anyOther: "zh"},
	"з": {anyOther: "z"},
	"и": {anyOther: "i"},
	"й": {anyOther: "j"},
	"е": {anyOther: "e"},
	"ё": {anyOther: "yo"},
	"о": {anyOther: "o"},
	"п": {anyOther: "p"},
	"р": {anyOther: "r"},
	"к": {anyOther: "k"},
	"л": {anyOther: "l"},
	"м": {anyOther: "m"},
	"н": {anyOther: "n"},
	"т": {anyOther: "t"},
	"ч": {anyOther: "ch"},
	"с": {
		anyOther: "s",
		"см":     "s",
		"дс":     "s",
		"бс":     "s",
		"вс":     "s",
		"су":     "c",
		"св":     "s",
	},
	"ц": {anyOther: "c"},
	"у": {anyOther: "u"},
	"ф": {anyOther: "f"},
	"х": {
		anyOther: "h",
		"сх":     "kh",
		"ех":     "kh",
		"эх":     "kh",
	},
	"ш": {anyOther: "sh"},
	"щ": {anyOther: "shch"},
	"э": {anyOther: "eh"},
	"ю": {anyOther: "yu"},
	"я": {anyOther: "ya"},
	"ы": {anyOther: "y"},
	"ь": {anyOther: ""},
	"ъ": {anyOther: ""},
	// Some special characters
	special: {anyOther: "-"},
}
