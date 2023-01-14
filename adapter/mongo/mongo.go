package mongo

import (
	"fmt"
	"micro_services/msbase/db/mongo"
	"micro_services/msbase/logger"

	"go.mongodb.org/mongo-driver/bson"
)

func Create(collection string, model interface{}) error {
	var (
		err error
	)

	if err = mongo.DBClient.Create(collection, model); err != nil {
		logger.Error(fmt.Sprintf("Document insert unsuccessfully to %s, error : ", err.Error()))
		return err
	}

	logger.Info("Document insert successfully to : ", collection)

	return err
}

func Read(collection string, q bson.M, result interface{}, params ...interface{}) error {
	var (
		err error
	)

	if err = mongo.DBClient.Read(collection, q, result, params...); err != nil {
		logger.Error(fmt.Sprintf("Document read unsuccessfully from %s, error : ", err.Error()))
		return err
	}

	logger.Info("Document read successfully from : ", collection)

	return err
}

func ReadOne(collection string, q bson.M, result interface{}) error {
	var (
		err error
	)

	if err = mongo.DBClient.ReadOne(collection, q, result); err != nil {
		logger.Error(fmt.Sprintf("Document read one unsuccessfully from %s, error : ", err.Error()))
		return err
	}

	logger.Info("Document read one successfully from : ", collection)

	return err
}

func UpdateOne(collection string, filter bson.M, update bson.M) error {
	var (
		err error
	)

	if err = mongo.DBClient.UpdateOne(collection, filter, update); err != nil {
		logger.Error(fmt.Sprintf("Document update one unsuccessfully on %s, error : ", err.Error()))
		return err
	}

	logger.Info("Document update one successfully on : ", collection)

	return err
}

func UpdateMany(collection string, filter bson.M, update bson.M) error {
	var (
		err error
	)

	if err = mongo.DBClient.UpdateMany(collection, filter, update); err != nil {
		logger.Error(fmt.Sprintf("Document update many unsuccessfully on %s, error : ", err.Error()))
		return err
	}

	logger.Info("Document update many successfully on : ", collection)

	return err
}

func DeleteOne(collection string, filter bson.M) error {
	var (
		err error
	)

	if err = mongo.DBClient.DeleteOne(collection, filter); err != nil {
		logger.Error(fmt.Sprintf("Document delete one unsuccessfully from %s, error : ", err.Error()))
		return err
	}

	logger.Info("Document delete one successfully from : ", collection)

	return err
}

func DeleteMany(collection string, filter bson.M) error {
	var (
		err error
	)

	if err = mongo.DBClient.DeleteMany(collection, filter); err != nil {
		logger.Error(fmt.Sprintf("Document delete many unsuccessfully from %s, error : ", err.Error()))
		return err
	}

	logger.Info("Document delete many successfully from : ", collection)

	return err
}

func Count(collection string, filter bson.M) (int, error) {
	var (
		err   error
		count int
	)

	if count, err = mongo.DBClient.Count(collection, filter); err != nil {
		logger.Error(fmt.Sprintf("Get documents count unsuccessfully from %s, error : ", err.Error()))
		return 0, err
	}

	logger.Info("Get document count successfully from : ", collection)

	return count, err
}
