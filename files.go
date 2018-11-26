package main

import (
  "fmt"
  "time"
  "sync"
  "os"
  "os/exec"
  "path/filepath"
  // "strings"
)
var wg sync.WaitGroup

func preprocessing(n int) {
  defer wg.Done()

  fmt.Printf(">>>>>> %d \n", n)
  out, err := exec.Command("python","app.py").Output()
    if err != nil {
        fmt.Println("error occured")
        fmt.Printf("%s", err)
    }
  
  fmt.Printf(">>>>>> %s \n", out)
  // time.Sleep(time.Second * time.Duration(7))
}

func checksize(path string) {
  defer wg.Done()

  // newPath := strings.Replace( path , ".png", ".jpg", -1 )

  // Rename or move file from one location to another.
  // os.Rename(path, newPath)

  fmt.Println(path)
}

func main() {

  start := time.Now()
	
  // for i := 0; i < 100; i++ {
  //   wg.Add(1)
  //   go preprocessing(i)
  // }
  // wg.Wait()


  var files []string

  root := "/mnt/hdd/Projects/image-regression/images/"

  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
      files = append(files, path)
      return nil
  })
  if err != nil {
      panic(err)
  }

  for _, file := range files {
    wg.Add(1)
    go checksize(file)
  }
  
  wg.Wait()
  elapsed := time.Since(start)
  fmt.Printf("Preprocessing  took %s", elapsed)

}