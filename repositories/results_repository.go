package repositories

import (
	"database/sql"
	"net/http"

	"github.com/Vkanhan/go-marathon/models"
)

// ResultsRepository handles database operations for results.
type ResultsRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewResultsRepository(dbHandler *sql.DB) *ResultsRepository {
	return &ResultsRepository{
		dbHandler: dbHandler,
	}
}

func (rr ResultsRepository) CreateResult(result *models.Result) (*models.Result, *models.ResponseError) {
	query := `
	INSERT INTO results(runner_id, race_result, location, position, year)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	rows, err := rr.dbHandler.Query(query, result.RunnerID, result.RaceResult, result.Location, result.Position, result.Year)
	if err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	defer rows.Close()

	var resultId string
	for rows.Next() {
		if err := rows.Scan(&resultId); err != nil {
			return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	return &models.Result{
		ID: resultId, RunnerID: result.RunnerID, RaceResult: result.RaceResult, Location: result.Location, Position: result.Position, Year: result.Year,
	}, nil
}

func (rr ResultsRepository) DeleteResult(resultId string) (*models.Result, *models.ResponseError) {
	query := `DELETE FROM results WHERE id = $1 RETURNING runner_id, race_result, year`
	rows, err := rr.transaction.Query(query, resultId)
	if err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	defer rows.Close()

	var runnerId, raceResult string
	var year int
	for rows.Next() {
		if err := rows.Scan(&runnerId, &raceResult, &year); err != nil {
			return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	return &models.Result{ID: resultId, RunnerID: runnerId, RaceResult: raceResult, Year: year}, nil
}

func (rr ResultsRepository) GetAllRunnersResults(runnerId string) ([]*models.Result, *models.ResponseError) {
	query := `SELECT id, race_result, location, position, year FROM results WHERE runner_id = $1`
	rows, err := rr.dbHandler.Query(query, runnerId)
	if err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	defer rows.Close()

	results := make([]*models.Result, 0)
	var id, raceResult, location string
	var position, year int

	for rows.Next() {
		if err := rows.Scan(&id, &raceResult, &location, &position, &year); err != nil {
			return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
		}
		results = append(results, &models.Result{ID: id, RunnerID: runnerId, RaceResult: raceResult, Location: location, Position: position, Year: year})
	}

	if err := rows.Err(); err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	return results, nil
}

func (rr ResultsRepository) GetPersonalBestResults(runnerId string) (string, *models.ResponseError) {
	query := `SELECT MIN(race_result) FROM results WHERE runner_id = $1`
	rows, err := rr.dbHandler.Query(query, runnerId)
	if err != nil {
		return "", &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	defer rows.Close()

	var raceResult string
	for rows.Next() {
		if err := rows.Scan(&raceResult); err != nil {
			return "", &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
		}
	}

	if err := rows.Err(); err != nil {
		return "", &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	return raceResult, nil
}

func (rr ResultsRepository) GetSeasonBestResults(runnerId string, year int) (string, *models.ResponseError) {
	query := `SELECT MIN(race_result) FROM results WHERE runner_id = $1 AND year = $2`
	rows, err := rr.dbHandler.Query(query, runnerId, year)
	if err != nil {
		return "", &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	defer rows.Close()

	var raceResult string
	for rows.Next() {
		if err := rows.Scan(&raceResult); err != nil {
			return "", &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
		}
	}

	if err := rows.Err(); err != nil {
		return "", &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	return raceResult, nil
}
