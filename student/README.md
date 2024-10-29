## Guess-It-2: Predicting Number Ranges with Linear Regression

### Description:

Guess-it-2 is a program that reads a series of numbers from standard input, calculates a linear regression line based on those numbers, and predicts the range in which the next number should fall. This project combines statistical knowledge (linear regression, standard error, prediction intervals) with Go programming, allowing the program to estimate future values with a degree of uncertainty.

### Features

* Reads a stream of numbers (representing the y-axis of a graph) from standard input.
* Automatically assigns x-axis values as the line numbers of the inputs.
* Uses linear regression to calculate the trend and predict the next number's range.
* Provides both lower and upper bounds of the predicted range for the next input.

### Files
1. main.go
        The main program that reads input, processes it using the linear regression functions, and prints out the predicted range for the next number.

2. student/stats/linearreg.go
        Contains the functions for calculating the slope, intercept, standard error, and prediction intervals using linear regression.

3. script.sh (Optional)
        A shell script to automate tasks such as compiling the Go program, running tests, or other project-related commands.

### Setup
### Requirements

* Go Programming Language (version 1.18 or higher)

### Installation

Clone the repository:
```
git clone https://learn.zone01kisumu.ke/git/shaokoth/guess-it-2
```

Move into student directory:
```
cd student/
```

Run the program:

```
go run main.go
```

### How to use

1. Run the program using the following command:

```
go run . main.go
```

2. Enter a sequence of numbers as input, each on a new line. After each input, the program will output a range for the next number.


### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.