package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"github.com/AsimovNetwork/data-sync/library/mongo/service"
)

type InviteNewMember struct{}

var daoMemberService = service.DaoMemberService{}
var daoTodoListService = service.DaoTodoListService{}

func (inviteNewMember InviteNewMember) handle(b blockInfo, args map[string]interface{}) error {
	contractAddress, ok := args["contractAddress"]
	if !ok {
		return inviteNewMember.checkArg("contractAddress")
	}

	memberAddress, ok := args["memberAddress"]
	if !ok {
		return inviteNewMember.checkArg("memberAddress")
	}

	err := daoMemberService.Insert(b.height, b.time, b.txHash, contractAddress.(*asimovCommon.Address).Hex(), model.MemberRoleOrdinary, memberAddress.(*asimovCommon.Address).Hex(), model.MemberStatusInvited)
	if err != nil {
		return err
	}

	err = daoTodoListService.Insert(b.height, b.time, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex(), model.TodoTypeInviteMember, -1, model.MemberRoleOrdinary)
	if err != nil {
		return err
	}

	return nil
}

func (inviteNewMember InviteNewMember) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao invite new member event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
