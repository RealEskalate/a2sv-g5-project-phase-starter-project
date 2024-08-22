import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../model/product_model.dart';

abstract class localDataSource{
 List<ProductModel> getallproduct();
Future<void> addproduct(ProductModel newproduct) ;
Future<void> cachedproduct(List<ProductModel> cachedproducts) ;
Future<ProductModel> updateproduct(ProductModel newproduct) ; 
 Future<bool> deleteproduct(String id) ;
 ProductModel getproduct(String id) ;

}

class localDataSourceImpl implements localDataSource{
  var keyname = 'cachedprods';
  SharedPreferences sharedPreference;
  localDataSourceImpl({required this.sharedPreference});
  @override
  Future<void> addproduct(ProductModel newproduct) {
   
    throw UnimplementedError();
  }

  @override
  Future<void> cachedproduct(List<ProductModel> cachedproducts)async {
    var mapped = cachedproducts.map((product) => product.toJson()).toList();
    var JsonMap = json.encode(mapped);
    bool response = await sharedPreference.setString(keyname, JsonMap);
    if (response==false){
      throw Exception('local error');
    }
  }

  @override
  Future<bool> deleteproduct(String id) async {
    var response = sharedPreference.getString(keyname);
    if(response!=null){ 
      var listofproducts = _jsonToProductList(response);
      var i =0;
      for(var prod in listofproducts){
        if (prod.id==id){
          listofproducts.remove(listofproducts[i]);
           await cachedproduct(listofproducts);
           return true;
        }
        i++;
      }
       throw Exception("Can't find product with this id");
      }else{
        throw Exception('error with shared preference');
      }
  }

  @override
  List<ProductModel> getallproduct()  {
    var response = sharedPreference.getString(keyname);
    if (response!=null){
      List<ProductModel> listOfCachedProducts = _jsonToProductList(response);
      return listOfCachedProducts;
    }else{
      throw Exception('cached missed')
;    }
  }

  List<ProductModel> _jsonToProductList(String response) {
     var listJson = json.decode(response);
    
    List<ProductModel> listOfCachedProducts = [];
    for(var li in listJson){
      listOfCachedProducts.add(ProductModel.fromJson(li));
    }
    return listOfCachedProducts;
  }

  @override
  ProductModel getproduct(String id) {
    var response = sharedPreference.getString(keyname);
    if(response!=null){ 
      var listofproducts = _jsonToProductList(response);
      for(var prod in listofproducts){
        if (prod.id==id){
          return prod;
        }
      }
      throw Exception('missed product');
      }else{
        throw Exception('error with shared preference');
      }
  }

  @override
  Future<ProductModel> updateproduct(ProductModel newproduct) async  {
    
    var response = sharedPreference.getString(keyname);
    if(response!=null){ 
      var listofproducts = _jsonToProductList(response);
      var i =0;
      for(var prod in listofproducts){
        if (prod.id==newproduct.id){
          listofproducts[i] = newproduct;
           await cachedproduct(listofproducts);
           return listofproducts[i];
        }
        i++;
      }
       throw Exception("Can't find product with this id");
      }else{
        throw Exception('error with shared preference');
      }
  }




}