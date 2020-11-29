package main

import (
	"encoding/json"
	"github.com/drkiet/stix21"
)

func GenerateBundle(object json.RawMessage) (bundle stix21.Bundle){
	bundle = stix21.Bundle{}
	bundle.Type = "bundle"
	bundle.ID = stix21.MakeIdentifier(bundle.Type)
	bundle.Objects = make([]json.RawMessage, 1,1)
	bundle.Objects[0] = object
	return bundle
}