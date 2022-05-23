### 1. 默认web界面
    http://127.0.0.1:6060/debug/pprof/
    下载 profile
### 2. 执行测试用例
    go test -bench=. -cpuprofile=cpu.prof
    输出 cpu.prof
### 3. 可视化
    brew install graphviz
    go get -u github.com/google/pprof
    pprof -http=:8080 cpu.prof
    pprof -http=:8080 profile