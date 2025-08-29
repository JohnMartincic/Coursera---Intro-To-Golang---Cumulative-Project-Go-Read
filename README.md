# Coursera - Cumulative Go Project - Go Read

This repository contains my final, cumulative project, "go read," developed from start to finish as a capstone challenge. The project is a command-line tool that allows users to search for books by title from a predefined list.

## 🎯 Project Goal

The application was built to satisfy a simple but practical use case:

1. Prompt a user to enter a search term.
2. Read the user's input from the console.
3. Search through a list of book titles for any that contain the user's search term.
4. Print a list of all matching book titles, handling the search in a case-insensitive manner.

## ✨ Key Features & Implementation

This project was built from scratch and demonstrates the ability to combine several fundamental Go concepts to create a functional application.

* **Interactive User Input**: The program waits for and accepts a string directly from the user in the console.

* **Case-Insensitive Search Logic**: To ensure the search is not case-sensitive, the application uses the Go standard `strings` package. Both the user's input and the book titles from the data source are converted to a consistent case (lowercase) before the comparison is made.

* **Partial String Matching**: The core search functionality is powered by the `strings.Contains()` function, which efficiently checks if a book title includes the user's search query as a substring.

* **Dynamic Results Display**: The program iterates through the entire collection of books and prints only the titles that match the search criteria, providing a clean list of results to the user.

## 💡 Key Concepts Applied

This challenge required the practical application of several key Golang concepts in a single program:

* **Standard Library Usage**: Heavy reliance on the `fmt` package for console I/O and the `strings` package for text manipulation (`ToLower`, `Contains`).
* **Data Structures**: Storing the list of book titles in a slice (`[]string`) for easy iteration.
* **Control Flow**: Using a `for` loop with `range` to iterate through the book data and an `if` statement to apply the filtering logic.
* **User Input Handling**: Reading and processing input from the command line.

## 🎓 Project Context

This project serves as a cumulative, start-to-finish challenge, requiring the independent implementation of all concepts learned throughout the Golang for Beginners series on Coursera. It demonstrates the ability to build a complete, functional command-line application from a set of requirements.

* **Course Link:** [Golang for Beginners: Data Types, Functions, and Packages](https://www.coursera.org/projects/golang-beginners-data-types-functions-packages)
