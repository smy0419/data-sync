package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type ConfirmNewPresident struct{}

func (confirmNewPresident ConfirmNewPresident) handle(b blockInfo, args map[string]interface{}) error {
	contractAddress, ok := args["contractAddress"]
	if !ok {
		return confirmNewPresident.checkArg("contractAddress")
	}

	memberAddress, ok := args["memberAddress"]
	if !ok {
		return confirmNewPresident.checkArg("memberAddress")
	}

	err := daoMemberService.RemovePresident(b.height, contractAddress.(*asimovCommon.Address).Hex())
	if err != nil {
		return err
	}

	err = daoMemberService.UpdatePresident(b.height, b.time, b.txHash, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex())
	if err != nil {
		return err
	}

	err = daoOrganizationService.UpdatePresident(b.height, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex())

	err = daoTodoListService.Release(b.height, contractAddress.(*asimovCommon.Address).Hex(), memberAddress.(*asimovCommon.Address).Hex(), model.TodoTypeInvitePresident)
	if err != nil {
		return err
	}

	return nil
}

func (confirmNewPresident ConfirmNewPresident) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao confirm new president event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
