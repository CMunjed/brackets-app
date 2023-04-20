# brackets-app
This is our brackets app for CEN3031 Intro to Software Engineering Spring 2023 @ UF.

## Project Description
This is an online tournament bracket creation tool with an emphasis on flexibility, customizability, and ease of use. <br /><br />

The goal is to create a product that is more user-friendly and intuitive than existing alternatives. For example, this tool will allow modification of existing brackets, a feature which is surprisingly scarce among alternatives. We plan to include a wide range of functionality, including a large selection of tournament styles such as single elimination, double elimination, round robin, etc. Other features will include the ability to input, change, and customize contestant information and bracket configurations, the ability to merge existing brackets, the ability to randomize/shuffle brackets, and more.<br /><br />

This tool will utilize Angular for the front-end and Go for the back-end.

## Project Members
Logan Bialek - Front-end <br />
Shawn Banks - Front-end <br />
Carlos Avila - Back-end <br />
Connor Munjed - Back-end <br />
CEN3031 - Group 74 <br />

## How To Run
1. Needed Software:
   - [Golang](https://go.dev/dl/)
   - [Node.js](https://nodejs.org/en/download)
   - GCC
     - If you are using a Windows machine, follow [this](https://code.visualstudio.com/docs/cpp/config-mingw) link to install GCC
     - If you are using a Mac machine, follow [this](http://cs.millersville.edu/~gzoppetti/InstallingGccMac.html) link to install GCC
     - If you are using a Linux machine, follow [this](https://www.geeksforgeeks.org/how-to-install-gcc-compiler-on-linux/) link to install GCC
    - [SQLite 3](https://sqlite.org/download.html)
    -Angular materials is needed as well for some functionality used(ex. buttons)
    

2. Clone the repository from the repository
3. Open a terminal in the project folder in your local machine, and run the following commands:
    - ```go mod tidy```
    - ```cd brackets-app```
    - ```npm update```
4. Once all of the dependencies have been installed, run the following commands in two separate terminals to start both frontend and backend servers starting from the project root directory
    - Backend Server  
      - ```go run .```
    - Frontend Server
      - ```cd brackets-app```
      - ```ng serve```
5. Go to [http://localhost:4200](http://localhost:4200) to see the webpage

6.How to run website

- Click on Insert Bracket Name to rename
- Click on How many teams text box to enter number of teams(ex. recommend use 8 shows enough matchups and not to many to progress)
- Select the slide button for type of bracket(ex. select the single elimination-it is easier to use)
- Once type of bracket selected click add teams to generate the bracket
- Once bracket is generated you can click the edit team name button to bring up menu and buttons to modify bracket
- Use the drop down to select team # you want to change and enter team names or organize matchups how you want 
- To use bracket for each matchup in the first column select one winner to progress to next match(finish all matchups in first till second ones are created)
- Once first column games winners are select just continue with each column to completion until you get to your winner
- And then your done 
