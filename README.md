# gin-mall


## Bugs

### 88001 
- ListProducts : can not get products in different pages using `page_size` and `page_num`

### 8802
- SearchProducts : can not get products by info searching  `info`

### 8803
 - DeleteFavorite : ``[1.480ms] [rows:0] UPDATE `favorite` SET `deleted_at`='2022-10-21 15:55:21.15' WHERE (id=2 AND user_id=1) AND `favorite`.`deleted_at` IS NULL``
 - ShowFavorites : still showing the data after deleted`"status": 200,
   "data": {
   "item": [
   {
   "user_id": 2,
   "product_id": 11,
   "create_at": 1666338648,
   "name": "clothes",
   "category_id": 1,
   "title": "man t-shirt",
   "info": "best choise for young man",
   "img_path": "127.0.0.1:3000/static/imgs/product/boss1/clotches.jpg",
   "price": "100",
   "discount_price": "99",
   "boss_id": 2,
   "num": 10,
   "on_sale": true
   }
   ],
   "total": 1
   },
   "msg": "ok",
   "error": ""
   }`