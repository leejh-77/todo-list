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

func (t *workspaceMemberTable) FindByUserIdAndWorkspaceId(w *WorkspaceMember, uid int64, wid int64) error {
	 return t.Find(&w, "userId = ? AND workspaceId = ?", uid, wid)
}

func (t *workspaceMemberTable) DeleteByWorkspaceId(wid int64) error {
	return t.Delete("workspaceId = ?", wid)
}
