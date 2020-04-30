package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type JoinNewMember struct{}

func (joinNewMember JoinNewMember) handle(b blockInfo, args map[string]interface{}) error {
	contractAddress, ok := args["contractAddress"]
	if !ok {
		return joinNewMember.checkArg("contractAddress")
	}

	memberAddress, ok := args["memberAddress"]
	if !ok {
		return joinNewMember.checkArg("memberAddress")
	}

	err := daoMemberService.Update(b.height, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex(), model.MemberStatusAgreed)
	if err != nil {
		return err
	}

	err = daoTodoListService.Release(b.height, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex(), model.TodoTypeInviteMember)
	if err != nil {
		return err
	}

	return nil
}

func (joinNewMember JoinNewMember) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao join new member event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
