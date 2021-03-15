# Flipper
  
Flipper was designed for personal usage to take a list of potential web server file names, utilise a range of techniques to mutate the values and (hopefully) generate a list of filenames that could exist on the target. Ideally, this will help with content discovery during the enumeration phase.  

Techniques:  
* dashes to underscores
* underscores to dashes
* '1337' speak (letters to numbers) [not fully working]
  
## Installation  
  
```
go get github.com/crawl3r/flipper    
```
  
## Usage  
  
```
skid@life:~$ cat test.txt 
file
another_file


skid@life:~$ cat test.txt| ./flipper -q | uniq -u
file
fi1e
fi13
fil3
another_file
another-file
4nother_file
4n0ther_file
4n07her_file
4n07h3r_file
4n07h3r_fi1e
4n07h3r_fi13
4n07h3r_fil3
4n07her_fi1e
4n07h3r_fi1e
4n0th3r_file
4n0th3r_fi1e
4n0th3r_fi13
4n0th3r_fil3
4n0ther_fi1e
4n0th3r_fi1e
4no7her_file
4no7h3r_file
4no7h3r_fi1e
4no7h3r_fi13
4no7h3r_fil3
4no7her_fi1e
4no7h3r_fi1e
4noth3r_file
4noth3r_fi1e
4noth3r_fi13
4noth3r_fil3
4nother_fi1e
4noth3r_fi1e
an0ther_file
an07her_file
an07h3r_file
an07h3r_fi1e
an07h3r_fi13
an07h3r_fil3
an07her_fi1e
an07h3r_fi1e
an0th3r_file
an0th3r_fi1e
an0th3r_fi13
an0th3r_fil3
an0ther_fi1e
an0th3r_fi1e
ano7her_file
ano7h3r_file
ano7h3r_fi1e
ano7h3r_fi13
ano7h3r_fil3
ano7her_fi1e
ano7h3r_fi1e
anoth3r_file
anoth3r_fi1e
anoth3r_fi13
anoth3r_fil3
another_fi1e
anoth3r_fi1e 
```

## Known Bug  
  
1337 rule seems okay if the target word only has one instance of the replaceable letter. I haven't perfected the recursive replace to account for multiple of the same letters, ensuring all patterns are generated.  
  
## License  
  
I'm just a simple skid. Licensing isn't a big issue to me, I post things that I find helpful online in the hope that others can:  
A) learn from the code  
B) find use with the code or  
C) need to just have a laugh at something to make themselves feel better  
  
Either way, if this helped you - cool :)  
