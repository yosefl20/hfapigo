package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"time"

	"github.com/yosefl20/hfapigo"
)

const HuggingFaceTokenEnv = "HUGGING_FACE_TOKEN"

func init() {
	key := os.Getenv(HuggingFaceTokenEnv)
	if key != "" {
		hfapigo.SetAPIKey(key)
	}
}

func main() {
	input := "a beautiful landscape, oil paint style"
	const numReturnSequences = 3

	fmt.Printf("Input: \"%s\"\n", input)

	type ChanRv struct {
		resps image.Image
		err   error
	}
	ch := make(chan ChanRv)

	fmt.Print("Sending request")
	go func() {
		resps, err := hfapigo.SendTextToImageRequest(hfapigo.RecommendedTextToImageModel, &hfapigo.TextToImageRequest{
			Inputs: input,
		})
		ch <- ChanRv{resps, err}
	}()

	for {
		select {
		case chrv := <-ch:
			fmt.Println()
			if chrv.err != nil {
				fmt.Println(chrv.err)
				return
			}
			out, _ := os.Create("./output.jpeg")
			defer out.Close()

			var opts jpeg.Options
			opts.Quality = 90

			if err := jpeg.Encode(out, chrv.resps, &opts); err != nil {
				fmt.Println(err)
			}
			fmt.Println("image saved to output.jpeg")
			return
		default:
			fmt.Print(".")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
