package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"fitnessProgramApi/pkg/config"
)

// User represents a user in the system.
type User struct {
    Username             string  // Exported field
    BodyWeightKG         float64 // Exported field
    HeightCM             float64 // Exported field
    Age                  int     // Exported field
    ProjectedBodyWeightKG float64 // Exported field
}

// SetNewUser creates a new User instance with given parameters.
func SetNewUser(ctx context.Context, username string, bodyWeightKG, heightCM float64, age int, projectedBodyWeightKG float64) error {
    client, err := config.GetConnection()
    if err != nil {
        return err
    }
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            log.Println("Error disconnecting MongoDB client:", err)
        }
    }()

    collection := client.Database("FitnessApi").Collection("Users")
    user := User{
        Username:             username,
        BodyWeightKG:         bodyWeightKG,
        HeightCM:             heightCM,
        Age:                  age,
        ProjectedBodyWeightKG: projectedBodyWeightKG,
    }

    _, err = collection.InsertOne(ctx, user)
    if err != nil {
        return err
    }
    return nil
}

// UpdateUser updates the fields of an existing User.
func (u *User) UpdateUser(ctx context.Context, username string, bodyWeightKG, heightCM float64, age int, projectedBodyWeightKG float64) error {
    u.Username = username
    u.BodyWeightKG = bodyWeightKG
    u.HeightCM = heightCM
    u.Age = age
    u.ProjectedBodyWeightKG = projectedBodyWeightKG

    client, err := config.GetConnection()
    if err != nil {
        return err
    }
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            log.Println("Error disconnecting MongoDB client:", err)
        }
    }()

    collection := client.Database("FitnessApi").Collection("Users")
    _, err = collection.UpdateOne(ctx, bson.M{"_id": u.Username}, bson.M{"$set": u})
	if err != nil {
        return err
    }
    return nil
}

// DeleteUser deletes a user from the database.
func (u *User) DeleteUser(ctx context.Context, userID string) error {
    client, err := config.GetConnection()
    if err != nil {
        return err
    }
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            log.Println("Error disconnecting MongoDB client:", err)
        }
    }()

    collection := client.Database("FitnessApi").Collection("Users")
    _, err = collection.DeleteOne(ctx, bson.M{"_id": userID})
    if err != nil {
        return err
    }
    return nil
}