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

func (this *Zskiplist) NodeList() []*ZskiplistNode {
	nodeList := []*ZskiplistNode{}
	curNode := this.tail
	for curNode != nil {
		nodeList = append(nodeList, curNode)
		curNode = curNode.Backward
	}
	return nodeList
}
