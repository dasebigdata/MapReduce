mkdir -p ../target
# 编译mapreduce框架
javac ./mr/*.java -d ../target
mkdir -p ../target_udf
# 编译wordcount
javac -cp ../target -d ../target_udf ../udf/*.java
# 运行wordcount
java -cp ../target mr.Main ../../input.txt ../../output.txt ../target_udf/ udf.WordCountMapper udf.WordCountReducer
