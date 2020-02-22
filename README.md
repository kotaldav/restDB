# Rest api for DB

!!UNDER DEVELOPMENT!!

## Usage

```
/api/v1/databases             # List of databases
/api/v1/{database}/tables     # List of tables in specified database

## Retrieve data - GET request
/api/v1/{database}/{table}    # Return json representation of table content
/api/v1/{database}/{table}?{column}=value
/api/v1/{database}/{table}?orderBy={column}
/api/v1/{database}/{table}?orderByDesc={column}

## Insert data - POST request
/api/v1/{database}/{table}    # Insert new row with json body content

## Replace data - PUT request
/api/v1/{database}{table}?{column}=value # Replace row with json body Content

## Update row - PATCH request
/api/v1/{database}{table}?{column}=value&{col1}={val1}&{col2}={val2}

## delete row - DELETE request
/api/v1/{database}{table}?{colum}={value}

```
