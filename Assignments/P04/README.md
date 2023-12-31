## Program 4 - Concurrent/Parallel Image Downloader
### Ethan Bailey
### Description:

This program invloves functions that download images from a specified url.
One downloads the images individually, and the other two receives a slice of urls
and then calls the first function to download them all at once.
They do so sequentially and concurrently through the use of channels and goroutines.

### Files

|   #   | Folder             | Description                                        |
| :---: | ---------------- | -------------------------------------------------- |
|   1   | [main.go](https://github.com/EthanJBailey/4143-PLC/tree/main/Assignments/P04/main.go)          | Main driver that calls the ColorTest module. |
|   2   | [go.mod](https://github.com/EthanJBailey/4143-PLC/tree/main/Assignments/P04/go.mod)            | Base go file that defines all the module's module path |
|   3   | [images](https://github.com/EthanJBailey/4143-PLC/tree/main/Assignments/P04/images)              | A folder containing the images downloaded while testing the program      |
|   4   | [README.md](https://github.com/EthanJBailey/4143-PLC/tree/main/Assignments/P04/README.md)          | This readme file explaining how the code works     |

### Instructions

- Make sure you download pull all files from the package that you want to use
- My program expects you to use the existing URLs or enter your own to download images.
- Execution Command:
  - `go run main.go`
  

### Results

- The following tables contains the download times for the sequential and concurrent functions.
- Based on the following times, it can be inferred that downloading images sequentially is marginally slower
  than downloading them concurrently. 
-  The average of the times below are 540.25952 ms for the sequential time and 417.24366 ms.
-  Therefore, there is on average a 30% difference in the time it takes to download the images.

| #   | Sequential Time (ms) | Concurrent Time (ms) |
| ------------    | ----------- | ----------- |
| 1st run         | 506.7567 ms | 422.2275 ms |
| 2nd run         | 575.9808 ms | 547.1067 ms |
| 3rd run         | 673.0518 ms | 357.4730 ms |
| 4th run         | 502.8538 ms | 316.7495 ms |
| 5th run         | 442.6545 ms | 442.6616 ms |

  
