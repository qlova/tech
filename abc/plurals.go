package abc

import (
	"strings"

	"qlova.tech/min"
)

//Plural returns the plural form of the given english word.
func Plural(word string) string {
	if len(word) == 0 {
		return ""
	}

	lower := strings.ToLower(word)

	if plural, ok := plurals[lower]; ok {
		return plural
	}
	if _, ok := singles[lower]; ok {
		return word
	}

	for i := min.Int(len(word), len(suffixToPlural)-1); i > 0; i-- {
		suffix := word[len(word)-i:]

		if rule, ok := suffixToPlural[i][suffix]; ok {
			return word[:len(word)-i] + rule
		}
	}

	if strings.HasSuffix(word, "s") {
		return word
	}

	return word + "s"
}

var plurals = map[string]string{
	"aircraft":    "aircraft",
	"alias":       "aliases",
	"alumna":      "alumnae",
	"alumnus":     "alumni",
	"analysis":    "analyses",
	"antenna":     "antennas",
	"antithesis":  "antitheses",
	"apex":        "apexes",
	"appendix":    "appendices",
	"axis":        "axes",
	"bacillus":    "bacilli",
	"bacterium":   "bacteria",
	"basis":       "bases",
	"beau":        "beaus",
	"bison":       "bison",
	"bureau":      "bureaus",
	"bus":         "buses",
	"campus":      "campuses",
	"caucus":      "caucuses",
	"child":       "children",
	"château":     "châteaux",
	"circus":      "circuses",
	"codex":       "codices",
	"concerto":    "concertos",
	"corpus":      "corpora",
	"crisis":      "crises",
	"curriculum":  "curriculums",
	"datum":       "data",
	"deer":        "deer",
	"diagnosis":   "diagnoses",
	"die":         "dice",
	"dwarf":       "dwarves",
	"ellipsis":    "ellipses",
	"equipment":   "equipment",
	"erratum":     "errata",
	"faux pas":    "faux pas",
	"fez":         "fezzes",
	"fish":        "fish",
	"focus":       "foci",
	"foo":         "foos",
	"foot":        "feet",
	"formula":     "formulas",
	"fungus":      "fungi",
	"genus":       "genera",
	"goose":       "geese",
	"graffito":    "graffiti",
	"grouse":      "grouse",
	"half":        "halves",
	"halo":        "halos",
	"hoof":        "hooves",
	"human":       "humans",
	"hypothesis":  "hypotheses",
	"index":       "indices",
	"information": "information",
	"jeans":       "jeans",
	"larva":       "larvae",
	"libretto":    "librettos",
	"loaf":        "loaves",
	"locus":       "loci",
	"louse":       "lice",
	"matrix":      "matrices",
	"minutia":     "minutiae",
	"money":       "money",
	"moose":       "moose",
	"mouse":       "mice",
	"nebula":      "nebulae",
	"news":        "news",
	"nucleus":     "nuclei",
	"oasis":       "oases",
	"octopus":     "octopi",
	"offspring":   "offspring",
	"opus":        "opera",
	"ovum":        "ova",
	"ox":          "oxen",
	"parenthesis": "parentheses",
	"phenomenon":  "phenomena",
	"photo":       "photos",
	"phylum":      "phyla",
	"piano":       "pianos",
	"plus":        "pluses",
	"police":      "police",
	"prognosis":   "prognoses",
	"prometheus":  "prometheuses",
	"quiz":        "quizzes",
	"quota":       "quotas",
	"radius":      "radiuses",
	"referendum":  "referendums",
	"ress":        "resses",
	"rice":        "rice",
	"salmon":      "salmon",
	"sex":         "sexes",
	"series":      "series",
	"sheep":       "sheep",
	"shoe":        "shoes",
	"shrimp":      "shrimp",
	"species":     "species",
	"stimulus":    "stimuli",
	"stratum":     "strata",
	"swine":       "swine",
	"syllabus":    "syllabi",
	"symposium":   "symposiums",
	"synapse":     "synapses",
	"synopsis":    "synopses",
	"tableau":     "tableaus",
	"testis":      "testes",
	"thesis":      "theses",
	"thief":       "thieves",
	"tooth":       "teeth",
	"trout":       "trout",
	"tuna":        "tuna",
	"vedalia":     "vedalias",
	"vertebra":    "vertebrae",
	"vertix":      "vertices",
	"vita":        "vitae",
	"vortex":      "vortices",
	"wharf":       "wharves",
	"wife":        "wives",
	"woman":       "women",
	"wolf":        "wolves",
	"you":         "you",
}

var suffixToPlural = [8]map[string]string{
	1: {
		"o": "oes",
		"x": "xes",
	},
	2: {
		"by": "bies",
		"ch": "ches",
		"cy": "cies",
		"dy": "dies",
		"ex": "ices",
		"fy": "fies",
		"gy": "gies",
		"hy": "hies",
		"io": "ios",
		"jy": "jies",
		"ky": "kies",
		"lf": "lves",
		"ly": "lies",
		"my": "mies",
		"ny": "nies",
		"py": "pies",
		"qy": "qies",
		"rf": "rves",
		"ry": "ries",
		"sh": "shes",
		"ss": "sses",
		"sy": "sies",
		"ty": "ties",
		"tz": "tzes",
		"va": "vae",
		"vy": "vies",
		"wy": "wies",
		"xy": "xies",
		"zy": "zies",
		"zz": "zzes",
	},
	3: {
		"afe": "aves",
		"bfe": "bves",
		"box": "boxes",
		"cfe": "cves",
		"dfe": "dves",
		"dge": "dges",
		"efe": "eves",
		"gfe": "gves",
		"hfe": "hves",
		"ife": "ives",
		"itz": "itzes",
		"ium": "ia",
		"ize": "izes",
		"jfe": "jves",
		"kfe": "kves",
		"man": "men",
		"mfe": "mves",
		"nfe": "nves",
		"nna": "nnas",
		"oaf": "oaves",
		"oci": "ocus",
		"ode": "odes",
		"ofe": "oves",
		"oot": "eet",
		"pfe": "pves",
		"pse": "psis",
		"qfe": "qves",
		"quy": "quies",
		"rfe": "rves",
		"sfe": "sves",
		"tfe": "tves",
		"tum": "ta",
		"tus": "tuses",
		"ufe": "uves",
		"ula": "ulas",
		"uli": "ulus",
		"use": "uses",
		"uss": "usses",
		"vfe": "vves",
		"wfe": "wves",
		"xfe": "xves",
		"yfe": "yves",
		"you": "you",
		"zfe": "zves",
	},
	4: {
		"atum": "ata",
		"atus": "atuses",
		"base": "bases",
		"cess": "cesses",
		"dium": "diums",
		"eses": "esis",
		"half": "halves",
		"hive": "hives",
		"iano": "ianos",
		"irus": "iri",
		"isis": "ises",
		"leus": "li",
		"mnus": "mni",
		"move": "moves",
		"news": "news",
		"odex": "odice",
		"oose": "eese",
		"ouse": "ouses",
		"ovum": "ova",
		"rion": "ria",
		"shoe": "shoes",
		"stis": "stes",
		"tive": "tives",
		"wife": "wives",
	},
	5: {
		"actus": "acti",
		"adium": "adia",
		"basis": "basis",
		"child": "children",
		"chive": "chives",
		"focus": "foci",
		"hello": "hellos",
		"jeans": "jeans",
		"louse": "lice",
		"media": "media",
		"mouse": "mice",
		"movie": "movies",
		"oasis": "oasis",
	},
	6: {
		"campus": "campuses",
		"genera": "genus",
		"person": "people",
		"phylum": "phyla",
		"randum": "randa",
	},
	7: {
		"iterion": "iteria",
	},
}
