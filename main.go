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
	http.HandleFunc("/red", redHandler)
	http.HandleFunc("/purple", purpleHandler)
	http.HandleFunc("/yellow", yellowHandler)
	http.ListenAndServe(":8080", nil)
}

func redHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func purpleHandler(w http.ResponseWriter, r *http.Request) {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    // Purple = R:128, G:0, B:128 (fully opaque)
    draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{128, 0, 128, 255}}, image.Point{}, draw.Src)
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, img)
}

func yellowHandler(w http.ResponseWriter, r *http.Request) {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    // Yellow = R:255, G:255, B:0 (fully opaque)
    draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 0, 255}}, image.Point{}, draw.Src)
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, img)
}

