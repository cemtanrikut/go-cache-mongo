# go-cache-restapi

 This is an In memory key-value store database rest api project.  

 I use golang standard libs.


 This RestAPI deployed **HEROKU** same time. 
 https://go-cache-rest-api.herokuapp.com  

## Endpoints  
- ### **SET** 
 Sets key-value items. Require BODY 

 https://go-cache-mongo.herokuapp.com/set  
 localhost/set 

 Sample Json Body,  
```javascript
{
    "key": "active-tabs",
    "value": "getir"
}
```
    

- ### **/mongodb/fetch** 
 Fetch all data from db. Require BODY  

 Sample Json Body,  
```javascript
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```


 https://go-cache-mongo.herokuapp.com/mongodb/fetch  
 localhost/mongodb/fetch  

- ### **GET/{Key}**  
 Require Key item. And set the cache  

 https://go-cache-mongo.herokuapp.com/get/TAKwGc6Jr4i8Z487    
 localhost/get/TAKwGc6Jr4i8Z487  





