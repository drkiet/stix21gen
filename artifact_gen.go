package main

import (

	stix21 "github.com/drkiet/stix21"
)

// Artifact
/**
Type string                             `json:"type" binding:"required"`
SpecVersion string                      `json:"spec_version,omitempty"`
ID string                               `json:"id" binding:"required"`
ObjectMarkingRefs []string              `json:"object_marking_refs,omitempty"`
GranularMarkings []GranularMarking `json:"granular-markings,omitempty"`
Defanged bool                           `json:"defanged,omitempty"`
Extensions map[string]interface{}       `json:"extensions,omitempty"`

MimeType string `json:"mime_type,omitempty"`
PayloadBin string `json:"payload_bin,omitempty"` // binary type
Url string `json:"url,omitempty"`
Hashes map[string]interface{} `json:"hashes,omitempty"`
EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"` // enum
DecryptionAlgorithm string `json:"decryption_key,omitempty"`
 */
func copyHeader(artifact *stix21.Artifact, sdoHeader stix21.STIXObject) {
	artifact.Type = sdoHeader.Type
	artifact.ID = sdoHeader.ID
	artifact.SpecVersion = sdoHeader.SpecVersion
	artifact.ObjectMarkingRefs = sdoHeader.ObjectMarkingRefs
	artifact.GranularMarkings = sdoHeader.GranularMarkings
	artifact.Defanged = sdoHeader.Defanged
	artifact.Extensions = sdoHeader.Extensions
}

func GenerateArtifact() (artifact stix21.Artifact) {
	//artifact.Type = stix21.ArtifactType
	//artifact.ID = stix21.MakeIdentifier("artifact")
	//artifact.SpecVersion = "2.1"
	sdoHeader := stix21.STIXObject{}
	MakeSdoHeader(&sdoHeader, stix21.ArtifactType)
	copyHeader(&artifact, sdoHeader)
	artifact.MimeType = "application/zip"
	artifact.DecryptionAlgorithm = "decryption key value"
	artifact.EncryptionAlgorithm = stix21.AES_256_GCM
	return artifact
}
