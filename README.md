# lexer-parser

北京科技大学计算机专业编译原理实验作业，要求实现基本的**词法分析**和**语法分析**。

## lexer

一个应用于python语言的基本词法分析器。

**基本流程图如下**

![lexer流程图](https://github.com/DingZhenyuan/lexer-parser/blob/master/doc/lexer_pic.png)

## parser

SLR(1)语法分析器题目如下：
> G(E): E->E+T|T <br>
> T->T*F|F <br>
> F->(E)|i

**基本流程图如下**

![parser流程图](https://github.com/DingZhenyuan/lexer-parser/blob/master/doc/parser.png)