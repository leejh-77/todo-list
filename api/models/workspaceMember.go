package models

import "todo-list/orm"

const (
	MemberTypeOwner = 0
	MemberTypeParticipant = 1
)

type WorkspaceMember struct {
	Id int64
	Type int
	WorkspaceId int64
	UserId int64
}

type workspaceMemberQuery struct {
	s orm.Session
}

func WorkspaceMemberQuery(s orm.Session) *workspaceMemberQuery {
	return &workspaceMemberQuery{s}
}

func (q *workspaceMemberQuery) FindByUserIdAndWorkspaceId(m *WorkspaceMember, uid int64, wid int64) error {
	return q.s.Table(TableWorkspaceMember).Find(m, "userId = ? AND workspaceId = ?", uid, wid)
}

func (q *workspaceMemberQuery) FindByWorkspaceId(ms *[]WorkspaceMember, wid int64) error {
	return q.s.Table(TableWorkspaceMember).Find(ms, "workspaceId = ?", wid)
}