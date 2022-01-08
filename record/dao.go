package record

import (
	"context"
	"getir-assignment/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// Dao manages access of the database.
// it is used by a Service to construct a view model.
type Dao struct{
	Db *mongodb.Connection
}

// Find fetches the records filtering them according to tha values specified with FilterOptions.
func (d Dao) Find(options FilterOptions) ([]Dto, error) {
	var records []Dto

	// go to the database and fetch the collection of records.
	db := d.Db.Database(d.Db.DatabaseName())
	collection := db.Collection("records")

	// firstly, counts array should be summed up, and do not need all fields.
	// selecting the fields needed by using $project.
	// by using $sum, totalCount is calculated.
	// by $match, the required filtering options are applied.
	cursor, err := collection.Aggregate(context.Background(), []bson.M{
		{
			"$project": bson.M{
				"key": 3,
				"createdAt": 2,
				"totalCount": bson.M{
					"$sum": "$counts",
				},
			},
		},
		{
			"$match": bson.M{
				"createdAt": bson.M{ "$gte": options.StartDate, "$lte": options.EndDate },
				"totalCount": bson.M{ "$gte": options.MinCount, "$lte": options.MaxCount },
			},
		},
	})
	if err != nil {
		log.Printf("Error on finding in collection: %v", err)
		return records, err
	}

	var bsonRecords []bson.M
	err = cursor.All(context.Background(), &bsonRecords)
	if err != nil {
		log.Printf("error on iterating over the cursor: %v", err)
		return records, err
	}

	// iterating through the bson objects to process all of them.
	for _, rec := range bsonRecords {
		bsonBytes, err := bson.Marshal(rec)
		if err != nil {
			log.Printf("error on marshalling BSON: %v", err)
			return records, err
		}

		// for unmarshal bson objects easily, using Entity struct to represent it in the code.
		var entity Entity
		err = bson.Unmarshal(bsonBytes, &entity)
		if err != nil {
			log.Printf("error on unmarshalling BSON: %v", err)
			return records, err
		}

		// converting the entity model to a data transfer object.
		dto := Dto{
			Key:        entity.Key,
			CreatedAt:  entity.CreatedAt,
			TotalCount: entity.TotalCount,
		}
		records = append(records, dto)
	}

	return records, nil
}