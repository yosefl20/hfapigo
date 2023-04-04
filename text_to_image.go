package hfapigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
)

const RecommendedTextToImageModel = "stabilityai/stable-diffusion-2-1"

type TextToImageRequest struct {
	// (Required) a string to be generated from
	Input string `json:"input,omitempty"`
}

func SendTextToImageRequest(model string, request *TextToImageRequest) (image.Image, error) {
	if request == nil {
		return nil, errors.New("nil SummarizationRequest")
	}

	jsonBuf, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, err := MakeHFAPIRequest(jsonBuf, model)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(respBody))
	if err != nil {
		return nil, err
	}

	return img, nil
}
