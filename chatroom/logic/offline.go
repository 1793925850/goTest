package logic

import (
	"container/ring" // ring 实现了环形链表的操作

	"github.com/spf13/viper"
)

type offlineProcessor struct {
	n int

	// 保存所有用户最近的 n 条消息
	recentRing *ring.Ring

	// 保存某个用户离线消息（一样 n 条）
	userRing map[string]*ring.Ring
}

var OfflineProcessor = newOfflineProcessor()

func newOfflineProcessor() *offlineProcessor {
	n := viper.GetInt("offline-num")

	return &offlineProcessor{
		n:          n,
		recentRing: ring.New(n), // 创建了含有 n 个链表节点的环形链表
		userRing:   make(map[string]*ring.Ring),
	}
}

func (o *offlineProcessor) Save(msg *Message) {
	if msg.Type != MsgTypeNormal {
		return
	}

	o.recentRing.Value = msg
	o.recentRing = o.recentRing.Next() // 跳到写一个链表节点

	for _, nickname := range msg.Ats {
		nickname = nickname[1:] // 这里为什么从一开始？？？

		var (
			r  *ring.Ring
			ok bool
		)

		if r, ok = o.userRing[nickname]; !ok {
			r = ring.New(o.n)
		}

		r.Value = msg
		o.userRing[nickname] = r.Next()
	}
}

func (o *offlineProcessor) Send(user *User) {
	o.recentRing.Do(func(value interface{}) { // Do 按正向顺序调用环形链表的每个元素上的函数 f。
		if value != nil {
			user.MessageChannel <- value.(*Message)
		}
	})

	if user.isNew {
		return
	}

	if r, ok := o.userRing[user.NickName]; ok {
		r.Do(func(value interface{}) {
			if value != nil {
				user.MessageChannel <- value.(*Message)
			}
		})

		delete(o.userRing, user.NickName) // 删除某个用户的离线消息
	}
}
