# Sprint1.md

## User Stories
- As a tournament manager, I want to be able to dynamically add teams to a bracket so that I can organize matches more quickly.

- As a frequent user and bracket creator, I want to be able to save, manage, and make changes to existing brackets.

- As a sports fan, I want to be able to visualize the progression of tournaments and make predictions by creating brackets to represent them.

- As a member of a group that needs to create brackets, I want to be able to create brackets that can be shared with other members so that they can also edit them.

- As a party host, I want an easy, lightweight way to keep track of scores and matches for the mini-games I planned.

- As an event organizer, I want to automatically assign spots/pairings in a bracket based on provided criteria to create balanced and fair matches.

- As a sports predictor, I would like to create a bracket based on an ongoing tournament, and display my prediction for what will happen in the tournament for others to see.

- As a college student, I want to have a simple way to keep track of national championship games that need to be played.

- As a baseball tournament host, I want to be able to generate randomly seeded brackets for teams to play.


## What issues did your team plan to address?
Front end:
- We wanted to get the website running through the angular framework and to begin designing some elements of our project. Our goal is to get a basic design idea for our bracket, which displays the location and an example of the future functionality that will be implemented. We also want to add functionality with some buttons and a table that stores user-entered stats, along with possibly implementing some bracket functionality.

Back-end:
- In this sprint, we wanted to install the basic packages/technologies needed to create a basic functioning back-end and figure out how to implement a basic database which would eventually be used to store user authentication data. We planned to install Gorilla Mux, Gorm, and SQLite, and implement a basic database with data that could be manipulated.


## Which ones were successfully completed?
Front end:
- Got the site up and running on our local machines, added text and highlight cards in certain areas for descriptions and design purposes.
- We created a couple of buttons that function and will be linked to the future bracket we haven’t generated yet.
- A static data table was created to hold stats, which can be updated to have an update/refresh functionality in a future sprint.
- Wireframes for bracket designs were drawn up with our design in mind.

Backend:
- Implemented gorm and gorilla mux into our project.
- Created a basic sqlite3 database into our project.
- We linked our backend server to the database and were able to manipulate it through HTTP requests.
- We were able to run our backend server and make HTTP requests using curl, including requests to create, remove, edit, and view data entries, successfully giving the expected output from each function.


## Which ones weren't and why?
Front end:
- We failed to get a bracket to generate this attempt. Due to the unfamiliarity with angular it was a lot of experimenting and getting to learn how the framework works. We tried to install different barebones npm packages that utilize types of bracket trees but kept running into errors installing packages, importing the necessary file/required items, and licensing of certain packages. Functionality of bracket generation would have been nice to have at this point, but getting the features planned and mapped out was the main priority and was achieved.

Back-end:
- We didn’t really fail to accomplish any of the things we wanted to accomplish, but we ran into issues manipulating data entries in the database, which were eventually resolved. We didn’t create user objects to be stored in the database, but rather, we made employee objects, which will eventually be changed. 
- Although we have achieved all of our goals for this sprint, we have failed to connect our work to the front end. Even if we were able to, the API created was not compatible with the goals of the front-end, since we used dummy data and structures rather than any relevant functionality. 

