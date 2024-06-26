
# Installing Tree

![image](https://github.com/robert-ohurley/tree/assets/96722504/9a3d2b8a-54a6-4dc0-b6b8-186c291d4f69)

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
| Flag              | Description                        | Param Type | Default Value |
|-------------------|------------------------------------|------------|---------------|
| -d                | Select directory                   | String     | Current Dir   |
| -h                | Show hidden files                  | Boolean    | False         |
| -D                | Depth to traverse                  | Integer    | 10            |
| -fullpath         | Show full path name for directories| Boolean    | false         |
| -separator        | Character used to separate lines   | String     | \|            |
| --file-pointer    | String used to point to files      | String     | \|\--         |
| --directory-color | Color to print directories         | String     | blue          |
| --file-color      | Color to print files               | String     | white         |
| --dir-only        | Only print directories             | Boolean    | false         |

   
   Alias your preferred settings by executing, for example, 


   ```bash
   echo "alias tree='tree -h=true -fullpath=true --directory-color=cyan'" >> ~/.bashrc
   ```
   
Note. I just made this as a bit of fun.







