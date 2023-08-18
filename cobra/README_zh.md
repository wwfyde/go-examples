# Cobra
A Framework for Modern CLI Apps in Go

现代CLI应用的Go框架

[Official Document](https://cobra.dev)



# 核心概念

Cobra 构建在命令, 参数和标记的结构之上.

命令代表行为, 参数代表事情, 标记是操作的修改器.


最好的应用在使用的时候应该像是阅读语句一样. 用户将知道如何使用应用, 因为他们会天然地理解如何使用它.

遵守的模式是 `APPNAME VERB NOUN --ADJECTIVE`. 或 `APPNAME COMMAND ARG --FLAG`. `名称-动词-名词 --形容词` 或 `名称-命令-参数 --标记`. 

一些好的现实例子可能更好地说明.

下面的例子中, "server"是一个命令, "port"是一个标记":

`hugo server --port=1313`

下面这个例子我们告诉git去clone这个url bare:

`git clone URL --bare`

# 命令(command)


命令式应用的中心点.所有应用支持的交互行为都包含在一条命令中. 一条命令可以包含子命令和可选地运行一个行为. 
上面的例子中, server表示命令.

更多参见 [cobra.Command](https://pkg.go.dev/github.com/spf13/cobra?utm_source=godoc#Command)

# 标记(Flags)


标记是一种修改命令行为的方式. Cobra完全支持 POSIX-compliant 风格的标记和Go fla包风格.
上面的例子中, port表示标记.

flag的功能性有pflag库提供, pflag 是flag标准库的一个分叉, pflag维护了相同的接口但
增加了对POSIX的兼容.


