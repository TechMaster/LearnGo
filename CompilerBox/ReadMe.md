Chạy lệnh Go

``` 
micro call compiler Compile.Run '{"SourceCode": [ "package main", "import \"fmt\"", "func main() {", " fmt.Println(\"Hello\")", "}" ], "Language": "go"}'
```

Chạy lệnh Node.js
```
micro call compiler Compile.Run '{"SourceCode": [ "console.log(\"Chao cac ban\")" ], "Language": "js"}'
```

```
micro call compiler Compile.Run '{"SourceCode": [ "for (i = 0; i < 10; i++) {", "console.log(i);", "}" ], "Language": "js"}' 
```