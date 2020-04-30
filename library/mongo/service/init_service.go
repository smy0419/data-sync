package service

/**
Execute when the block height is 0
*/
type HeightInitService interface {
	HeightInit() error
}

/**
Execute when project start or rollback
*/
type ProjectStartupService interface {
	ProjectStart(height int32) error
}

// var heightInitHandler = []HeightInitService{AssetService{}}

// var projectStartupHandler []ProjectStartupService
// var projectStartupHandler = []ProjectStartupService{FoundationRoleActionService{}}

type InitService struct {
	HeightInitHandler     []HeightInitService
	ProjectStartupHandler []ProjectStartupService
}

/**
Height is next block to synchronize
*/
func (initService InitService) Init(height int32) error {
	if height == 0 {
		for _, handler := range initService.HeightInitHandler {
			err := handler.HeightInit()
			if err != nil {
				return err
			}

		}
	}

	for _, handler := range initService.ProjectStartupHandler {
		err := handler.ProjectStart(height - 1)
		if err != nil {
			return err
		}
	}

	return nil
}
