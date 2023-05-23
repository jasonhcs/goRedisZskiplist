package goRedisZskiplist

const (
	ZSKIPLIST_MAXLEVEL = 32
	ZSKIPLIST_P        = 0.25
)

type Zskiplist struct {
	//头、尾节点
	header, tail *ZskiplistNode
	//节点数量
	length uint64
	//当前最大层数
	level int
}

//创建跳跃表
func Create() *Zskiplist {
	zsl := new(Zskiplist)
	//当前层数置为1
	zsl.level = 1
	//初始化头节点（头节点不用于存储数据）
	zsl.header = new(ZskiplistNode)
	zsl.header.Level = make([]zskiplistLevel, ZSKIPLIST_MAXLEVEL)
	for i := 0; i < ZSKIPLIST_MAXLEVEL; i++ {
		zsl.header.Level[i].forward = nil
		zsl.header.Level[i].span = 0
	}
	zsl.header.Backward = nil
	zsl.tail = nil
	return zsl
}

// 从大到小排序
func (this *Zskiplist) NodeList() []*ZskiplistNode {
	nodeList := []*ZskiplistNode{}
	curNode := this.tail
	for curNode != nil && curNode != this.header {
		nodeList = append(nodeList, curNode)
		curNode = curNode.Backward
	}
	return nodeList
}

// 头结点(最小)
// 注：初始化头结点不存储数据
func (this *Zskiplist) HeaderObj() *ZskiplistNode {
	return this.header.Backward
}

// 尾结点(最大)
func (this *Zskiplist) TailObj() *ZskiplistNode {
	return this.tail
}
