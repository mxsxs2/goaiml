package goaiml

import "encoding/xml"

//The name of the bot
const BOT_NAME string = "Eliza"

type AIML struct {
	Memory map[string]string
	Bot    map[string]string
	Root   AIMLRoot
}

type AIMLRoot struct {
	XMLName    xml.Name       `xml:"aiml"`
	Categories []AIMLCategory `xml:"category"`
}

//Cateegory for each pattern and template
type AIMLCategory struct {
	XMLName  xml.Name     `xml:"category"`
	Pattern  AIMLPattern  `xml:"pattern"`
	Template AIMLTemplate `xml:"template"`
}

//Template structure
type AIMLTemplate struct {
	XMLName xml.Name `xml:"template"`
	Content string   `xml:",innerxml"`
	Looped  bool
}

//AIML pattern structure
type AIMLPattern struct {
	XMLName xml.Name `xml:"pattern"`
	Content string   `xml:",innerxml"`
}

//Used to map in <star/>
var Reflections = map[string]string{
	"am":        "are",
	"your":      "my",
	"me":        "you",
	"myself":    "yourself",
	"yourself":  "myself",
	"i":         "you",
	"you":       "I",
	"my":        "your",
	"i am":      "you are",
	"are":       "am",
	"i would":   "you would",
	"you would": "i'd",
	"i have":    "you have",
	"you have":  "i'd",
	"i will":    "you will",
	"you will":  "i'll",
}

//Used to preprocess the sentences for better regex match
var PreProcessWords = map[string]string{
	"dont":       "don't",
	"cant":       "can't",
	"wont":       "won't",
	"recollect":  "remember",
	"recall":     "remember",
	"dreamt":     "dreamed",
	"dreams":     "dream",
	"maybe":      "perhaps",
	"certainly":  "yes",
	"machine":    "computer",
	"machines":   "computer",
	"computers":  "computer",
	"were":       "was",
	"you're":     "you are",
	"i'm":        "i am",
	"same":       "alike",
	"identical":  "alike",
	"equivalent": "alike",
	"he'd":       "he would",
	"he'll":      "he will",
	"she'd":      "she would",
	"she'll":     "she will",
	"it'd":       "it would",
	"it'll":      "it will",
	"we'd":       "we would",
	"we'll":      "we will",
	"they'd":     "they would",
	"they'll":    "they will",
	"i've":       "I have",
	"ima":        "I am going to",
	"wanna":      "want to",
	"gonna":      "going to",
	"he's":       "he is",
	"she's":      "she is",
	"it's":       "it is",
	"that's":     "that is",
}

//Function used to start a new AIML parser
func NewAIML() *AIML {
	//Create a new AIML
	ret := &AIML{
		Memory: make(map[string]string),
		Bot:    make(map[string]string),
	}
	//Assign the bots name
	ret.Bot["name"] = BOT_NAME
	//Return the AIML parser
	return ret
}
