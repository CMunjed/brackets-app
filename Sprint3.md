# Sprint 3.md

## Video Link
[Link to Video]

## Work Done in Sprint 3
Front End:
During this sprint, much of the actual bracket creation was completed on the front end with a working system in place to create brackets with a user-entered number of teams. The user is also able to rename teams, change the title of the bracket, and click on the selected team to advance to the next stage. Work still needs to be done to integrate this with the back end and polish features, but much of the progress needed for this project was completed in this stage. Additionally, unit testing done through Cypress has provided to be a more versatile tool for what we're testing than using Karma, so the unit tests in this sprint were done entirely through Cypress.

Back End:
During this sprint, we added Google sign-up and sign-in and completely overhauled the existing user system to seamlessly integrate it with Google authentication. This included changing the means of identifying users in the code from a username-based approach to a mixed UUID and email-based approach. We also updated the routing system, brackets, unit tests, and everything else that worked with the username-based approach accordingly. Lastly, we created a better unit testing system using Go's included functionality rather than Postman or CURL commands, created a wiki (https://github.com/RetroSpaceMan123/brackets-app/wiki) as an easier method of sharing information between the front-end and back-end, and also began working on cookies.

## Front-End Unit Tests
create and run web application
ensure existence of and click on sign in component
ensure existence of and click on add teams component, and subsequent creation and clicking of edit team selection
ensure existence of and click on github star button

looking to add E2E tests to enter certain number of teams and click through created bracket, but text entry has been a pain so far


## Back-end Unit Tests

curl -X POST -H 'Content-Type:application/json' -d '{"iss":"test","nbf":123456789,"aud":"test","sub":"test","email":"test@email.com","email_verified":true,"azp":"test","name":"test","picture":"test","given_name":"test","iat":123456789,"exp":123456789,"jti":"test"}' localhost:3000/users/googlesignup

curl localhost:3000/users

curl -X PUT -H 'Content-Type:application/json' -d '{"iss":"test","nbf":123456789,"aud":"test","sub":"test","email":"test@email.com","email_verified":true,"azp":"test","name":"test","picture":"test","given_name":"test","iat":123456789,"exp":123456789,"jti":"test"}' localhost:3000/users/googlesignin

## Updated Back-end Documentation

https://github.com/RetroSpaceMan123/brackets-app/wiki
