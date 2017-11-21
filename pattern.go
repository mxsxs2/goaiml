package goaiml

import (
	"encoding/xml"
	"errors"
	"regexp"
	"strings"
	"unicode"
)

//Function used to remove white spaces from the input string
func stringMinifier(in string) (out string) {
	//White space tag
	white := false
	//Loop the inuput string
	for _, c := range in {
		//If the character is a space
		if unicode.IsSpace(c) {
			//If there was no white space yet
			if !white {
				//Add the string to the output
				out = out + " "
			}
			//Set the white space flag
			white = true
		} else {
			//Add the string to the output
			out = out + string(c)
			//Set the white space flag
			white = false
		}
	}
	//Return
	return
}

//Function used to parse the pattern into a gerex
func (aimlPattern *AIMLPattern) Regexify() *regexp.Regexp {
	//Get the pattern
	rString := aimlPattern.Content
	//Remove white spaces
	rString = stringMinifier(rString)
	//Replace the * to the regex capture *
	rString = strings.Replace(rString, "*", "(.*)", -1)
	//Return the regex
	return regexp.MustCompile("(?i)" + rString)
}

//Function used to process the bot tag form the AIML file
func (aimlPattern *AIMLPattern) ProcessBot(aiml *AIML) error {
	//New structure for the bot
	botStruct := struct {
		XMLName xml.Name `xml:"bot"`
		Name    string   `xml:"name,attr"`
	}{}
	//Parse the pattern content into the botstruct
	err := xml.Unmarshal([]byte(aimlPattern.Content), &botStruct)
	//If there was an error then return it
	if err != nil {
		return err
	}
	//Try to get the bot name
	content, ok := aiml.Bot[botStruct.Name]
	//If the bot name was not found then return an error
	if !ok {
		return errors.New("Key not found in memory")
	}
	//Replace the bots name in the pattern content
	aimlPattern.Content = strings.Replace(aimlPattern.Content, `<bot name="`+botStruct.Name+`"/>`, content, -1)
	//Return
	return nil
}
