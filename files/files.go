package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the example file and print its metadata.
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(f.Name())

	fileInfo, err := f.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Println("Is directory:", fileInfo.IsDir())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file mode:", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())

	defer f.Close()

	// Read and print the first 11 bytes from the open file.
	buf := make([]byte, 11)

	d, err := f.Read(buf)

	if err != nil {
		panic(err)
	}

	fmt.Println(d)
	for i := 0; i < len(buf); i++ {
		fmt.Println("Data is:", string(buf[i]))
	}

	// Read and print the complete file contents.
	data, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// List the names of entries in the current directory.
	dirFile, err := os.Open(".")

	if err != nil {
		panic(err)
	}

	defer dirFile.Close()

	dir, err := dirFile.ReadDir(-1)

	if err != nil {
		panic(err)
	}

	for _, fi := range dir {
		fmt.Println(fi.Name())
	}

	// Create a destination file and write sample data to it.
	file, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// file.WriteString("welcome to golang")

	bytes := []byte("hello golang")

	file.Write(bytes)

	// Copy the source file to the destination using buffered I/O.
	sourceFile, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destinationFile, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	defer destinationFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destinationFile)

	for {
		dataByte, err := reader.ReadByte()

		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(dataByte)

		if e != nil {
			panic(e)
		}
	}

	writer.Flush()
	fmt.Println("writing is done")

	// Remove the destination file created by this example.
	e := os.Remove("example2.txt")

	if e != nil {
		panic(e)
	}
}
