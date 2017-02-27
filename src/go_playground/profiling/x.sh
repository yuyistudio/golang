if [[ "$1" == "run" ]]; then
	# 编译运行,生成profile记录文件
	go build && ./profiling -cpuprofile=output.prof
elif [[ "$1" == "prof" ]]; then
	# 将运行的bin和生成的prof文件作为参数,打开prof分析
	go tool pprof profiling output.prof

	# 命令:

	# top10 显示sample最多的10个函数,按照sample数量排序.一个sample指的是在一次sample中程序所在的函数计数加一.
	# 每次sample时,会累加当前函数的计数,还会累加当前调用栈上所有函数的计数(函数的累加计数)
	# top10 -cum 按照累加计数进行排序 cumulation

	# web命令:调用Graphviz进行绘图,生成svg文件(Scalable Vector Graphics),然后使用系统默认程序设置打开该文件.
	# 所以首先要安装Graphviz: sudo brew install graphviz
	# 默认软件不支持SVG格式时,可以修改.SVG文件的默认打开方式为chrome或者其他支持SVG格式的软体
else
	echo "error. help: go prof/run"
	exit 1
fi
