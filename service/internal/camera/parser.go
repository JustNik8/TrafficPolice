package camera

import (
	"TrafficPolice/internal/services"
	"TrafficPolice/internal/transport/rest/dto"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/jcalabro/leb128"
	"log"
	"time"
)

const (
	typeCamerus1 = "camerus1"
	typeCamerus2 = "camerus2"
	typeCamerus3 = "camerus3"

	cameraIDKey = "camera_id"
	cameraKey   = "camera"
)

type Parser struct {
	cameraService services.CameraService
}

func NewParser(cameraService services.CameraService) *Parser {
	return &Parser{
		cameraService: cameraService,
	}
}

func (p *Parser) ParseCameraInfo(payload []byte) (dto.Case, error) {
	if len(payload) == 0 {
		return dto.Case{}, fmt.Errorf("payload is empty")
	}

	payload = payload[2:]
	info := p.parsePayload(payload)
	log.Println("INFO:")
	log.Println(info)

	var cameraType string
	var err error

	if cameraID, ok := info[cameraIDKey]; ok {
		cameraType, err = p.cameraService.GetCameraTypeByCameraID(cameraID.(string))
	} else if camera, ok := info[cameraKey].(map[string]any); ok {
		cameraID := camera["id"].(string)
		cameraType, err = p.cameraService.GetCameraTypeByCameraID(cameraID)
	} else {
		err = errors.New("not known cameraID")
	}

	if err != nil {
		return dto.Case{}, err
	}

	switch cameraType {
	case typeCamerus1:
		return p.parseCamerus1(info)
	case typeCamerus2:
		return p.parseCamerus2(info)
	case typeCamerus3:
		return p.parseCamerus3(info)
	default:
		return dto.Case{}, errors.New("unknown camera type")
	}
}

func (p *Parser) parsePayload(payload []byte) map[string]any {
	info := make(map[string]any)

	for len(payload) > 0 {
		keySize := binary.BigEndian.Uint16(payload[:2])
		payload = payload[2:]

		valueSize := binary.BigEndian.Uint16(payload[:2])
		payload = payload[2:]

		valueType := payload[0]
		payload = payload[1:]

		keyValue := payload[:keySize]
		payload = payload[keySize:]

		value := payload[:valueSize]
		payload = payload[valueSize:]

		log.Println(keySize, valueSize, valueType, keyValue, value)

		if valueType == 0 {
			info[string(keyValue)] = string(value)
		} else if valueType == 1 {
			n, _ := leb128.DecodeS64(bytes.NewBuffer(value))
			info[string(keyValue)] = n
		} else if valueType == 2 {
			dict := p.parsePayload(value)
			info[string(keyValue)] = dict
		}
	}

	return info
}

func (p *Parser) parseCamerus1(info map[string]any) (dto.Case, error) {
	date, err := time.Parse(time.RFC3339, info["datetime"].(string))
	if err != nil {
		return dto.Case{}, err
	}
	return dto.Case{
		Transport: dto.Transport{
			Chars:  info["transport_chars"].(string),
			Num:    info["transport_numbers"].(string),
			Region: info["transport_region"].(string),
		},
		Camera: dto.Camera{
			ID: info["camera_id"].(string),
		},
		Violation: dto.Violation{
			ID: info["violation_id"].(string),
		},
		ViolationValue: info["violation_value"].(string),
		RequiredSkill:  info["skill_value"].(int64),
		Date:           date,
	}, nil
}

func (p *Parser) parseCamerus2(info map[string]any) (dto.Case, error) {
	transport := info["transport"].(map[string]any)
	camera := info["camera"].(map[string]any)
	violation := info["violation"].(map[string]any)
	skill := info["skill"].(map[string]any)

	datetime := info["datetime"].(map[string]any)

	year := datetime["year"].(int64)
	month := datetime["month"].(int64)
	day := datetime["day"].(int64)
	hour := datetime["hour"].(int64)
	minute := datetime["minute"].(int64)
	seconds := datetime["seconds"].(int64)
	utcOffset := datetime["utc_offset"].(string)

	dateString := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%s", year, month, day, hour, minute, seconds, utcOffset)
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return dto.Case{}, err
	}

	return dto.Case{
		Transport: dto.Transport{
			Chars:  transport["chars"].(string),
			Num:    transport["numbers"].(string),
			Region: transport["region"].(string),
		},
		Camera: dto.Camera{
			ID: camera["id"].(string),
		},
		Violation: dto.Violation{
			ID: violation["id"].(string),
		},
		ViolationValue: violation["value"].(string),
		RequiredSkill:  skill["value"].(int64),
		Date:           date,
	}, nil

}

func (p *Parser) parseCamerus3(info map[string]any) (dto.Case, error) {
	transport := info["transport"].(string)
	chars := transport[1:4]
	num := string(transport[0]) + transport[4:6]
	region := transport[6:]

	camera := info["camera"].(map[string]any)
	violation := info["violation"].(map[string]any)

	return dto.Case{
		Transport: dto.Transport{
			Chars:  chars,
			Num:    num,
			Region: region,
		},
		Camera: dto.Camera{
			ID: camera["id"].(string),
		},
		Violation: dto.Violation{
			ID: violation["id"].(string),
		},
		ViolationValue: violation["value"].(string),
		RequiredSkill:  info["skill"].(int64),
		Date:           time.Unix(info["datetime"].(int64), 0),
	}, nil
}
