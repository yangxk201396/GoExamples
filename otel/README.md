# OpenTelemetry


### 1. OpenTelemetry 推荐
    Prometheus + Grafana做Metrics存储、展示
    使用Jaeger或Zipkin做分布式跟踪的存储和展示
    使用Elasticsearch做日志存储和展示

### 2. OpenTelemetry 目标
    实现Metrics、Tracing、Logging的融合，作为APM的数据采集终极解决方案
    Tracing：提供了一个请求从接收到处理完成整个生命周期的跟踪路径，也被称为分布式链路追踪
    Metrics：例如cpu、请求延迟、用户访问数等Counter、Gauge、Histogram指标
    Logging：传统的日志，提供精确的系统记录

### 3. OpenTelemetry 使用
    基于Metrics告警发现异常，通过Tracing定位到具体的系统和方法
    根据模块的日志最终定位到错误详情和根源
    调整Metrics等设置，更精确的告警/发现问题