# Shopping API with GraphQL

----
## How to use this API
I recommend using [Postman](https://www.getpostman.com/), but querying through the browser is completely fine.

Currently deployed to [Heroku](https://its-a-shop.herokuapp.com)

Queries to the API look like
   
```
/graphql?query={}
```
Within the curly braces, describe the GraphQL schema desired. As I implemented this naively, there is no need for "query" and "mutation" semantics. The schemas implemented look as follows:

```
shop {
  id
  name
  products{}
  orders{}
}
product {
  id
  shop_id
  name
  value
  quantity
  line_items{}
}
order {
  id
  shop_id
  value
  line_items{}
}
line_item {
  id
  product_id
  order_id
  value
}
```
CRUD functionality is included:
```
shop(id,name): Gets and returns shops. Arguments are optional, and will return based on satisfying ALL conditions. No arguments will return all shops.
createShop(name): Creates a new shop with provided name (mandatory).
updateShop(id,name): Updates shop with provided ID with new name.
deleteShop(id): Deletes shop with given ID.

product(id,name,shop_id,value,quantity)
createProduct(shop_id,name,value,quantity): All fields mandatory
updateProduct(id,shop_id,name,value,quantity): Can update whichever fields you want.
deleteProduct(id)

order(id,name,shop_id)
createOrder(shop_id)
updateOrder(id,shop_id): You can change an order's shop, not sure why you would want to though.
deleteOrder(id)

lineItem(id)
createLineItem(product_id,order_id,quantity): Again, all fields mandatory.
updateLineItem(id,product_id,order_id,quantity)
deleteLineItem(id)
```
The CRUD "methods" need to be verbose and in paretheses '( )', where you specify parameter key and value. For example:
```
{product(name:"Wizard Hat"){id,shop_id,name}}
```
Now, all these queries/mutations must be followed by a return schema in more curly braces '{ }'. For example:
```
{shop{id,name}}
{updateShop(id:2,name:"Hallelujah"){id,name,products{id,shop_id,name,value}}
{deleteOrder(id:4){id}}
{createOrder(shop_id:5){id,shop_id}}
{order(id:3){value, list_items{quantity,value}}}
{lineItem{id,value,order_id,product_id}}
```
Here's what a query to the entire database looks like:
```
https://its-a-shop.herokuapp.com/graphql?query={shop{id,name,products{id,shop_id,name,value,quantity,line_items{id,product_id,order_id,quantity}}orders{id,value,line_items{id,product_id,order_id,value}}}}
```
### I didn't do a great job with naming consistency, so pay attention to the differences in the singular camelCase methods and the plural snake_case schema returns. Sorry!.

Play around, and if you delete too much and don't want to painstakingly recreate data, just hit the reseed endpoint:
```
https://its-a-shop.herokuapp.com/reseed
```
and that should reseed the data back to the defaults I've provided.

Now all of this has been written for GET queries, but POST queries are also supported. However, as I have not written the GraphQL implementation well, it's not really that much of an improved experience to be sending schemas as JSON bodies through POST. Sorry!
