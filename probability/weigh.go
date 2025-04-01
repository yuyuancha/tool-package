package probability

import (
	"errors"

	"github.com/yuyuancha/tool-package/random"
)

// Weight 機率權重結構
type Weight struct {
	list  []int64
	total int64
}

// CreateWeight 創建機率權重
func CreateWeight(items ...int64) Weight {
	var w Weight

	w.list = make([]int64, len(items))
	for index := range items {
		w.list[index] = items[index]
		w.total += items[index]
	}

	return w
}

// Add 增加權重
func (w *Weight) Add(item int64) {
	w.list = append(w.list, item)
	w.total += item
}

// GetCount 取得權重數量
func (w *Weight) GetCount() int64 {
	return int64(len(w.list))
}

// GetIndexByWeight 根據權重獲取索引
func (w *Weight) GetIndexByWeight() (int, error) {
	if w.total <= 0 || len(w.list) == 0 {
		return 0, errors.New("empty weight error")
	}

	var result int
	randomValue := random.GetRandomInt(w.total)
	for index := range w.list {
		if randomValue >= w.list[index] {
			randomValue -= w.list[index]
			continue
		}
		result = index
		break
	}

	return result, nil
}
