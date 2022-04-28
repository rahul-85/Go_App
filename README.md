# Go_App
This is an API build up to fetch for covid data in Echo Golang

## /data API fetches the covid data and stores it in database
For e.g:
GET http://localhost:1323/data
### Output:
Data loaded Successfully!


## /data/location?lat={lat}&long={long} fetches the data from the database as per latitude and longitude entered by the user
For e.g:
GET http://localhost:1323/data/location?lat=28.5707841&long=77.3271074

### Output:
[{"regional_id":"UP","covidCases":1710158,"last_updated":"2021-11-01T09:54:14+05:30"},{"regional_id":"IND","covidCases":34285612,"last_updated":"2021-11-01T11:20:24+05:30"}]
