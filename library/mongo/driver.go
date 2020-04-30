package mongo

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/response"
	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

func FindOne(collection string, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	err := MongoDB.Collection(collection).FindOne(context.TODO(), filter, opts...).Decode(result)
	if err != nil {
		// common.Logger.Errorf("find one result from %s error. err: %s", collection, err)
		if err.Error() == "mongo: no documents in result" {
			err = response.NewDataNoExistError()
		}
		return err
	}
	return nil
}

func Find(collection string, filter interface{}, t reflect.Type, tPointer reflect.Type, opts ...*options.FindOptions) (interface{}, error) {
	results := reflect.MakeSlice(reflect.SliceOf(tPointer), 0, 0)
	cursor, err := MongoDB.Collection(collection).Find(context.TODO(), filter, opts...)

	defer cursor.Close(context.TODO())
	if err != nil {
		common.Logger.Errorf("find results from %s error. err: %s", collection, err)
		return nil, err
	}
	if err := cursor.Err(); err != nil {
		common.Logger.Errorf("find results from %s error. err: %s", collection, err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		elem := bson.M{}
		err := cursor.Decode(&elem)
		if err != nil {
			common.Logger.Errorf("find results from %s error. err: %s", collection, err)
			return nil, err
		}

		structElem := reflect.New(t).Interface()
		// structElem := reflect.New(t).Elem().Interface()  value or pointer
		err = bson2Odj(elem, structElem)
		if err != nil {
			common.Logger.Errorf("bson to obj failed, err: %s", err)
			return nil, err
		}
		results = reflect.Append(results, reflect.ValueOf(structElem))
	}

	return results.Interface(), nil
}

func QueryByPage(collection string, filter interface{}, findOptions *options.FindOptions, t reflect.Type, tPointer reflect.Type) (int64, interface{}, error) {
	results := reflect.MakeSlice(reflect.SliceOf(tPointer), 0, 0)
	total, err := MongoDB.Collection(collection).CountDocuments(context.TODO(), filter)
	if err != nil {
		common.Logger.Errorf("query %s error. err: %s", collection, err)
		return 0, nil, err
	}
	if total == 0 {
		return 0, results.Interface(), nil
	}

	cursor, err := MongoDB.Collection(collection).Find(context.TODO(), filter, findOptions)
	defer cursor.Close(context.TODO())
	if err != nil {
		common.Logger.Errorf("query %s error. err: %s", collection, err)
		return 0, nil, err
	}
	if err := cursor.Err(); err != nil {
		common.Logger.Errorf("query %s error. err: %s", collection, err)
		return 0, nil, err
	}

	for cursor.Next(context.TODO()) {
		elem := bson.M{}
		err := cursor.Decode(&elem)
		if err != nil {
			common.Logger.Errorf("query %s error. err: %s", collection, err)
			return 0, nil, err
		}

		structElem := reflect.New(t).Interface()
		err = bson2Odj(elem, structElem)
		if err != nil {
			common.Logger.Errorf("bson to obj failed, err: %s", err)
			return 0, nil, err
		}
		results = reflect.Append(results, reflect.ValueOf(structElem))
	}

	return total, results.Interface(), nil
}

func Exist(collection string, filter interface{}) (bool, error) {
	count, err := MongoDB.Collection(collection).CountDocuments(context.TODO(), filter)
	if err != nil {
		common.Logger.Errorf("call %s exist error. err: %s", collection, err)
		return false, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func Count(collection string, filter interface{}) (int64, error) {
	count, err := MongoDB.Collection(collection).CountDocuments(context.TODO(), filter)
	if err != nil {
		common.Logger.Errorf("call %s count error. err: %s", collection, err)
		return 0, err
	}
	return count, nil
}

func Sum(collection string, condition []byte, t reflect.Type) (interface{}, error) {
	pipeLine := mongoDriver.Pipeline{}
	// err := bson.UnmarshalExtJSON([]byte(condition), true, &pipeLine)
	err := bson.UnmarshalExtJSON(condition, true, &pipeLine)
	if err != nil {
		return reflect.Zero(t).Interface(), err
	}

	cursor, err := MongoDB.Collection(collection).Aggregate(context.TODO(), pipeLine)
	if err != nil {
		return reflect.Zero(t).Interface(), err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result bson.M
		err = cursor.Decode(&result)
		if err != nil {
			return reflect.Zero(t).Interface(), err
		}
		return result["total"], nil
	}

	return reflect.Zero(t).Interface(), nil
}

func bson2Odj(val interface{}, obj interface{}) error {
	data, err := bson.Marshal(val)
	if err != nil {
		return err
	}
	return bson.Unmarshal(data, obj)
}
