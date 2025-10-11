package userRepository

import (
	"context"
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Search(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, id string, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (ref *userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	record := model.UserFromDomain(user)

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	created, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id := created.InsertedID.(primitive.ObjectID)

	result := ref.collection.FindOne(ctx, bson.M{"_id": id})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.User
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}

func (ref *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var record model.User
	if err = result.Decode(&record); err != nil {
		return nil, err
	}

	return record.ToDomain(), nil
}

func (ref *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := ref.collection.FindOne(ctx, bson.M{"email": email})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var record model.User
	if err := result.Decode(&record); err != nil {
		return nil, err
	}

	return record.ToDomain(), nil
}

func (ref *userRepository) Search(ctx context.Context) ([]entity.User, error) {
	cursor, err := ref.collection.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNilCursor {
			return nil, nil
		}
		return nil, err
	}

	records := make([]entity.User, 0)

	for cursor.Next(ctx) {
		var record model.User
		if err = cursor.Decode(&record); err != nil {
			return nil, err
		}

		records = append(records, *record.ToDomain())
	}

	return records, nil
}

func (ref *userRepository) Update(ctx context.Context, id string, user entity.User) (*entity.User, error) {
	record := model.UserFromDomain(user)
	record.UpdatedAt = time.Now()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": record,
	}

	_, err = ref.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.User
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}

func (ref *userRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = ref.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
