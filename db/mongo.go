package db

import (
	"context"
	"time"

	"github.com/MxHonesty/change4backend/auth"
	"github.com/MxHonesty/change4backend/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	client *mongo.Client
	Users  *mongo.Collection
}

func NewMongodb() *Mongodb {
	logging.InfoLogger.Println("Starting Mongo connection")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://root:root@change4.j4oct.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logging.ErrorLogger.Fatal(err)
	}

	return &Mongodb{client: client,
		Users: client.Database("Change").Collection("User"),
	}
}

func (m *Mongodb) CloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logging.InfoLogger.Println("Closing Mongo Connection")
	if err := m.client.Disconnect(ctx); err != nil {
		logging.ErrorLogger.Fatal(err)
	}
}

func (m *Mongodb) FindAllUsers() []auth.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	curr, currErr := m.Users.Find(ctx, bson.D{})
	if currErr != nil {
		logging.ErrorLogger.Fatal(currErr)
	}

	var users []auth.User
	if err := curr.All(ctx, &users); err != nil {
		panic(err)
	}
	return users
}

func (m *Mongodb) FindOneUser(username, password string) (auth.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundUser auth.User
	res := m.Users.FindOne(ctx, bson.D{{Key: "userName", Value: username}, {Key: "password", Value: password}})
	err := res.Decode(&foundUser)
	return foundUser, err
}

func (m *Mongodb) AddUser(username, password string, userType uint8) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	passHash, _ := auth.HashPassword(password)
	add := bson.D{
		{Key: "userName", Value: username},
		{Key: "password", Value: passHash},
		{Key: "userType", Value: userType},
	}

	res, err := m.Users.InsertOne(ctx, add)
	return res.InsertedID.(primitive.ObjectID).String(), err
}
