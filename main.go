package main

import (
  "image"
  "image/color"
  "image/png"
  "image/draw"
  "os"
  "math/rand"
)

const WIDTH = 400
const HEIGHT = 400
const KILLRATE = 60
const BIRTHRATE = 55
const MAXPOINTS = 100

func remove(slice []image.Point, s int) []image.Point {
    return append(slice[:s], slice[s+1:]...)
}

func main() {
  upLeft := image.Point{0, 0}
  lowRight := image.Point{WIDTH, HEIGHT}

  img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
  draw.Draw(img, img.Bounds(), image.White, image.Pt(0, 0), draw.Src)

  points := []image.Point{image.Point{int(WIDTH / 2), 0}}

  for len(points) > 0 {
    tempPoints := []image.Point{} 

    for i, point := range points {
      img.Set(point.X, point.Y, color.Black)
      
      // If 0, kill this point
      deathNum := rand.Intn(KILLRATE)
      
      if deathNum == 0 && len(points) > 2{
        continue
      }

      randNum := rand.Intn(70)
      switch {
      // Move left if not on border
      case randNum < 32:
        if point.X >= WIDTH - 1 {
          continue
        } else {
          tempPoints = append(tempPoints, image.Point{points[i].X + 1, points[i].Y})
        }
      case randNum < 64:
        if point.X <= 0 {
          continue
        } else {
          tempPoints = append(tempPoints, image.Point{points[i].X - 1, points[i].Y})
        }
      default:
        if point.Y >= HEIGHT - 1 {
          continue
        } else {
          tempPoints = append(tempPoints, image.Point{points[i].X, points[i].Y + 1})
        }
      }
      createNum := rand.Intn(BIRTHRATE)

      if createNum == 0 && len(points) < MAXPOINTS {
        tempPoints = append(tempPoints, image.Point{point.X, point.Y})
      }
    }
    points = tempPoints
  }

  // Encode as PNG.
  f, _ := os.Create("image.png")
  png.Encode(f, img)
}
