package repository

import (
	"TrafficPolice/internal/domain"
	"TrafficPolice/internal/repository"
	"context"
	"github.com/jackc/pgx/v5"
)

type cameraRepoPostgres struct {
	conn *pgx.Conn
}

func NewCameraRepoPostgres(conn *pgx.Conn) repository.CameraDB {
	return &cameraRepoPostgres{conn: conn}
}

func (db *cameraRepoPostgres) AddCameraType(cameraType domain.CameraType) error {
	query := "INSERT INTO camera_types (camera_type_id, camera_type_name) VALUES ($1, $2)"

	_, err := db.conn.Exec(context.Background(), query, cameraType.ID, cameraType.Name)
	if err != nil {
		return err
	}

	return nil
}

func (db *cameraRepoPostgres) RegisterCamera(camera domain.Camera) error {
	query := `INSERT INTO cameras (camera_id, camera_type_id, camera_latitude, camera_longitude, short_desc) 
		VALUES ($1, $2, $3, $4, $5)`

	_, err := db.conn.Exec(context.Background(), query,
		camera.ID, camera.CameraType.ID, camera.Latitude, camera.Longitude, camera.ShortDesc)

	if err != nil {
		return err
	}

	return nil
}
