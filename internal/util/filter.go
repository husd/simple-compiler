package util

/**
 * 过滤器接口
 * @author hushengdong
 */
type Filter interface {

	/**
	 * 满足要求就返回True
	 */
	Accept(t interface{}) bool

	Filter_()
}
