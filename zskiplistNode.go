package goRedisZskiplist

type zskiplistLevel struct {
	//节点在该层的下一个节点
	forward *ZskiplistNode
	//节点距离该层下一个节点的距离
	span uint
}

type ZskiplistNode struct {
	//节点内容
	Obj interface{}
	//节点分数（链表按照分数从下到大排序）
	Score float64
	//上一个节点
	Backward *ZskiplistNode
	//下一个节点
	Forward *ZskiplistNode
	//该节点在各层的信息
	Level []zskiplistLevel
}

func createNode(level int, score float64, obj interface{}) *ZskiplistNode {
	zn := new(ZskiplistNode)
	zn.Level = make([]zskiplistLevel, level)
	zn.Score = score
	zn.Obj = obj
	return zn
}
