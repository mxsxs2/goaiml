# AIML (Artificial Intelligence Modelling Language) parser in Go-Lang
It is a fork of https://github.com/eduardonunesp/goaiml with a few modifications.

## How to use with it
In command line write ```go get github.com/mxsxs2/goaiml```. When it is finished.
The ```github.com/mxsxs2/goaiml``` package is going to be available in your Go projects.

## Modifications
* Added comment for every statement to support better understanding of the library
* ```processStar``` function now replaces every occurance of the ```<star/>``` tag in the templates before the random
* Added reflection support to the ```<star/>``` tag
* Added pre processor for better pattern matching
* Fix recognition of ```<srai>``` tags

## How it works
1. The parser loads in the AIML file on every request, which is helpful since the server does not have to be restarted if the AIML file is changed.
2. When a request is sent to the parser it searches through the categories in the AIML file top to bottom.
3. Once match found the tags (srai,star,bot,set) are parsed.
4. On a random tag the parser choses one item and returns back to the fucntion which called the AIML parser.
