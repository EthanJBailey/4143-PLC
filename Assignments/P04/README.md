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
|   3   | [go.sum](https://github.com/EthanJBailey/4143-PLC/tree/main/Assignments/P04/go.sum)            | Go file that contains information used by Go to record specific hashes and versions of dependencies |
|   4   | [images](https://github.com/EthanJBailey/4143-PLC/tree/main/Assignments/P04/go.sum)              | A folder containing the images downloaded while testing the program      |
|   5   | README.md          | This readme file explaining how the code works     |

### Instructions

- Make sure you download pull all files from the package that you want to use
- My program expects you to use the existing URLs or enter your own to download images.
- Execution Command:
  - `go run main.go`
  

### Results

- The following tables contains the download times for the sequential and concurrent functions.
  | #   | Sequential Time (ms) | Concurrent Time (ms) |
  | ------------    | ----------- | ----------- |
  |                 | 506.7567 ms | 422.2275 ms |
  |                 | 575.9808 ms | 547.1067 ms |
  |                 | 673.0518 ms | 357.4730 ms |
  |                 | 502.8538 ms | 316.7495 ms |
  |                 | 442.6545 ms | 442.6616 ms |

  
