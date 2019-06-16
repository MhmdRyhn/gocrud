Creating Author:
================
```
{
    "on": "author",
    "apply": "create",
    "name": "Daron Acemoglu",
    "email": "daron@mail.com",
    "phone": "+123456789",
    "age": 45,
    "address": "Turkey"
}


{
    "on": "author",
    "apply": "create",
    "name": "Joe Navarro",
    "email": "joe@mail.com",
    "phone": "+102456489",
    "age": 52,
    "address": "USA"
}


{
    "on": "author",
    "apply": "create",
    "name": "Robert C. Martin",
    "email": "martin@mail.com",
    "phone": "+1234034728",
    "age": 42,
    "address": "USA"
}
```


Get all author
===============
```
{
    "on": "author",
    "apply": "fetch_all"
}
```


Get single author
==================
```
{
    "on": "author",
    "apply": "fetch_one",
    "email": "daron@mail.com"
}
```


Update Author
==============
```
{
    "on": "author",
    "apply": "update",
    "email": "daron@mail.com",
    "name": "Daron Acemogluuuu"
}
```


Delete Author
==============
```
{
    "on": "author",
    "apply": "delete",
    "email": "daron@mail.com"
}
```











.
