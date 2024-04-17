package repositories

import (
	"context"
	. "go-fiber-proton/domain/datasources"
	"go-fiber-proton/domain/entities"
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type usersRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IUsersRepository interface {
	InsertNewUser(data *entities.NewUserBody) bool
	FindAll() (*[]entities.UserDataFormat, error)
	UpdateUser(data *entities.NewUserBody) bool
	DeleteUser(name string) bool
	FindUser(name string) bool
}

func NewUsersRepository(db *MongoDB) IUsersRepository {
	return &usersRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("users"),
	}
}

func (repo *usersRepository) InsertNewUser(data *entities.NewUserBody) bool {
	go func() {
		if _, err := repo.Collection.InsertOne(repo.Context, data); err != nil {
			fiberlog.Errorf("Users -> InsertNewUser: %s \n", err)
		}

	}()
	return true
}

func (repo *usersRepository) FindAll() (*[]entities.UserDataFormat, error) {
	options := options.Find()
	filter := bson.M{}
	cursor, err := repo.Collection.Find(repo.Context, filter, options)
	if err != nil {
		fiberlog.Errorf("Users -> FindAll: %s \n", err)
		return nil, err
	}
	defer cursor.Close(repo.Context)
	users := []entities.UserDataFormat{}
	for cursor.Next(repo.Context) {
		var item entities.UserDataFormat

		err := cursor.Decode(&item)
		if err != nil {
			continue
		}

		users = append(users, item)
	}
	return &users, nil
}

func (repo *usersRepository) UpdateUser(data *entities.NewUserBody) bool {
	filter := bson.M{"name": data.Name}
	dataTmp := bson.M{}
	if data.Age != 0 {
		dataTmp["age"] = data.Age
	}
	if data.Text != "" {
		dataTmp["text"] = data.Text
	}
	_, err := repo.Collection.UpdateOne(repo.Context, filter, bson.M{"$set": dataTmp})
	if err != nil {
		fiberlog.Errorf("Users -> UpdateText: %s \n", err)
		return false
	}
	return true

}

func (repo *usersRepository) DeleteUser(name string) bool {
	filter := bson.M{"name": name}
	_, err := repo.Collection.DeleteOne(repo.Context, filter)
	if err != nil {
		fiberlog.Errorf("Users -> DeleteUser: %s \n", err)
		return false
	}
	return true
}

func (repo *usersRepository) FindUser(name string) bool {
	filter := bson.M{"name": name}
	data := entities.UserDataFormat{}
	if err := repo.Collection.FindOne(repo.Context, filter).Decode(&data); err != nil {
		fiberlog.Errorf("Users -> FindUser: %s \n", err)
		return false
	}
	return true
}
