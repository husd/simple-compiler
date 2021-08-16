# simple-compiler

编译器教练机，为了学习编译原理用go语言写的一个简单的编译器。 目录里如果有1个Keep.go，就是为了让目录里有个文件，这样就能把目录提交到github了， 之后会删除这个无用的文件。

这个项目，最终目的是实现一个javac的功能，将Java源代码编译为.class文件。

## 1.1 词法分析器

第一步的目的是实现一个词法分析器。

### 1.1.1 简单的分割字符串

把目标再缩小下，编写一个方法，入参是字符串，按空格和换行符分割，返回所有的token， 这样就实现了一个最简单的词法分析器。代码见 word_lexical_v1.go

### 1.1.2 读取文件

有了最简单的分割字符串的功能，下一步就是处理文件和目录了，处理目录暂且不开发，先开发编译单个文件的程序，词法分析的初步目标是解析出token，这个 token已经在 word_lexical_v3.go 中实现了部分功能

- 去掉空格
- 去掉了注释
- 识别出了 标识符、赋值、数字、字符串等


