package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
)

type ImpeachCitizenEvent struct{}

func (impeachCitizenEvent ImpeachCitizenEvent) handle(b blockInfo, args map[string]interface{}) error {
	member, ok := args["oldCitizen"]
	if !ok {
		return impeachCitizenEvent.checkArg("oldCitizen")
	}
	return foundationMemberService.OutService(b.height, b.time, member.(*asimovCommon.Address).Hex())
}

func (impeachCitizenEvent ImpeachCitizenEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("impeach citizen event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
