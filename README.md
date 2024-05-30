
# Installing Tree

## Prerequisites

Before you begin, make sure you have the following installed on your system:

- Git: You can download and install Git from [here](https://git-scm.com/downloads).
- Go Programming Language: You can download and install Go from [here](https://golang.org/dl/).

## Steps

1. **Clone the Git Repository:**

   Open your terminal and clone the repository.

   ```bash
   git clone https://github.com/robert-ohurley/tree
   ```
2. **Install:**
   
   cd into the repository and install with the following command.


   ```bash
   go install
   ```
3. **Run:**
   
   Once installed, you can run the application using its executable name (ensuring $GOPATH is in your path).


   ```bash
   tree
   ```

## Options
| Flag  | Description             | Param Type | Default Value |
|-------|-------------------------|------------|---------------|
| -d    | Select directory        | String     | Current Dir   |
| -h    | Show hidden files       | Boolean    | False         |


Note. I just made this as a bit of fun.
