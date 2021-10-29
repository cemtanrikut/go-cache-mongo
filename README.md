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
 Fetch all data from db

 https://go-cache-mongo.herokuapp.com/mongodb/fetch  
 localhost/mongodb/fetch

- ### **GET/{Key}**  
 Require Key item. And set the cache

 https://go-cache-mongo.herokuapp.com/get/1  
 localhost/get/TAKwGc6Jr4i8Z487  





