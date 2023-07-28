
# Cobra

A Framework for Modern CLI Apps in Go

[Official Document](https://cobra.dev)

## Concepts



Cobra is built on a structure of commands, arguments & flags.

Commands represent actions, Args are things and Flags are modifiers for those actions.

The best applications will read like sentences when used. Users will know how to use the application because they will natively understand how to use it.

The pattern to follow is `APPNAME VERB NOUN --ADJECTIVE`. or `APPNAME COMMAND ARG --FLAG`

A few good real world examples may better illustrate this point.

In the following example, ‘server’ is a command, and ‘port’ is a flag:

hugo server --port=1313
In this command we are telling Git to clone the url bare.

git clone URL --bare

Commands

Command is the central point of the application. Each interaction that the application supports will be contained in a Command. A command can have children commands and optionally run an action.

In the example above, ‘server’ is the command.

More about cobra.Command

Flags

A flag is a way to modify the behavior of a command. Cobra supports fully POSIX-compliant flags as well as the Go flag package. A Cobra command can define flags that persist through to children commands and flags that are only available to that command.

In the example above, ‘port’ is the flag.

Flag functionality is provided by the pflag library, a fork of the flag standard library which maintains the same interface while adding POSIX compliance.

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

