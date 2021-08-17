# simple-compiler

编译器教练机，为了学习编译原理用go语言写的一个简单的编译器。 目录里如果有1个Keep.go，就是为了让目录里有个文件，这样就能把目录提交到github了， 之后会删除这个无用的文件。

这个项目，最终目的是实现一个javac的功能，将Java源代码编译为.class文件。

这里简要列举一下编译器的几个步骤：

- 源程序
- 词法分析器
- 语法分析器
- 语义分析器
- 中间代码生成器
- 代码优化器
- 代码生成器
- 目标程序

全程有2个模块，符号表+错误处理器 是在每个阶段都参与的。

这个项目的初步目标是做到 中间代码生成器 这一阶段。

错误处理器需要详细说一下：

- 词法分析阶段能够检测出输入中不能形成源语言任何记号的错误字符串。
- 语法分析阶段 分析下语法规则的错误
- 语义分析阶段 例如将2个标识符相加， 一个数组 一个过程

## 1.0 一些笔记

这里列举下，编译原理这本书中提到的一些关于编译器的经典语句：

- 发现错误就立即停止运行的编译器不是一个好的编译器。

## 1.1 词法分析器

第一步的目的是实现一个词法分析器。

### 1.1.1 简单的分割字符串

把目标再缩小下，编写一个方法，入参是字符串，按空格和换行符分割，返回所有的token， 这样就实现了一个最简单的词法分析器。代码见 word_lexical_v1.go

### 1.1.2 读取文件

有了最简单的分割字符串的功能，下一步就是处理文件和目录了，处理目录暂且不开发，先开发编译单个文件的程序，词法分析的初步目标是解析出token，这个 token已经在 word_lexical_v3.go 中实现了部分功能

- 去掉空格
- 去掉了注释
- 识别出了 标识符、赋值、数字、字符串等


