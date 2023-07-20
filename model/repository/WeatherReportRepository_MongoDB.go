package repository

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
	"workspace/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ==================================================================================================
// DBConfig struct
// ==================================================================================================

// -----------------
// properties,fields
// -----------------
type DBConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

// ==================================================================================================
//  WeatherReportRepository_MongoDB struct implements WeatherReportRepository interface
// ==================================================================================================

// -----------------
// properties,fields
// -----------------
type WeatherReportRepository_MongoDB struct {
	DBConfig DBConfig
	Client   *mongo.Client
}

// -----------------
// constructors
// -----------------
func NewWeatherReportRepository_MongoDB() (*WeatherReportRepository_MongoDB, error) {

	repo := &WeatherReportRepository_MongoDB{}

	err := repo.initDBConfig()
	if err != nil {
		return nil, err
	}

	err = repo.connect()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// -----------------
// methods
// -----------------
func (repo *WeatherReportRepository_MongoDB) initDBConfig() error {

	/* Reads external config from file */
	// data, err := ioutil.ReadFile("DBConfigMongo.json")
	// if err != nil {
	// 	return err
	// }

	// err = json.Unmarshal(data, &repo.DBConfig)
	// if err != nil {
	// 	return err
	// }

	// return nil

	/* Reads from environment variables */
	repo.DBConfig.Host = os.Getenv("DB_HOST")
	repo.DBConfig.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	repo.DBConfig.Username = os.Getenv("DB_USERNAME")
	repo.DBConfig.Password = os.Getenv("DB_PASSWORD")
	repo.DBConfig.Database = os.Getenv("DB_DATABASE")
	repo.DBConfig.Collection = os.Getenv("DB_COLLECTION")

	return nil
}

func (repo *WeatherReportRepository_MongoDB) connect() error {

	/* Outputs the config */
	fmt.Println("Connecting to MongoDB...")
	fmt.Println(repo.DBConfig.Host)
	fmt.Println(repo.DBConfig.Port)
	fmt.Println(repo.DBConfig.Username)
	fmt.Println(repo.DBConfig.Password)
	fmt.Println(repo.DBConfig.Database)
	fmt.Println(repo.DBConfig.Collection)

	/* Builds the connection string */
	port := strconv.Itoa(repo.DBConfig.Port)
	connectionString := "mongodb://" + repo.DBConfig.Username + ":" + repo.DBConfig.Password + "@" + repo.DBConfig.Host + ":" + port + "/?maxPoolSize=20&w=majority"
	clientOptions := options.Client().ApplyURI(connectionString)

	/* Connects to MongoDB */
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	/* Checks the connection */
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}

	/* Assigns the client to the repository */
	repo.Client = client

	fmt.Println("Connected to MongoDB!")
	return nil
}

func (repo *WeatherReportRepository_MongoDB) Create(p_weatherReport any) error {
	_, err := repo.Client.Database(repo.DBConfig.Database).Collection(repo.DBConfig.Collection).InsertOne(context.TODO(), p_weatherReport)
	return err
}

func (repo *WeatherReportRepository_MongoDB) Update(p_weatherReport any) error {
	p_wr := p_weatherReport.(domain.Daily)
	collection := repo.Client.Database(repo.DBConfig.Database).Collection(repo.DBConfig.Collection)

	/* Defines the filter to find the document based on the Time field */
	filter := bson.M{"time": p_wr.Time}

	/* Defines the update to set the new values for the document */
	update := bson.M{
		"$set": bson.M{
			"temperature_2m_max":         p_wr.Temperature2MMax,
			"temperature_2m_min":         p_wr.Temperature2MMin,
			"rain_sum":                   p_wr.RainSum,
			"windspeed_10m_max":          p_wr.Windspeed10MMax,
			"winddirection_10m_dominant": p_wr.Winddirection10MDominant,
		},
	}

	/* Updates the document in the collection */
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
