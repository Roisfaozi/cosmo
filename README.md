# Cosmo Image Metadata Tool

Cosmo is a command-line tool for managing and updating image metadata in bulk. It offers functionality to rename images sequentially and update metadata based on a CSV file input. This project leverages ExifTool to handle metadata extraction and updating.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
    - [Rename Command](#rename-command)
    - [Update Metadata Command](#update-metadata-command)
    - [Process Command](#process-command)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Bulk Metadata Update**: Update image metadata (e.g., title, keywords, copyright status) based on a CSV file.
- **Sequential Image Renaming**: Rename images in a specified directory and export the results to a CSV file.
- **Combined Processing**: Run both renaming and metadata updates with a single command.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/cosmo.git
   cd cosmo
   ```

2. Install the required dependencies:
   ```bash
   go get ./...
   ```
3.  Build the project:
    ```bash
    go build -o cosmo ./cmd/main.go
    ```

## Usage
### Rename Command

The ```rename``` command renames all images in a specified directory sequentially and exports the renaming information to a CSV file.

##### Command Syntax
```bash
./cosmo rename [directory] --ext=[file extension]
```
##### Example
```bash
./cosmo rename ./testdata --ext=".jpg"
```
This command renames all .jpg images in the ./testdata directory and saves the renaming details in renamed_files.csv within the same directory.

### Update Metadata Command

The ```update-metadata``` command updates the metadata of all images in a specified directory based on a CSV file.

##### Command Syntax
```bash
./cosmo update-metadata [csv-file]
```
##### Example        
```bash
./cosmo update-metadata ./testdata/metadata.csv
```
The CSV file should have the following columns:

- ```SourceFile```: Path to the image file
- ```ObjectName```: Title or name of the object in the image
- ```Keywords```: Keywords associated with the image
- ```CopyrightStatus```: Copyright status of the image
- ```Marked```: Marked status (e.g., TRUE or FALSE)
- ```CopyrightNotice```: Copyright notice information

### Process Command

The ```process``` command combines the functionality of both the ```rename``` and ```update-metadata``` commands.

##### Command Syntax
```bash
./cosmo process [directory] --ext=[file extension]
```
##### Example
```bash
./cosmo process ./testdata --ext=".jpg"
```
This command renames all .jpg images in the ./testdata directory and updates the metadata based on the metadata.csv file.

## Testing

To run the tests, run the following command:
```bash
go test ./...
```
For verbose output:
```bash
go test -v ./... 
```
To run specific tests, use the -run flag:
```bash
go test -v -run=TestUpdateMetadata  ./...
```
This project uses testify for assertions and exiftool for metadata handling.

## Contributing 

If you would like to contribute to this project, please open an issue or submit a pull request. Thank you!
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes and push your branch.
4. Open a pull request describing your changes.

## License

This project is licensed under the [MIT License](LICENSE).

### Support this project
If you find this project useful, please consider supporting it by making a donation.

[![PayPal Button](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?business=roisfaozi55@gmail.com&cmd=_donations&currency_code=USD)