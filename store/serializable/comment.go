package serializable

import "time"

type serializableComment struct {
	SerializedAuthor    string
	SerializedBody      string
	SerializedCreatedAt time.Time
	SerializedUpdatedAt time.Time
}

func (c *serializableComment) Author() string       { return c.SerializedAuthor }
func (c *serializableComment) Body() string         { return c.SerializedBody }
func (c *serializableComment) CreatedAt() time.Time { return c.SerializedCreatedAt }
func (c *serializableComment) UpdatedAt() time.Time { return c.UpdatedAt() }
