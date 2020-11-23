package utils

import (
	"fmt"
	"gim/server"
	"gim/utils/skiplist"
	"strconv"
	"sync"
)

type WaitAckMsg struct {
	Msg        *server.GimProtocol
	RetryTime  int64 //下一次发起重试的时间
	RetryCount int32 //重试次数
	Uid        int64 //发送给谁
	MsgId      int64
}

func (m *WaitAckMsg) ExtractKey() float64 {
	key, _ := strconv.ParseFloat(fmt.Sprintf("%d.%d", m.RetryTime, m.Uid), 64)
	return key
}

func (m *WaitAckMsg) String() string {
	return fmt.Sprintf("%d.%d", m.RetryTime, m.Uid)
}

type RetryList struct {
	sync.RWMutex
	skipList *skiplist.SkipList
	msgIdMap sync.Map
}

func NewRetryList() *RetryList {
	retryList := RetryList{}
	list := skiplist.New()
	retryList.skipList = &list
	return &retryList
}

const (
	//重试间隔 3s
	RETRY_DURATION = 3
	//重试次数 5次
	RETRY_COUNT = 5
)

func (m *RetryList) AddRetryMsg(msg *server.GimProtocol, uid int64, msgId int64) error {
	m.Lock()
	defer m.Unlock()
	ack := WaitAckMsg{Msg: msg, RetryTime: NowMillisecond() + RETRY_DURATION,
		RetryCount: RETRY_COUNT, Uid: uid, MsgId: msgId}
	m.skipList.Insert(&ack)
	m.msgIdMap.Store(msgId, &ack)
	return nil
}

func (m *RetryList) RemoveRetryMsg(msgId int64) error {
	m.Lock()
	defer m.Unlock()
	ack, ok := m.msgIdMap.Load(msgId)
	if !ok {
		return nil
	}
	ackMsg := ack.(*WaitAckMsg)
	m.skipList.Delete(ackMsg)
	return nil
}

//获取等待重试的消息，如果发现retryCount=0则删除，否则删除老消息并更新retryCount跟retryTime后重新插入
func (m *RetryList) GetWaitRetryMsg() (list []*WaitAckMsg, err error) {
	now := WaitAckMsg{RetryTime: NowMillisecond()}
	mid, ok := m.skipList.FindGreaterOrEqual(&now)
	if ok {
		pre := m.skipList.Prev(mid)
		//如果已经到了尾部了，则证明循环结束了
		if pre == m.skipList.GetLargestNode() {
			len := len(list)
			newList := make([]*WaitAckMsg, len)
			for k, v := range list {
				newList[len-k-1] = v
			}
			return newList, nil
		}
		msg, _ := pre.GetValue().(*WaitAckMsg)
		list = append(list, msg)
	}
	res := make([]*WaitAckMsg, 0, len(list))
	//遍历所有已经到期的消息，检查每个消息，如果已经到达重试最大次数，则直接删除，如果没有达到最大次数，修改下次重试时间后插入
	for _, v := range list {
		m.skipList.Delete(v)
		if v.RetryCount <= 0 {
			m.msgIdMap.Delete(v.MsgId)
		} else {
			v.RetryCount--
			v.RetryTime += RETRY_DURATION
			m.skipList.Insert(v)
		}
		res = append(res, v)
	}
	return res, nil
}
