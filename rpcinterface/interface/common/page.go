/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  page
 * @Version: 1.0.0
 * @Date: 2021/1/11 20:31
 */
package common

import "apiDemo/upgin"

type Paging struct {
	Page     int // 分页
	PageSize int //每页显示多少条
}

func (p *Paging) GetStart() int {
	if p.Page <= 1 {
		return 0
	}
	return (p.Page - 1) * p.GetPageSize()
}

func (p *Paging) GetPageSize() int {
	maxSize := upgin.AppConfig.DefaultInt("page.maxsize", 100)
	if p.PageSize > maxSize {
		p.PageSize = maxSize
	}
	if p.PageSize < 1 {
		if pagesize, err := upgin.AppConfig.Int("page.pagesize"); err == nil {
			return pagesize
		}
	}
	return p.PageSize
}
