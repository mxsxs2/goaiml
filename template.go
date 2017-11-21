package goaiml

import (
	"encoding/xml"
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//Function used to process the set tag
func (aimlTemplate *AIMLTemplate) ProcessSet(aiml *AIML) error {
	//Structure for the set tag
	setStruct := struct {
		XMLName xml.Name `xml:"set"`
		Name    string   `xml:"name,attr"`
		Content string   `xml:",innerxml"`
	}{}
	//Parse the content into the struct
	err := xml.Unmarshal([]byte(aimlTemplate.Content), &setStruct)
	//Return the error if there was any
	if err != nil {
		return err
	}
	//Remove the set tag from the content
	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `<set name="`+setStruct.Name+`">`, "", -1)
	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `</set>`, "", -1)
	//Set the value in the memory
	aiml.Memory[setStruct.Name] = setStruct.Content
	//Return
	return nil
}

//Function used to process the get tag
func (aimlTemplate *AIMLTemplate) ProcessGet(aiml *AIML) error {
	//Structure for the get tag
	getStruct := struct {
		XMLName xml.Name `xml:"get"`
		Name    string   `xml:"name,attr"`
	}{}
	//Parse the content into the struct
	err := xml.Unmarshal([]byte(aimlTemplate.Content), &getStruct)
	//Return the error if there was any
	if err != nil {
		return err
	}
	//Try to get the value from the memory
	content, ok := aiml.Memory[getStruct.Name]
	//Return the error if there was any
	if !ok {
		return errors.New("Key not found in memory")
	}
	//Replace the tag from the memory in the pattern
	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `<get name="`+getStruct.Name+`"/>`, content, -1)
	//Return
	return nil
}

//Function used to process the bot tag
func (aimlTemplate *AIMLTemplate) ProcessBot(aiml *AIML) error {
	//Structure for the bot tag
	botStruct := struct {
		XMLName xml.Name `xml:"bot"`
		Name    string   `xml:"name,attr"`
	}{}
	//Parse the content into the struct
	err := xml.Unmarshal([]byte(aimlTemplate.Content), &botStruct)
	//Return the error if there was any
	if err != nil {
		return err
	}
	//Try to get the value from the memory
	content, ok := aiml.Bot[botStruct.Name]
	//Return the error if there was any
	if !ok {
		return errors.New("Key not found in memory")
	}
	//Replace the tag from the memory in the pattern
	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `<bot name="`+botStruct.Name+`"/>`, content, -1)
	//Return
	return nil
}

//Function used to process the bot tag
func (aimlTemplate *AIMLTemplate) ProcessSrai(aiml *AIML) (*AIMLTemplate, error) {
	//Structure for the srai tag
	sraiStruct := struct {
		XMLName xml.Name `xml:"srai"`
		Content string   `xml:",innerxml"`
	}{}
	//Parse the content into the struct
	err := xml.Unmarshal([]byte(aimlTemplate.Content), &sraiStruct)
	//Return the error if there was any
	if err != nil {
		return nil, err
	}
	//Remove the srai tag from the content
	sraiStruct.Content = strings.Replace(sraiStruct.Content, `<srai>`, "", -1)
	sraiStruct.Content = strings.Replace(sraiStruct.Content, `</srai>`, "", -1)
	//Try to find the reference tag from the srai
	ret, errPattern := aiml.findPattern(sraiStruct.Content, true)
	//Return the error if there was any
	if errPattern != nil {
		return nil, errPattern
	}
	//Retrun the referenced template
	return ret, nil
}

//Function used to get a random number within given boundaries
func random(min, max int) int {
	//Sedd the number
	rand.Seed(time.Now().Unix())
	//Get a random number within the boundaries
	return rand.Intn(max-min) + min
}

//Function used to process the random tag
func (aimlTemplate *AIMLTemplate) ProcessRandom(aiml *AIML) error {
	//Structurefor the content of the random tag
	randomStruct := struct {
		XMLName xml.Name `xml:"random"`
		List    []struct {
			XMLName xml.Name `xml:"li"`
			Content string   `xml:",innerxml"`
		} `xml:"li"`
	}{}
	//Try to parse the tag from the AIML into the struct
	err := xml.Unmarshal([]byte(aimlTemplate.Content), &randomStruct)
	//Return the error if there was any
	if err != nil {
		return err
	}
	//Get a random index number
	randIdx := random(0, len(randomStruct.List))
	//Get the random content
	randContent := randomStruct.List[randIdx]
	//Set the content for the tehmplate
	aimlTemplate.Content = randContent.Content
	//Create an empty string array for the matches
	arr := []string{}
	//Process the rest of the tags in the content. Get,set,bot,star,srai
	_, errT := aiml.processTemplateTags(aimlTemplate, arr, true)
	//Return the error if there was any
	if errT != nil {
		return errT
	}
	//Return
	return nil
}

//Function used to process the star tag
func (aimlTemplate *AIMLTemplate) ProcessStar(starContent []string) {
	//Loop the content
	for idx, sContent := range starContent {
		//If it is not the first.(The first one is always the full sentence the rest is capture from the regex)
		if idx > 0 {
			//Clean the string and change reflections in it
			sContent = ProcessStarContent(strings.TrimSpace(sContent))

			//Edited, remplace all count of <star/>
			aimlTemplate.Content = strings.Replace(aimlTemplate.Content, "<star/>", sContent, -1)
		}
	}
	//Replace the star tag with an empty string
	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, "<star/>", "", -1)
}

//Function used to process the star tags after it was replaced with the regex capture
func ProcessStarContent(content string) string {
	//Regex to remove ,;!.? from he input string
	reg, err := regexp.Compile("[,;!.?]+")
	if err == nil {
		//Remove the commas and semicolons dots and questionmarks
		content = reg.ReplaceAllString(content, "")
	}

	//Split the sentence
	sentence := strings.Split(content, " ")
	//Loop the sentence
	for i, word := range sentence {
		//Try to swap the word by key
		if new_word, ok := Reflections[word]; ok {
			//If it could be swapped then do it
			sentence[i] = new_word
		}
	}
	return strings.Join(sentence, " ")
}
