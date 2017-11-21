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
	"am":       "are",
	"your":     "my",
	"me":       "you",
	"myself":   "yourself",
	"yourself": "myself",
	"i":        "you",
	"you":      "I",
	"my":       "your",
	"i'm":      "you are",
	"are":      "am",
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
