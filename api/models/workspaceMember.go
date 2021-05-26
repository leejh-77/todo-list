package models

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

func SaveWorkspaceMember(m *WorkspaceMember) (int64, error) {
	ret, err := DB.Exec("INSERT INTO workspaceMembers (type, workspaceId, userId) VALUES (?, ?, ?)",
		m.Type, m.WorkspaceId, m.UserId)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}
