# Exercism Go Track Solutions

Welcome to my repository where I document my solutions to the Go (Golang) exercises on [Exercism.org](https://exercism.org). This repository serves as a personal archive of my learning journey through the Go programming language, as well as a resource for others who might be looking for different approaches to solving these exercises.

## Table of Contents

- [About Exercism](#about-exercism)
- [Repository Structure](#repository-structure)
- [How to Use This Repository](#how-to-use-this-repository)
- [Contributing](#contributing)
- [License](#license)

## About Exercism

Exercism is an online platform designed to help you improve your coding skills through practice and mentorship. It offers exercises in over 50 different programming languages, including Go. Each exercise comes with a set of tests that your solution must pass, and you can receive feedback from mentors to improve your code.

## Repository Structure

This repository is organized by the exercise names as provided by Exercism. Each exercise is contained within its directory, which includes:

- A `README.md` file with a brief description of the exercise and any relevant links.
- A Go source file (`exercise_name.go`) containing my solution.
- A test file (`exercise_name_test.go`) which is usually provided by Exercism to validate the solution.

```
/go-learning
│
├── /hello-world
| ├── go.mod
│ ├── README.md
| ├── HELP.md
│ ├── hello_world.go
│ └── hello_world_test.go
│
├── /lasagna
| ├── go.mod
│ ├── README.md
| ├── HELP.md
| ├── HINTS.md
│ ├── lasagna.go
│ └── lasagna_test.go
│
└── ...
```
## How to Use This Repository

1. **Clone the Repository**: Start by cloning this repository to your local machine using `git clone https://github.com/pumnea/go-learning.git`.
2. **Navigate to an Exercise**: Change into the directory of the exercise you're interested in.
3. **Review the Solution**: Open the `.go` file to see my solution to the exercise.
4. **Run the Tests**: You can run the tests using the Go test command to see if the solution passes all the test cases.

```bash

cd path/to/exercise
go test
```


## Contributing

While this repository is primarily a personal archive, I welcome any constructive feedback, suggestions, or discussions about the solutions provided. If you have a different approach or an improvement, feel free to open an issue or submit a pull request.

## License

This repository is licensed under the MIT License. This means you are free to use, modify, and distribute the code as you see fit, as long as you include the original license.