package statistic

// RequestNotify 在新连接进入统计管理器时触发，供无 HTTP Controller 的原生客户端记录请求。
type RequestNotify func(c Tracker)

var DefaultRequestNotify RequestNotify
