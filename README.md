# Introduction
Hey! yobuffer! No it's your buffer.

It contains a custom-friendly [IDL] and some data interchange formats, which is inspired by protobuf. Of cousrse it's not a replacement of anything, but just a choose for super fast serialize/unserialize in gamedev.

# Long-winded introduction(in Chinese)
本项目旨在开发一个定制性更高的 [IDL] 和一些性能更好的数据交换格式，专为游戏开发中常见的高频率对象复制做优化。相比 protobuf, thift 的实现并不见得多先进，仅仅是去掉一些不必要的高级特性。


# Status
heavy in dev

# Usage

```bash
# Install /usr/local/bin/yobuffer
make && sudo make install

# Compile IDL file
yobuffer compile --src _examples/rpc.yb
```

# Benchmarks
NOTE: It's so fast just because it lacking a lot features like protobuf.  
```bash
$ cd ./_examples
$ go test -bench .
goos: windows
goarch: amd64
pkg: github.com/cupen/yobuffer/_examples
cpu: AMD Ryzen 7 3700X 8-Core Processor
BenchmarkEncoding/Encode-16       31008231       38.30 ns/op     1227.21 MB/s       48 B/op       1 allocs/op
BenchmarkEncoding/Decode-16       25000260       54.63 ns/op      860.33 MB/s       16 B/op       3 allocs/op
ok      github.com/cupen/yobuffer/_examples     2.858s
```


[IDL]: https://en.wikipedia.org/wiki/Interface_description_language
