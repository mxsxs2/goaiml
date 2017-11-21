# AIML (Artificial Intelligence Modelling Language) parser in Go-Lang
It is a fork of https://github.com/eduardonunesp/goaiml with a few modifications.

## How to use with it
In command line write ```go get github.com/mxsxs2/goaiml```. When it is finished.
The ```github.com/mxsxs2/goaiml``` package is going to be available in your Go projects.

## Modifications
* Added comment for every statement to support better understanding of the librayr
* ```processStar``` function now replaces every occurance of the ```<star/>``` tag in the templates before the random
* Added reflection support to the ```<star/>``` tag
* Added pre processor for better pattern matching
* Fix recognition of ```<srai>``` tags
