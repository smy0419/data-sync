package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type InviteNewPresident struct{}

func (inviteNewPresident InviteNewPresident) handle(b blockInfo, args map[string]interface{}) error {
	contractAddress, ok := args["contractAddress"]
	if !ok {
		return inviteNewPresident.checkArg("contractAddress")
	}

	memberAddress, ok := args["memberAddress"]
	if !ok {
		return inviteNewPresident.checkArg("memberAddress")
	}

	err := daoTodoListService.Insert(b.height, b.time, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex(), model.TodoTypeInvitePresident, -1, model.MemberRolePresident)
	if err != nil {
		return err
	}

	return nil
}

func (inviteNewPresident InviteNewPresident) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao invite new president event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
