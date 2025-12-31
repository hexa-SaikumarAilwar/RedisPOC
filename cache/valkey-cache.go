// cache/valkey-cache.go
package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hexa-SaikumarAilwar/RedisPOC.git/entity"
	"github.com/valkey-io/valkey-go"
)

type valkeyCache struct {
	client  valkey.Client
	expires time.Duration
}

func NewValkeyCache(addr string, expires time.Duration) PostCache {
	client, _ := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{addr},
	})
	return &valkeyCache{client: client, expires: expires}
}

func (v *valkeyCache) Get(key string) *entity.Post {
	resp := v.client.Do(
		context.Background(),
		v.client.B().Get().Key(key).Build(),
	)

	if resp.Error() != nil {
		return nil
	}

	str, err := resp.ToString()
	if err != nil {
		return nil
	}

	var post entity.Post
	if err := json.Unmarshal([]byte(str), &post); err != nil {
		return nil
	}

	return &post
}

func (v *valkeyCache) Set(key string, value *entity.Post) {
	bytes, _ := json.Marshal(value)
	v.client.Do(context.Background(),
		v.client.B().Set().Key(key).Value(string(bytes)).Ex(v.expires).Build(),
	)
}
