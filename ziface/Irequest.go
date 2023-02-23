package ziface
// 对请求对封装
type IRequest interface{
	// 得到当前的连接
	GetConnection() IConnection
	// 得到请求的消息数据
	GetData() []byte
}