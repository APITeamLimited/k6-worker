// Copyright (C) MongoDB, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package internal

import (
	"fmt"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

const (
	EncryptedCacheCollection      = "ecc"
	EncryptedStateCollection      = "esc"
	EncryptedCompactionCollection = "ecoc"
)

// GetEncryptedStateCollectionName returns the encrypted state collection name associated with dataCollectionName.
func GetEncryptedStateCollectionName(efBSON bsoncore.Document, dataCollectionName string, stateCollection string) (string, error) ***REMOVED***
	fieldName := stateCollection + "Collection"
	val, err := efBSON.LookupErr(fieldName)
	if err != nil ***REMOVED***
		if err != bsoncore.ErrElementNotFound ***REMOVED***
			return "", err
		***REMOVED***
		// Return default name.
		defaultName := "enxcol_." + dataCollectionName + "." + stateCollection
		return defaultName, nil
	***REMOVED***

	stateCollectionName, ok := val.StringValueOK()
	if !ok ***REMOVED***
		return "", fmt.Errorf("expected string for '%v', got: %v", fieldName, val.Type)
	***REMOVED***
	return stateCollectionName, nil
***REMOVED***
