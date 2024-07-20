# Guess It 1

Guess It 1 is a command-line tool written in Go that reads numbers from standard input, calculates the average, variance, and standard deviation, and determines a range based on these statistics. The range is calculated as mean ± 1.5 * standard deviation.

## Table of Contents

- [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [Author](#author)

## Installation

To use this project, you need to have Go installed. Follow the steps below to get started:

1. Clone the repository:
   ```sh
   git clone https://learn.zone01kisumu.ke/git/bernaotieno/guess-it-1.git
   cd guess-it-1
   ```


## Project Structure
 * `main.go`: The main entry point for the application.
 * `inputhandler/handler.go`: Handles user input and calls calculation functions.
 * `calculation/average.go`: Calculates the average of a slice of float64 numbers.
 * `calculation/range.go`: Calculates the range based on mean and standard deviation.
 * `calculation/standarddev.go`: Calculates the standard deviation from variance.
 * `calculation/variance.go`: Calculates the variance of a slice of float64 numbers.

## Usage
To run the project, use the provided bash script:

```sh
./run.sh
```
Enter numbers one by one and press Enter. The program will calculate and print the range after each entry. To exit, simply press Enter without typing a number.

## Testing
**To test the project, follow these steps:**

 1. Download the test [zip file](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip) containing the tester.
 2. Extract the zip file in the root directory of the project.
 3. Move the student/ folder into the extracted directory.
 4. Follow the instructions provided in the tester's README file.

## Contributing
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## Author
 * [Benard Okumu](https://learn.zone01kisumu.ke/git/bernaotieno)
