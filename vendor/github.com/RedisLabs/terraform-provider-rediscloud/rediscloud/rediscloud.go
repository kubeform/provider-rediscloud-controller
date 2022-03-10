package rediscloud

import (
	"github.com/RedisLabs/terraform-provider-rediscloud/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New(version string) func() *schema.Provider {
	return provider.New(version)
}
