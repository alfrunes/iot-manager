// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/mongo/driver/operation"
)

// BulkWriteResult is the result type returned by a BulkWrite operation.
type BulkWriteResult struct {
	// The number of documents inserted.
	InsertedCount int64

	// The number of documents matched by filters in update and replace operations.
	MatchedCount int64

	// The number of documents modified by update and replace operations.
	ModifiedCount int64

	// The number of documents deleted.
	DeletedCount int64

	// The number of documents upserted by update and replace operations.
	UpsertedCount int64

	// A map of operation index to the _id of each upserted document.
	UpsertedIDs map[int64]interface{}
}

// InsertOneResult is the result type returned by an InsertOne operation.
type InsertOneResult struct {
	// The _id of the inserted document. A value generated by the driver will be of type primitive.ObjectID.
	InsertedID interface{}
}

// InsertManyResult is a result type returned by an InsertMany operation.
type InsertManyResult struct {
	// The _id values of the inserted documents. Values generated by the driver will be of type primitive.ObjectID.
	InsertedIDs []interface{}
}

// DeleteResult is the result type returned by DeleteOne and DeleteMany operations.
type DeleteResult struct {
	DeletedCount int64 `bson:"n"` // The number of documents deleted.
}

// ListDatabasesResult is a result of a ListDatabases operation.
type ListDatabasesResult struct {
	// A slice containing one DatabaseSpecification for each database matched by the operation's filter.
	Databases []DatabaseSpecification

	// The total size of the database files of the returned databases in bytes.
	// This will be the sum of the SizeOnDisk field for each specification in Databases.
	TotalSize int64
}

func newListDatabasesResultFromOperation(res operation.ListDatabasesResult) ListDatabasesResult {
	var ldr ListDatabasesResult
	ldr.Databases = make([]DatabaseSpecification, 0, len(res.Databases))
	for _, spec := range res.Databases {
		ldr.Databases = append(
			ldr.Databases,
			DatabaseSpecification{Name: spec.Name, SizeOnDisk: spec.SizeOnDisk, Empty: spec.Empty},
		)
	}
	ldr.TotalSize = res.TotalSize
	return ldr
}

// DatabaseSpecification contains information for a database. This type is returned as part of ListDatabasesResult.
type DatabaseSpecification struct {
	Name       string // The name of the database.
	SizeOnDisk int64  // The total size of the database files on disk in bytes.
	Empty      bool   // Specfies whether or not the database is empty.
}

// UpdateResult is the result type returned from UpdateOne, UpdateMany, and ReplaceOne operations.
type UpdateResult struct {
	MatchedCount  int64       // The number of documents matched by the filter.
	ModifiedCount int64       // The number of documents modified by the operation.
	UpsertedCount int64       // The number of documents upserted by the operation.
	UpsertedID    interface{} // The _id field of the upserted document, or nil if no upsert was done.
}

// UnmarshalBSON implements the bson.Unmarshaler interface.
func (result *UpdateResult) UnmarshalBSON(b []byte) error {
	elems, err := bson.Raw(b).Elements()
	if err != nil {
		return err
	}

	for _, elem := range elems {
		switch elem.Key() {
		case "n":
			switch elem.Value().Type {
			case bson.TypeInt32:
				result.MatchedCount = int64(elem.Value().Int32())
			case bson.TypeInt64:
				result.MatchedCount = elem.Value().Int64()
			default:
				return fmt.Errorf("Received invalid type for n, should be Int32 or Int64, received %s", elem.Value().Type)
			}
		case "nModified":
			switch elem.Value().Type {
			case bson.TypeInt32:
				result.ModifiedCount = int64(elem.Value().Int32())
			case bson.TypeInt64:
				result.ModifiedCount = elem.Value().Int64()
			default:
				return fmt.Errorf("Received invalid type for nModified, should be Int32 or Int64, received %s", elem.Value().Type)
			}
		case "upserted":
			switch elem.Value().Type {
			case bson.TypeArray:
				e, err := elem.Value().Array().IndexErr(0)
				if err != nil {
					break
				}
				if e.Value().Type != bson.TypeEmbeddedDocument {
					break
				}
				var d struct {
					ID interface{} `bson:"_id"`
				}
				err = bson.Unmarshal(e.Value().Document(), &d)
				if err != nil {
					return err
				}
				result.UpsertedID = d.ID
			default:
				return fmt.Errorf("Received invalid type for upserted, should be Array, received %s", elem.Value().Type)
			}
		}
	}

	return nil
}

// IndexSpecification represents an index in a database. This type is returned by the IndexView.ListSpecifications
// function and is also used in the CollectionSpecification type.
type IndexSpecification struct {
	// The index name.
	Name string

	// The namespace for the index. This is a string in the format "databaseName.collectionName".
	Namespace string

	// The keys specification document for the index.
	KeysDocument bson.Raw

	// The index version.
	Version int32
}

var _ bson.Unmarshaler = (*IndexSpecification)(nil)

type unmarshalIndexSpecification struct {
	Name         string   `bson:"name"`
	Namespace    string   `bson:"ns"`
	KeysDocument bson.Raw `bson:"key"`
	Version      int32    `bson:"v"`
}

// UnmarshalBSON implements the bson.Unmarshaler interface.
func (i *IndexSpecification) UnmarshalBSON(data []byte) error {
	var temp unmarshalIndexSpecification
	if err := bson.Unmarshal(data, &temp); err != nil {
		return err
	}

	i.Name = temp.Name
	i.Namespace = temp.Namespace
	i.KeysDocument = temp.KeysDocument
	i.Version = temp.Version
	return nil
}

// CollectionSpecification represents a collection in a database. This type is returned by the
// Database.ListCollectionSpecifications function.
type CollectionSpecification struct {
	// The collection name.
	Name string

	// The type of the collection. This will either be "collection" or "view".
	Type string

	// Whether or not the collection is readOnly. This will be false for MongoDB versions < 3.4.
	ReadOnly bool

	// The collection UUID. This field will be nil for MongoDB versions < 3.6. For versions 3.6 and higher, this will
	// be a primitive.Binary with Subtype 4.
	UUID *primitive.Binary

	// A document containing the options used to construct the collection.
	Options bson.Raw

	// An IndexSpecification instance with details about the collection's _id index. This will be nil if the NameOnly
	// option is used and for MongoDB versions < 3.4.
	IDIndex *IndexSpecification
}

var _ bson.Unmarshaler = (*CollectionSpecification)(nil)

// unmarshalCollectionSpecification is used to unmarshal BSON bytes from a listCollections command into a
// CollectionSpecification.
type unmarshalCollectionSpecification struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
	Info *struct {
		ReadOnly bool              `bson:"readOnly"`
		UUID     *primitive.Binary `bson:"uuid"`
	} `bson:"info"`
	Options bson.Raw            `bson:"options"`
	IDIndex *IndexSpecification `bson:"idIndex"`
}

// UnmarshalBSON implements the bson.Unmarshaler interface.
func (cs *CollectionSpecification) UnmarshalBSON(data []byte) error {
	var temp unmarshalCollectionSpecification
	if err := bson.Unmarshal(data, &temp); err != nil {
		return err
	}

	cs.Name = temp.Name
	cs.Type = temp.Type
	if cs.Type == "" {
		// The "type" field is only present on 3.4+ because views were introduced in 3.4, so we implicitly set the
		// value to "collection" if it's empty.
		cs.Type = "collection"
	}
	if temp.Info != nil {
		cs.ReadOnly = temp.Info.ReadOnly
		cs.UUID = temp.Info.UUID
	}
	cs.Options = temp.Options
	cs.IDIndex = temp.IDIndex
	return nil
}
