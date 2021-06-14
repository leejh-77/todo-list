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

type workspaceMemberTable struct {
	*orm.ORMTable
}

func (t *workspaceMemberTable) FindByUserIdAndWorkspaceId(m *WorkspaceMember, uid int64, wid int64) error {
	 return t.Find(m, "userId = ? AND workspaceId = ?", uid, wid)
}

func (t *workspaceMemberTable) DeleteByWorkspaceId(wid int64) error {
	return t.Delete("workspaceId = ?", wid)
}

func (t *workspaceMemberTable) FindByWorkspace(ms *[]WorkspaceMember, wid int64) error {
	return t.Find(ms, "workspaceId = ?", wid)
}
