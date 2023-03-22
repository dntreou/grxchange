package domain

type CampaignName string
type Targeting string
type Brand string
type CampaignType string
type Event string

// Campaign represents the domain level entity, the first field is being passed as is to ensure that the entity is not used
type Campaign struct {
	_            interface{}
	name         CampaignName
	targeting    Targeting
	brand        Brand
	campaignType CampaignType
	event        Event
}

func NewCampaign(
	name CampaignName,
	targeting Targeting,
	brand Brand,
	campaignType CampaignType,
	event Event,
) (*Campaign, error) {
	// TODO add validation of each field, this is an operation that should happend at the domain level
	return &Campaign{
		name:         name,
		targeting:    targeting,
		brand:        brand,
		campaignType: campaignType,
		event:        event,
	}, nil
}
