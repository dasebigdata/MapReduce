# 编译wordcount_mapper
go build -buildmode=plugin -o=wordcount_mapper.so ./udf/wordcount_mapper.go
# 编译wordcount_reducer
go build -buildmode=plugin -o=wordcount_reducer.so ./udf/wordcount_reducer.go
# 运行wordcount
go run main.go ../input.txt ../output.txt ./wordcount_mapper.so ./wordcount_reducer.so