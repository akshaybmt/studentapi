# Student-api

First create a DB with name 'crud'

## Custom Port Facility

 User can pass the custom port facility by passing its value in the flag during the runtime call as " **go run \.main.go -port 8070**" .(Here I used 8070 as custom port)
 Otherwise it will open the application on **Port 8080 by default**.


# Available Endpoints

POST   /api/v1/student                  
GET    /api/v1/student                 
GET    /api/v1/student/:id              
PUT    /api/v1/student/:id              
DELETE /api/v1/student/:id                                                                                                                                                                 
POST   /api/v1/signup                                                                                                                                                                     
POST   /api/v1/login



### Validation

Email is checked for its validness.

Phone number is checked for its valid 10-digit length.

Name is checked if it is not empty.


### Sorting 

Sorting is defined in the URL by using string` "sort_field= field_name"  & "dir = direction_of_sorting" `

For example` "/student?sort_field=phone&sort_dir=asc" `(Here we have choosen Column field as phone and Direction for Sorting as Ascending)


 ### Pagination
 
 Pagination is added on GET request of the API, it defines the custom limit and pages on request.
 
 ### Check For Existing Data (Phone Number or Email)
 
 In order to avoid duplicacy of Data, the API will check for Existing Data with similar Phone number and Email Id and will prevent from creating duplicate entries.
 
 
 ### Data Filtering 
 
 User can now filter the data from the tables by Passing Search string as` "search= search_key " `(Here search_key is the value to be searched) in the URL. 
 It will filter all the data that matches the passed string. 
 Note:- It is not necessary to pass the full length of the search key, the API will search for all the data fields that matches the passed string wheather it is partial or full.
 
 ### Time Based Filtering
 
 User can now filter data on daily/weekly/monthly/yearly basis
 User has to pass the Filter String as ` "time= time_key"  `(Here time_key can be daily/weekly/monthly/yearly) in the URL.
 
 For Example = ` "\student?time=daily" ` (Here we are filtering based on data created in last 24hours or Daily Basis)

 ### Date wise Filtering 

 User can now filter the data from the data based on the Date Created at and Date Updated at .
 User has to pass the Filter String as ` "date= date_key" ` (Here date_key can be the date value) in the URL.

For Example = ` "/student?date=2021-08-08" ` (Here we are filtering the data that was either created on 8th August or updated on 8th August)


### SignUp and Login Feature

SignUp feature will allow user to register themselves as a user to access the database.        
Validation is available at the time of User SignUp.                                                            
Login feature will allow user to login using the existing credentials inside the database.    


### Generation of JWT Token

A JWT token is generated at the time of User Login.


### Grouping of EndPoints

Endpoints of the API are grouped in the Route using ` .Group ` function to create a new router Group.
(Here we have created v1 as our router group for Crud Functions (`/api/v1`) and v2 for SignUp and Login(`/api/v2`))
