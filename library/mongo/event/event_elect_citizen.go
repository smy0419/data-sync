package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
)

type ElectCitizenEvent struct{}

func (electCitizenEvent ElectCitizenEvent) handle(b blockInfo, args map[string]interface{}) error {
	member, ok := args["newCitizen"]
	if !ok {
		return electCitizenEvent.checkArg("newCitizen")
	}
	return foundationMemberService.Insert(b.height, b.time, []string{member.(*asimovCommon.Address).Hex()})
}

func (electCitizenEvent ElectCitizenEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("elect citizen event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
