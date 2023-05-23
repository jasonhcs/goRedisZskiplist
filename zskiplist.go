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
	zsl.header.Level = make([]ZskiplistLevel, ZSKIPLIST_MAXLEVEL)
	for i := 0; i < ZSKIPLIST_MAXLEVEL; i++ {
		zsl.header.Level[i].Forward = nil
		zsl.header.Level[i].Span = 0
	}
	zsl.header.Backward = nil
	zsl.tail = nil
	return zsl
}

// 从小到大排序列表
func (this *Zskiplist) UpNodeList() []*ZskiplistNode {
	nodeList := []*ZskiplistNode{}
	curNode := this.header.Level[0].Forward
	for curNode != nil {
		nodeList = append(nodeList, curNode)
		curNode = curNode.Level[0].Forward
	}
	return nodeList
}

// 从大到小排序列表
func (this *Zskiplist) DownNodeList() []*ZskiplistNode {
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
	return this.header.Level[0].Forward
}

// 尾结点(最大)
func (this *Zskiplist) TailObj() *ZskiplistNode {
	return this.tail
}

// 节点数量
func (this *Zskiplist) Length() uint64 {
	return this.length
}

// 当前最大层数
func (this *Zskiplist) Level() int {
	return this.level
}

func (this *Zskiplist) RemoveHeader() {
	if this.length == 0 {
		return
	}
	x := this.header.Level[0].Forward
	if x == nil {
		return
	}
	for i := 0; i < this.level; i++ {
		if this.header.Level[i].Forward != x {
			break
		}
		this.header.Level[i].Forward = x.Level[i].Forward
	}
	if x.Level[0].Forward != nil {
		x.Level[0].Forward.Backward = nil
	} else {
		this.tail = x.Backward
	}
	for this.level > 1 && this.header.Level[this.level-1].Forward == nil {
		this.level--
	}
	this.length--
}

func (this *Zskiplist) RemoveTail() {
	if this.length == 0 {
		return
	}
	x := this.tail
	if x == nil {
		return
	}
	for i := 0; i < this.level; i++ {
		if this.header.Level[i].Forward == nil || this.header.Level[i].Forward == x {
			this.header.Level[i].Forward = nil
		} else {
			break
		}
	}
	if x.Backward != nil {
		x.Backward.Level[0].Forward = nil
		this.tail = x.Backward
	} else {
		this.header.Level[0].Forward = nil
		this.tail = nil
	}
	for this.level > 1 && this.header.Level[this.level-1].Forward == nil {
		this.level--
	}
	this.length--
}

func (this *Zskiplist) RemoveNode(node *ZskiplistNode) {
	update := make([]*ZskiplistNode, ZSKIPLIST_MAXLEVEL)
	x := this.header
	for i := this.level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil && x.Level[i].Forward != node {
			x = x.Level[i].Forward
		}
		update[i] = x
	}
	if x.Level[0].Forward != nil && x.Level[0].Forward == node {
		for i := this.level - 1; i >= 0; i-- {
			if update[i].Level[i].Forward == node {
				update[i].Level[i].Forward = node.Level[i].Forward
			}
		}
		for this.level > 1 && this.header.Level[this.level-1].Forward == nil {
			this.level--
		}
		this.length--
	}
}
