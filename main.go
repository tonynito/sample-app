/**
 * Copyright 2023 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)

func main() {
	http.HandleFunc("/blue", blueHandler)
	http.HandleFunc("/red", redHandler)
	http.HandleFunc("/purple", purpleHandler)
	http.ListenAndServe(":8080", nil)
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func redHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func purpleHandler(w http.ResponseWriter, r *http.Request) {
    // Create a 100x100 RGBA image
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))

    // Define a purple color: Red + Blue = Purple
    purple := color.RGBA{128, 0, 128, 255} // RGB(128, 0, 128)
    
    // Fill the image with purple
    draw.Draw(img, img.Bounds(), &image.Uniform{purple}, image.Point{}, draw.Src)

    // Set the content type to PNG
    w.Header().Set("Content-Type", "image/png")

    // Encode the image as PNG
    if err := png.Encode(w, img); err != nil {
        http.Error(w, "Failed to encode image", http.StatusInternalServerError)
    }
}
