/*************************************************************************************************
*
*  Email:            ejbailey1030@my.msutexas.edu
*  Label:            P04
*  Title:            Concurrent/Parallel Image Downloader
*  Course:           CMPS 4143
*  Semester:         Fall 2023
*
*  Description:
*        This program downloads multiple images across several different sites.
*        It does so both concurrently and sequentially through the use of goroutines
*        and channels respectively. It also count the total time each version takes
*        to complete.
*
*  Usage:
*        - Run the main.go file
*
*  Files:
*        main.go   			 :  driverprogram
*        go.mod	   			 :  defines the module's module path
*        go.sum	   			 :  record specific hashes and versions of dependencies
*
*************************************************************************************************/

// Package declaration
package main

// Import required packages
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Slice of strings that contain URLs of images
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://cdn.stocksnap.io/img-thumbs/960w/nature-mountains_5QRJ3CKJDI.jpg",
		"https://images.pexels.com/photos/164729/pexels-photo-164729.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://cdn.stocksnap.io/img-thumbs/960w/vinyl-music_JR3YRE4G2N.jpg",
		"https://images.unsplash.com/photo-1531986627054-7f294d095acd?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1646330024721-d726c353e519?q=80&w=2071&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://cdn.stocksnap.io/img-thumbs/960w/simple-pumpkin_WPXAMOFWKY.jpg",
		"https://images.squarespace-cdn.com/content/v1/60a31f4004aa3e0256864270/c310ba3d-b756-4479-811e-da4a56166b15/image_1700493828815821000.jpg",
		"https://images.squarespace-cdn.com/content/v1/60a31f4004aa3e0256864270/ad1cb4c0-734a-4707-8840-3b8068336404/image_1700493828316511000.jpg",
		"https://cdn.pixabay.com/photo/2023/11/04/07/57/owl-8364426_1280.jpg",
		"https://cdn.stocksnap.io/img-thumbs/960w/rose-flowers_UHDDRHM5UZ.jpg",
		"https://www.pexels.com/photo/18705295/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.squarespace-cdn.com/content/v1/60a31f4004aa3e0256864270/ad1cb4c0-734a-4707-8840-3b8068336404/image_1700493828316511000.jpg",
		"https://images.unsplash.com/photo-1575470522418-b88b692b8084?q=80&w=3288&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1699704674521-54b3252e52e9?q=80&w=2072&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.freeimages.com/images/large-previews/77c/nemo-the-horse-1339807.jpg",
	}

	// Print Program Heading
	fmt.Println("Program 4 - Image Downloader")

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// Create a new `http.Request` object.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a new `http.Client` object.
	client := &http.Client{}

	// Do the request and get the response.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response status code:", resp.StatusCode)
		return err
	}

	// Create a new file to save the image to.
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Copy the image from the response body to the file.
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Close the file.
	f.Close()
	return nil
}

/**
* Private: downloadImagesConcurrent
*
* Description:
*      Function to concurrently download images using goroutines
*
* Params:
*      urls []string
*
* Returns:
*      Downloaded images to the specified folder
 */
func downloadImagesConcurrent(urls []string) {
	// Used for synchronization between goroutines.
	var wg sync.WaitGroup
	// Loop through slice of urls
	for i, url := range urls {
		// Add one to wait group for each goroutine
		wg.Add(1)
		go func(index int, u string) {
			// Decrements the WaitGroup counter by one; done when this function returns.
			defer wg.Done()
			// Download the image
			filename := "images/Image" + strconv.Itoa(index+1) + ".jpg"
			_ = downloadImage(u, filename)
		}(i, url)
	}
	// Wait until all goroutines are done
	wg.Wait()
}

//
/**
* Private: downloadImagesSequential
*
* Description:
*      Function to sequentially download images using channels
*
* Params:
*      urls []string
*
* Returns:
*      Downloaded images to the specified folder
*		OR returns error message
 */
func downloadImagesSequential(urls []string) {
	// Creates channels
	ch := make(chan error)
	// Loop through slice of urls, start downloading images sequentially
	for i, url := range urls {
		filename := "images/Image" + strconv.Itoa(i+1) + ".jpg"
		go func(url string, ch chan<- error) {
			// Downloads an image and sends it back on the channel
			ch <- downloadImage(url, filename)
		}(url, ch)
	}
	// Loops through urls and print errors, if any
	for range urls {
		e := <-ch
		if e != nil {
			log.Fatal("Error downloading images: ", e)
		}
	}
}
