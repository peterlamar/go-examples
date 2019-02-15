# GDrive in golang

This example lists files, creates a directory and uploads an
image into the directory in google drive. You will need an 
active Google account with google drive and credentials.json 
in the root directory of this project. 

## Usage

List files
```golang
// List all files in the google drive
if len(r.Files) == 0 {
    fmt.Println("No files found.")
} else {
    for _, i := range r.Files {
        fmt.Printf("%s (%s)\n", i.Name, i.Id)
    }
}
```

Create a dir and upload a file
```golang
// Open local file
f, err := os.Open("./knightsofni.png")

if err != nil {
    panic(fmt.Sprintf("cannot open file: %v", err))
}

// Create the directory
dir, err := createDir(srv, "My Folder", "root")

if err != nil {
    panic(fmt.Sprintf("Could not create dir: %v\n", err))
}

defer f.Close()

// Create the file and upload its content
file, err := createFile(srv, "uploaded-image.png", "image/png", f, dir.Id)

if err != nil {
    panic(fmt.Sprintf("Could not create file: %v\n", err))
}
fmt.Printf("File '%s' successfully uploaded in '%s' directory", file.Name,
    dir.Name)
```

## Run the code

go run main.go


## References

* [Upload files in Google Drive](https://medium.com/@devtud/upload-files-in-google-drive-with-golang-and-google-drive-api-d686fb62f884)
* [Golang quickstart](https://developers.google.com/drive/api/v3/quickstart/go)