# Tesla mock API

Deployed on Vercel Now

Currently supporting 4 endpoints:

- POST `/oauth/token` with appropriate parameters
- GET `/api/1/vehicles/<any_id>/data` with appropriate auth header
- POST `/api/1/vehicles/<any_id>/command/door_unlock` with appropriate auth header
- POST `/api/1/vehicles/<any_id>/command/door_lock` with appropriate auth header