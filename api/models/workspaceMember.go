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