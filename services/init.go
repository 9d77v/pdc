package services

import "github.com/9d77v/go-lib/ptrs"

//GetPageInfo 获取分页信息
func GetPageInfo(page, pageSize *int64) (int64, int64) {
	offset := ptrs.Int64(page)
	limit := ptrs.Int64(pageSize)
	if offset < 1 {
		offset = 1
	}
	if limit == 0 {
		limit = 10
	}
	if limit < 0 {
		limit = -1
	}
	if limit > 100 {
		limit = 100
	}
	offset = (offset - 1) * limit
	return offset, limit
}
