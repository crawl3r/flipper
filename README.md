# Flipper
  
Flipper was designed for personal usage to take a list of potential web server file names, utilise a range of techniques to mutate the values and (hopefully) generate a list of filenames that could exist on the target. Ideally, this will help with content discovery during the enumeration phase.  

Techniques:  
* dashes to underscores
* underscores to dashes
* '1337' speak (letters to numbers)
  
## Installation  
  
```
go get github.com/crawl3r/flipper    
```
  
## Usage  
  
```
skid@life:~$ cat filenames.txt | ./flipper  
```
  
## License  
  
I'm just a simple skid. Licensing isn't a big issue to me, I post things that I find helpful online in the hope that others can:  
A) learn from the code  
B) find use with the code or  
C) need to just have a laugh at something to make themselves feel better  
  
Either way, if this helped you - cool :)  
