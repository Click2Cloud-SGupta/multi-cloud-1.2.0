/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateVolumeTagsResponse struct {
}

func (o BatchCreateVolumeTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateVolumeTagsResponse", string(data)}, " ")
}
