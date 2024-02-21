package database

import "TrafficPolice/internal/models"

type CameraDB interface {
	AddCameraType(cameraType models.CameraType) error
	RegisterCamera(camera models.Camera) error
}