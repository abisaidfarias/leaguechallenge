# leaguechallenge


The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
