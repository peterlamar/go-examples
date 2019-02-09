package pythonmovie

import (
	"fmt"
	"time"

	redisCache "github.com/go-redis/cache"
	"github.com/peterlamar/go-examples/sqltesting/database"
	log "github.com/sirupsen/logrus"
)

// MovieName stores the name of the Monty Python movie
type MovieName struct {
	TableID     int    `db:"table_id"`
	DataString  string `db:"movie_title"`
	DomesticBox int    `db:"domestic_box"`
	WordwideBox int    `db:"worldwide_box"`
}

// GetMovieInfo queries the name of the Monty Python voie
func GetMovieInfo(movieID int) (rtn MovieName) {

	var err error

	QueryKeyFormat := "simpletable:%d"

	queryKey := fmt.Sprintf(QueryKeyFormat, movieID)

	// If there is a cache miss
	if database.IsCacheConnected() {
		err = database.GetRedisCacheCodec().Get(queryKey, &rtn)
	}

	if err != nil || !database.IsCacheConnected() {

		query := `select
    		table_id,
    		movie_title,
				domestic_box,
				worldwide_box
    	from
    		simpletable
    	where
    		table_id = :table_id`

		// Get data from db
		nstmt, err := database.GetDB().PrepareNamed(query)

		if err != nil {
			log.Fatal(err)
		}

		args := map[string]interface{}{
			"table_id": movieID,
		}

		err = nstmt.Get(&rtn, args)

		if err != nil {
			log.Fatal(err)
		}

		if database.IsCacheConnected() {
			// Place data into cache for next time, expire in 1 hour
			err = database.GetRedisCacheCodec().Set(&redisCache.Item{
				Key:        queryKey,
				Object:     rtn,
				Expiration: time.Hour,
			})
			if err != nil {
				panic(err)
			}
		}
		log.Debug("DB hit GetMovieName")
	} else {
		log.Debug("Cache hit GetMovieName")
	}

	return
}
