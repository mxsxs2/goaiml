package goaiml

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

//Function used to read in the AIML file
func (aiml *AIML) Learn(mainFile string) error {
	//Try to open the file
	xmlFile, err := os.Open(mainFile)
	//Return error if the file could not be opened
	if err != nil {
		return err
	}
	//Close the file
	defer xmlFile.Close()
	//Read in the content of the file
	bytes, _ := ioutil.ReadAll(xmlFile)
	//Parse the xml and return back
	return xml.Unmarshal(bytes, &aiml.Root)
}

//Function used to pre process the input for better pattern match
func (aiml *AIML) PreProcessInput(input string) string {
	//The processed Sentence
	processedSentence := strings.Split(strings.ToLower(input), " ")
	//Loop the input text
	for i, word := range processedSentence {
		//Try to process the word
		if processed, ok := PreProcessWords[word]; ok {
			//If it could be rpocessed then set it
			processedSentence[i] = processed
		}
	}
	//Return the pre processed sentence
	return strings.Join(processedSentence, " ")
}

//Function used to get a response form the AIML file
func (aiml *AIML) Respond(input string) (string, error) {
	//Find a template by matching pattern
	aimlTemplate, err := aiml.findPattern(aiml.PreProcessInput(input), false)
	//If there was a error then return an empty string
	if err != nil {
		return "", err
	}

	//fmt.Println(aimlTemplate.Content)
	//If the template countains the srai redierct tag and the reference could not be found, return an error
	if strings.Contains(aimlTemplate.Content, "<srai") {
		return "", errors.New("Srai reference not found")
	}
	//Return the response string
	return strings.TrimSpace(aimlTemplate.Content), nil
}

//Function used to process the AIML tags
func (aiml *AIML) processTemplateTags(template *AIMLTemplate, matchRes []string, looped bool) (*AIMLTemplate, error) {
	//Check if there is a star tag
	if strings.Contains(template.Content, "<star") { //Star is for replacing atched string
		template.ProcessStar(matchRes)
	}
	//Check if there is a set tag
	if strings.Contains(template.Content, "<set") { //Set is to set a variable
		template.ProcessSet(aiml)
	}
	//Check if there is a get tag
	if strings.Contains(template.Content, "<get") { //Get is to get a variable
		template.ProcessGet(aiml)
	}
	//Check if there is a bot tag
	if strings.Contains(template.Content, "<bot") { //Bot to use the bots name
		template.ProcessBot(aiml)
	}
	//Check if there is a srai tag
	if strings.Contains(template.Content, "<srai") && !looped { //Srai to redirect to another category
		return template.ProcessSrai(aiml)
	}
	//Check if there is a random tag
	if strings.Contains(template.Content, "<random") { //Rabdom is to get a random ansfer from a given category
		template.ProcessRandom(aiml)
	}

	//Return the template
	return template, nil
}

//Function used to find a matching patter to a given text input
func (aiml *AIML) findPattern(input string, looped bool) (*AIMLTemplate, error) {
	//Loop each category
	for _, category := range aiml.Root.Categories {
		//Add padding ot the input text
		input = " " + input + " "
		//Check if the ctegory content contains the "<bot" keyword
		if strings.Contains(category.Pattern.Content, "<bot") {
			//If it icontains the process the bot
			category.Pattern.ProcessBot(aiml)
		}
		//Modified strip input from extra spaces. srai wouldnt work otherwise
		matchRes := category.Pattern.Regexify().FindStringSubmatch(strings.TrimSpace(input))
		//If there was any match
		if len(matchRes) > 0 {
			//Process the templates and return the response
			return aiml.processTemplateTags(&category.Template, matchRes, looped)
		}
	}
	//Return an error response
	return nil, errors.New("Template not found")
}
