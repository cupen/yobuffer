# Introduction
Hey! yobuffer! No, it's your buffer.

It contains a custom-friendly [IDL] and some data interchange formats, which is inspired by protobuf. Of cousrse it's not a replacement of anything, but just a choose for super fast serialize in gamedev.

# Long-winded introduction(in Chinese)
本项目旨在开发易于定制的 [IDL] 和数据交换格式，内置的编解码器专为游戏开发中常见的高频率对象复制做优化。相比 protobuf, thift 的实现并不见得多先进，仅仅是去掉一些不必要的高级特性。

# Status
heavy in dev

# Usage

* IDL (not stable)
  ```
  package example

  @meta(id=10001, encoding=tiny)
  struct PlayerInfo {
      uint64  userId  = 1;
      string  name    = 2;
      uint64  shortId = 3;
      int32   level   = 4;
      float64 winRate = 5;
      bool    isGM = 6;
  }

  @meta(id=10002)
  external protobuf PlayerData

  @meta(id=20001, encoding=tiny, desc="blahblah……")
  rpc GetPlayerInfo() (PlayerInfo)
  ```
* Compiler
  ```bash
  make && sudo make install
  yobuffer compile --src _examples/tiny/rpc.yb
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
