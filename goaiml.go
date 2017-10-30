package goaiml

import "encoding/xml"

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

type AIMLCategory struct {
	XMLName  xml.Name     `xml:"category"`
	Pattern  AIMLPattern  `xml:"pattern"`
	Template AIMLTemplate `xml:"template"`
}

type AIMLTemplate struct {
	XMLName xml.Name `xml:"template"`
	Content string   `xml:",innerxml"`
	Looped  bool
}

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
}

func NewAIML() *AIML {
	ret := &AIML{
		Memory: make(map[string]string),
		Bot:    make(map[string]string),
	}

	ret.Bot["name"] = BOT_NAME

	return ret
}
