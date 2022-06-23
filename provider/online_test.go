// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package provider

import (
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/google/uuid"
	"os"
	"reflect"
	"testing"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return s
}

type OnlineResource struct {
	Entity string
	Value  interface{}
	Type   ValueType
}

func TestOnlineStores(t *testing.T) {
	testFns := map[string]func(*testing.T, OnlineStore){
		"CreateGetTable":     testCreateGetTable,
		"TableAlreadyExists": testTableAlreadyExists,
		"TableNotFound":      testTableNotFound,
		"SetGetEntity":       testSetGetEntity,
		"EntityNotFound":     testEntityNotFound,
		"TypeCasting":        testTypeCasting,
	}

	//Redis (Mock)
	// miniRedis := mockRedis()
	// defer miniRedis.Close()
	// mockRedisAddr := miniRedis.Addr()
	// redisMockConfig := &RedisConfig{
	// 	Addr: mockRedisAddr,
	// }

	// //Redis (Live)
	// redisPort := os.Getenv("REDIS_PORT")
	// liveAddr := fmt.Sprintf("%s:%s", "localhost", redisPort)
	// redisLiveConfig := &RedisConfig{
	// 	Addr: liveAddr,
	// }

	dynamoPort, ok := os.LookupEnv("DYNAMO_PORT")
	dynamoAccessKey := os.Getenv("DYNAMO_ACCESS_KEY")
	dynamoSecretKey := os.Getenv("DYNAMO_SECRET_KEY")
	if !ok {
		dynamoPort = "8000"
		dynamoAccessKey = "AKIAYUXQ37DFTRWBOL7J"
		dynamoSecretKey = "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"
	} 
	dynamoAddr := fmt.Sprintf("%s:%s", "localhost", dynamoPort)
	dynamoConfig := &DynamodbConfig{
		Addr:      dynamoAddr,
		Region:    "us-east-1",
		AccessKey: dynamoAccessKey,
		SecretKey: dynamoSecretKey,
	}

	testList := []struct {
		t               Type
		c               SerializedConfig
		integrationTest bool
	}{
		{LocalOnline, []byte{}, false},
		// {RedisOnline, redisMockConfig.Serialized(), false},
		// {RedisOnline, redisLiveConfig.Serialized(), true},
		{DynamoDBOnline, dynamoConfig.Serialized(), true},
	}
	for _, testItem := range testList {
		if testing.Short() && testItem.integrationTest {
			t.Logf("Skipping %s, because it is an integration test", testItem.t)
			continue
		}
		for name, fn := range testFns {
			provider, err := Get(testItem.t, testItem.c)
			if err != nil {
				t.Fatalf("Failed to get provider %s: %s", testItem.t, err)
			}
			store, err := provider.AsOnlineStore()
			if err != nil {
				t.Fatalf("Failed to use provider %s as OnlineStore: %s", testItem.t, err)
			}
			var prefix string
			if testItem.integrationTest {
				prefix = "INTEGRATION"
			} else {
				prefix = "UNIT"
			}
			testName := fmt.Sprintf("%s_%s_%s", testItem.t, prefix, name)
			t.Run(testName, func(t *testing.T) {
				fn(t, store)
			})
		}
	}
}

func randomFeatureVariant() (string, string) {
	return uuid.NewString(), uuid.NewString()
}

func testCreateGetTable(t *testing.T, store OnlineStore) {
	mockFeature, mockVariant := randomFeatureVariant()
	if store.Type() == "DYNAMODB_ONLINE" {
		key := dynamodbTableKey{mockFeature, mockVariant, String}
		defer dynamoDeleteTable(key.String())
	}
	if tab, err := store.CreateTable(mockFeature, mockVariant, String); tab == nil || err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}
	if tab, err := store.GetTable(mockFeature, mockVariant); tab == nil || err != nil {
		t.Fatalf("Failed to get table: %s", err)
	}
}

func testTableAlreadyExists(t *testing.T, store OnlineStore) {
	mockFeature, mockVariant := randomFeatureVariant()
	if _, err := store.CreateTable(mockFeature, mockVariant, String); err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}
	if _, err := store.CreateTable(mockFeature, mockVariant, String); err == nil {
		t.Fatalf("Succeeded in creating table twice")
	} else if casted, valid := err.(*TableAlreadyExists); !valid {
		t.Fatalf("Wrong error for table already exists: %T", err)
	} else if casted.Error() == "" {
		t.Fatalf("TableAlreadyExists has empty error message")
	}
}

func testTableNotFound(t *testing.T, store OnlineStore) {
	mockFeature, mockVariant := randomFeatureVariant()
	if _, err := store.GetTable(mockFeature, mockVariant); err == nil {
		t.Fatalf("Succeeded in getting non-existant table")
	} else if casted, valid := err.(*TableNotFound); !valid {
		t.Fatalf("Wrong error for table not found: %s,%T", err, err)
	} else if casted.Error() == "" {
		t.Fatalf("TableNotFound has empty error message")
	}
}

func testSetGetEntity(t *testing.T, store OnlineStore) {
	mockFeature, mockVariant := randomFeatureVariant()
	entity, val := "e", "val"
	tab, err := store.CreateTable(mockFeature, mockVariant, String)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}
	if err := tab.Set(entity, val); err != nil {
		t.Fatalf("Failed to set entity: %s", err)
	}
	gotVal, err := tab.Get(entity)
	if err != nil {
		t.Fatalf("Failed to get entity: %s", err)
	}
	if !reflect.DeepEqual(val, gotVal) {
		t.Fatalf("Values are not the same %v %v", val, gotVal)
	}
}

func testEntityNotFound(t *testing.T, store OnlineStore) {
	mockFeature, mockVariant := uuid.NewString(), "v"
	entity := "e"
	tab, err := store.CreateTable(mockFeature, mockVariant, String)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}
	if _, err := tab.Get(entity); err == nil {
		t.Fatalf("succeeded in getting non-existent entity")
	} else if casted, valid := err.(*EntityNotFound); !valid {
		t.Fatalf("Wrong error for entity not found: %T", err)
	} else if casted.Error() == "" {
		t.Fatalf("EntityNotFound has empty error message")
	}
}

func testTypeCasting(t *testing.T, store OnlineStore) {
	onlineResources := []OnlineResource{
		{
			Entity: "a",
			Value:  int(1),
			Type:   Int,
		},
		{
			Entity: "b",
			Value:  int64(1),
			Type:   Int64,
		},
		{
			Entity: "c",
			Value:  float32(1.0),
			Type:   Float32,
		},
		{
			Entity: "d",
			Value:  float64(1.0),
			Type:   Float64,
		},
		{
			Entity: "e",
			Value:  "1.0",
			Type:   String,
		},
		{
			Entity: "f",
			Value:  false,
			Type:   Bool,
		},
	}
	for _, resource := range onlineResources {
		featureName := uuid.New().String()
		tab, err := store.CreateTable(featureName, "", resource.Type)
		if err != nil {
			t.Fatalf("Failed to create table: %s", err)
		}
		if err := tab.Set(resource.Entity, resource.Value); err != nil {
			t.Fatalf("Failed to set entity: %s", err)
		}
		gotVal, err := tab.Get(resource.Entity)
		if err != nil {
			t.Fatalf("Failed to get entity: %s", err)
		}
		if !reflect.DeepEqual(resource.Value, gotVal) {
			t.Fatalf("Values are not the same %v, type %T. %v, type %T", resource.Value, resource.Value, gotVal, gotVal)
		}
	}
}

func dynamoDeleteTable(tablename string) error{
	dynamoPort := os.Getenv("DYNAMO_PORT")
	dynamoAccessKey := os.Getenv("DYNAMO_ACCESS_KEY")
	dynamoSecretKey := os.Getenv("DYNAMO_SECRET_KEY")
	dynamoAddr := fmt.Sprintf("%s:%s", "localhost", dynamoPort)
	dynamoConfig := &DynamodbConfig{
		Addr:      dynamoAddr,
		Region:    "us-east-1",
		AccessKey: dynamoAccessKey,
		SecretKey: dynamoSecretKey,
	}
	sess := session.Must(session.NewSession(getDynamodbConfig(dynamoConfig)))
	dynamodbClient := dynamodb.New(sess)
	params := &dynamodb.DeleteTableInput{
		TableName: aws.String(tablename), 
	}
	_, err := dynamodbClient.DeleteTable(params)
	if err != nil {
		return err
	}
	return nil
}
