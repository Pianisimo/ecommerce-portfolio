# Demo e-commerce Project

This is a demo e-commerce project it uses:
- react as the frontend
- go as the backend
- postgres as the database
- nginx for serving the react application


## React 

The frontend part is really basic, it only connects to the backend for the authentication part
the products of the project are just hardcoded and the shopping cart is using redux-persistor to cache it

## GO

The backend handle the API calls from the frontend.
It checks if the user is authenticated via http-only cookies, if not it marks as not authenticated and users can
register or login to the page
it uses CSRF on top of JWT to handle authentication, checks cookie validity and compares it to the CSRF token

## Postgres

Just the chosen database to store the user info, token, refresh-token, and csrf-token.

# Deploy

to deploy and test just go to the root directory of the project and run

`docker-compose up`
