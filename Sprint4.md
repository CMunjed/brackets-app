# Sprint 4.md

## Video Link
[Link to Video]()

## Work Done in Sprint 4
Front End:


Back End:
During this sprint, we implemented user sessions and user permisions. We also redid our initialization function enable CORS permissions, which allowed us to integrate with the frontend. Users sessions work through cookies. Once a user logs in, they recieve a cookie that represents thier current session. When interacting with brackets, this cookie will be passed to the backend to validate what user they are. User permissions also work through cookies. Brackets now store a whitelist of users which can view or edit the bracket depending on what settings the creator of the bracket sets. Only the creator of the bracket can delete the bracket. An admin flag has also been added, which was intended to block regular users from calling certain API calls, however, due to time constraints, this functionally has not been implemented at this time.

## Front-End Unit Tests



## Back-end Unit Tests



## Updated Back-end Documentation

https://github.com/RetroSpaceMan123/brackets-app/wiki
