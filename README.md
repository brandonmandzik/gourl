# gourl
__Author__ - _Brandon Mandzik, 2020_ <br>
_Gourl_  is a cli tool that prints the html-body into the standard-ouput. With the flag -o you can specify a output file with path.
with -header you will get printed the url information of a url 
```go
$ gourl -o htmlBody.htm https://golang.org
$ gorul -header https://golang.org
```

__Requieres__: go 1.13 AND darwin system

__Installation__:
cd into the directory. use $ go install and execute in the bin directory with $ go run gocat or call it globally by exporting into your PATH variable.

