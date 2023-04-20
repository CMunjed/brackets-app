# Sprint 4.md

## Video Link
[Link to Video]()

## Work Done in Sprint 4
Front End:
During this sprint, features of the bracket were added and polished up, integration with the back end was done, and the footer was rewrote to adjust position based on how much content is on screen and keep its size static. Additionally, work was done to add a double elimination system of bracket generation and it is functional, although it isn't quite where we would like it to be at this point. UI elements were added and modified to better support the new integration with the backend to make a functioning login system. Finally, unit testing through cypress was successful in testing new funcionality, such as the sign in button and double elimination slider.

Back End:
During this sprint, we implemented user sessions and user permisions. We also redid our initialization function enable CORS permissions, which allowed us to integrate with the frontend. Users sessions work through cookies. Once a user logs in, they recieve a cookie that represents their current session. When interacting with brackets, this cookie will be passed to the backend to validate what user they are. User permissions also work through cookies. Brackets now store a whitelist of users which can view or edit the bracket depending on what settings the creator of the bracket sets. Only the creator of the bracket can delete the bracket. An admin flag has also been added, which was intended to block regular users from calling certain API calls, however, due to time constraints, this functionally has not been implemented at this time.

## Front-End Unit Tests

open web application on local host
compounding test, click on add teams button and edit teams dropdown
ensure existence of and test click on google sign in button
ensure existence of and test click on other sign in button
ensure existence of and test click on github button
test slide functionality
test slide functionality through multiple uses

## Back-end Unit Tests



## Updated Back-end Documentation

https://github.com/RetroSpaceMan123/brackets-app/wiki
