# gourl
__Author__ - _Brandon Mandzik, 2020_ <br>
_Gourl_  is a cli tool that can either print a html-body or header information to the user.

### Flags usage 
```go
$ gourl -o htmlBody.htm https://golang.org // it will also create new directory if needed
$ gorul -header https://golang.org // it's a bool flag
```

__Requieres__: go 1.13 AND darwin system

__Installation__:
cd into the directory. use $ go install and execute in the bin directory with $ go run gourl or call it globally by exporting into your PATH variable.

