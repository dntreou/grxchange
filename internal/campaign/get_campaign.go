package campaign

import (
	"context"

	"grxchange/internal/domain"
)

type GetCampaign struct {
	campaigns interface{} // TODO replace with persistence layer implementation, now a fake interface is being used
}

func (g GetCampaign) GetCampaign(ctx context.Context) (domain.Campaign, error) {

	return domain.Campaign{}, nil
}
