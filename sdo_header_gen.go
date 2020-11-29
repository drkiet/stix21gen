package main

import "github.com/drkiet/stix21"

const DEFAULT_VERSION = "2.1"

/**
Type string                             `json:"type" binding:"required"`
SpecVersion string                      `json:"spec_version,omitempty"`
ID string                               `json:"id" binding:"required"`
ObjectMarkingRefs []string              `json:"object_marking_refs,omitempty"`
GranularMarkings []GranularMarking `json:"granular-markings,omitempty"`
Defanged bool                           `json:"defanged,omitempty"`
Extensions map[string]interface{}       `json:"extensions,omitempty"`
 */
func MakeSdoHeader (object *stix21.STIXObject, sdoType string) {
	object.Type = sdoType
	object.SpecVersion = DEFAULT_VERSION
	object.ID = stix21.MakeIdentifier(sdoType)
}
