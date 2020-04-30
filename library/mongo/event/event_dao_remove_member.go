package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type RemoveMemberEvent struct{}

func (removeMemberEvent RemoveMemberEvent) handle(b blockInfo, args map[string]interface{}) error {
	contractAddress, ok := args["contractAddress"]
	if !ok {
		return removeMemberEvent.checkArg("contractAddress")
	}

	memberAddress, ok := args["memberAddress"]
	if !ok {
		return removeMemberEvent.checkArg("memberAddress")
	}

	err := daoMemberService.Update(b.height, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex(), model.MemberStatusRemoved)
	if err != nil {
		return err
	}
	return nil
}

func (removeMemberEvent RemoveMemberEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao remove member event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
