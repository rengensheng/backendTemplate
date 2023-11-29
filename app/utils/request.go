package utils

import (
	"strings"

	"github.com/hyahm/golog"
	"xorm.io/xorm"
)

type Request struct {
	PageNo   int               `form:"pageNo" json:"pageNo"`
	PageSize int               `form:"pageSize" json:"pageSize"`
	SortItem string            `form:"sortItem" json:"sortItem"`
	SortType string            `form:"sortType" json:"sortType"`
	Query    map[string]string `form:"query" json:"query"`
}

func (req *Request) DisposeRequest(session *xorm.Session) *xorm.Session {
	var disposedSession *xorm.Session = session
	for k, v := range req.Query {
		queryStr := ""
		if strings.HasPrefix(v, "*") {
			// start like query
			queryStr += "%" + strings.TrimPrefix(v, "*")
			if strings.HasSuffix(v, "*") {
				queryStr = strings.TrimSuffix(queryStr, "*") + "%"
			}
			disposedSession = session.Where(k+" like ?", queryStr)
		} else if strings.HasSuffix(v, "*") {
			queryStr = strings.TrimSuffix(v, "*") + "%"
			disposedSession = session.Where(k+" like ?", queryStr)
		} else if strings.Contains(v, ",") {
			disposedSession = session.In(k, strings.Split(v, ","))
		} else {
			disposedSession = session.Where(k+" = ?", v)
		}
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.SortItem == "" {
		req.SortItem = "created_time"
	}
	if req.SortType == "" {
		req.SortType = "desc"
	}
	golog.Info("sort method", req.SortItem, req.SortType)
	if req.SortType == "desc" {
		golog.Info("desc")
		disposedSession = disposedSession.Desc(req.SortItem)
	} else {
		disposedSession = disposedSession.Asc(req.SortItem)
	}
	disposedSession = disposedSession.Limit(req.PageSize, (req.PageNo-1)*req.PageSize)
	return disposedSession
}
