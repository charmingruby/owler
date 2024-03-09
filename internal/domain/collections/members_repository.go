package collections

type CollectionMembersRepository interface {
	Create(member *CollectionMember) error
	FindMemberInCollection(accountID, collectionID string) (CollectionMember, error)
}
