package constant

const (
	InitStatus = iota // default status
	SuccessStatus
	FailedStatus
)

// Message Type
const (
	MessageCategoryBeMember         int = 1
	MessageCategoryAddNewMember     int = 2
	MessageCategoryCreateOrg        int = 3
	MessageCategoryBeenRemoved      int = 4
	MessageCategoryRemoveMember     int = 5
	MessageCategoryBePresident      int = 6
	MessageCategoryChangePresident  int = 7
	MessageCategoryTransferAsset    int = 8
	MessageCategoryIssueAsset       int = 9
	MessageCategoryModifyOrgLogo    int = 10
	MessageCategoryModifyOrgName    int = 11
	MessageCategoryNewVote          int = 12
	MessageCategoryProposalRejected int = 13
	MessageCategoryProposalExpired  int = 14
	MessageCategoryInvited          int = 15
	MessageCategoryCloseOrg         int = 16
	MessageCategoryReceiveAsset     int = 17
	MessageCategoryMintAsset        int = 18
)

const (
	MessageStateUnread int = iota
	MessageStateRead
	MessageStateDisagree
	MessageStateAgree
)

const (
	MessageTypeReadOnly int = iota
	MessageTypeDirectAction
	MessageTypeVote
)

const (
	MessagePositionWeb int = iota
	MessagePositionDao
	MessagePositionBoth
)

// Operation Type
const (
	OperationTypeCreateOrg          int = 1
	OperationTypeCloseOrg           int = 2
	OperationTypeModifyOrgName      int = 3
	OperationTypeRemoveMember       int = 4
	OperationTypeIssueAsset         int = 5
	OperationTypeTransferAsset      int = 6
	OperationTypeInviteNewPresident int = 7
	OperationTypeConfirmPresident   int = 8
	OperationTypeInviteNewMember    int = 9
	OperationTypeJoinNewMember      int = 10
	OperationTypeVote               int = 11
	OperationTypeModifyOrgLogo      int = 12
	OperationTypeMintAsset          int = 13
)
