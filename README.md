# MapReduce
本实验需要实现一个简易版 mapreduce 框架，不限语言。

我们提供了单机版 mapreduce 框架在三种语言下的实现供大家参考，并基于该框架实现了 wordcount 程序。

项目结构：
```
.
├── golang                                  # golang版本mapreduce实现
│   ├── go.mod
│   ├── main.go                             # mapreduce框架主体
│   ├── mr
│   │   ├── mapper.go                       # map函数类型定义
│   │   ├── pair.go
│   │   └── reducer.go                      # reduce函数类型定义
│   ├── run.sh                              # wordcount运行脚本
│   └── udf
│       ├── wordcount_mapper.go             # wrodcount的map函数
│       └── wordcount_reducer.go            # wordcount的reduce函数
├── input.txt                               # 测试数据
|
|
├── java                                    # java版本mapreduce实现
│   ├── src
│   │   ├── mr
│   │   │   ├── Main.java                   # mapreduce框架主体
│   │   │   ├── Mapper.java                 # map接口
│   │   │   ├── Pair.java                   
│   │   │   └── Reducer.java                # reduce接口
│   │   └── run.sh                          # wordcount运行脚本
│   └── udf
│       ├── WordCountMapper.java            # wrodcount的map方法
│       └── WordCountReducer.java           # wordcount的reduce方法
|
|
├── python                                  # python版本mapreduce实现
│   ├── main.py                             # mapreduce框架主体
│   ├── mapper.py                           # map接口
│   ├── reducer.py                          # reduce接口
│   ├── run.sh                              # wordcount运行脚本
│   ├── save.py         
│   ├── wordcount_mapper.py                 # wrodcount的map方法
│   └── wordcount_reducer.py                # wordcount的reduce方法
└── README.md
```

运行方式：

```
# golang版本
cd ./golang
sh run.sh

# java版本
cd ./java/src
sh run.sh

# python版本
cd ./python
sh run.sh
```

运行过程：

1. 加载用户自定义的 map 和 reduce 函数；
2. 读取输入文件；
3. map -> shuffle -> reduce；
4. 写入输出文件。

## 基础要求

1. 实现伪分布式 mapreduce 框架（可基于提供的任意单机版框架进行扩展，也可以自由实现）：
    
    a. 包含 map、shuffle 和 reduce 三个阶段，提供可编程的 map 和 reduce 接口；
    
    b. master（JobTracker）和 worker（TaskTracker）运行在不同的进程中。

2. 端到端测试：基于伪分布式框架实现 wordcount 程序。

## 进阶要求

1. 实现分布式 mapreduce 框架：JobTracker 和 TaskTracker 运行在不同的机器上。

2. 端到端测试：基于分布式框架实现 wordcount 程序。
