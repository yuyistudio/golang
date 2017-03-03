#!/usr/bin/env bash
if [[ "$1" == "run" ]]; then
	# 编译运行,生成profile记录文件
	go build && ./profiling -cpuprofile=cpu.prof -memprofile=mem.prof
elif [[ "$1" == "cpu" ]]; then
	# 将运行的bin和生成的prof文件作为参数,打开prof分析
	go tool pprof profiling cpu.prof

	# 命令:

	# top10 显示sample最多的10个函数,按照sample数量排序.一个sample指的是在一次sample中程序所在的函数计数加一.
	# 每次sample时,会累加当前函数的计数,还会累加当前调用栈上所有函数的计数(函数的累加计数)
	# top10 -cum 按照累加计数进行排序 cumulation

	# web命令:调用Graphviz进行绘图,生成svg文件(Scalable Vector Graphics),然后使用系统默认程序设置打开该文件.
	# 所以首先要安装Graphviz: sudo brew install graphviz
	# 默认软件不支持SVG格式时,可以修改.SVG文件的默认打开方式为chrome或者其他支持SVG格式的软体

	# 图像中线的粗线表示该调研耗时的长短
	# 方框的大小表示该函数的真实耗时(不包含调用的其他函数)
	# 图像中每个方块的含义:
	# main.delay   1.2s of 1.99s(81%)
	# 函数名   当前函数的执行时间(真实时间)  当前函数执行时间+调用的其他函数的执行时间(总时间)  该时间和占据程序总运行时间的比例
	# 所以性能瓶颈就在"真实时间"最多的那个函数

	# web fn_name
	# 绘图时只考虑指定的函数

	# list fn_name
	# 显示该函数每一行代码的耗时
elif [[ "$1" == "mem" ]]; then
	flogs='--inuse_objects' # 显示内存分配次数,而不是占用的内存总量
	flags='--nodefraction=0.5' # 消耗比例小于某个阈值的节点不显示
	#flags=''
	go tool pprof $flags profiling mem.prof
else
	echo "error"
	exit 1
fi
