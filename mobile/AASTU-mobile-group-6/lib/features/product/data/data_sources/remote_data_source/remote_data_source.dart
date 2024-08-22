import 'dart:convert';
import 'dart:typed_data';
import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/constants/constants.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:get/get.dart';
import  'package:http/http.dart' as http;
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:http_parser/http_parser.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../../core/errors/failure/failures.dart';
import '../../../domain/entities/product_entity.dart';
import 'dart:io';


abstract class ProductRemoteDataSource {

  Future <ProductModel> getProduct(String id);
  Future<String> addProduct(ProductEntity product);
  Future<String> deleteProduct(String id);
  Future<String> updateProduct(ProductModel product);
  Future<List<ProductModel>> getProducts();
  Future<UserModel> getUserInfo();
 
}

class ProductRemoteDataSourceImpl extends ProductRemoteDataSource {
  final http.Client client;
  ProductRemoteDataSourceImpl({required this.client});
  
  get global_access_token => null;
  
  @override
  Future<ProductModel> getProduct(String id) async {
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',
    };
    final response = await client.get(Uri.parse(Urls.getProduct(id)),headers: head);
    print(response);
    if (response.statusCode == 200) {
      var ans = jsonDecode(response.body);
      print(ans);
      return ProductModel.fromJson(json.decode(ans['data']));

    } else {
      throw Exception('Failed to load data');
    }
  }
  @override
  Future<UserModel> getUserInfo()async {
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v2/users/me';
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',
    };
    final response = await client.get(Uri.parse(url),headers: head);
    print(response);
    if (response.statusCode == 200) {
      var ans = jsonDecode(response.body)['data'];
      var user_data = UserModel(name: ans['name'], email: ans['email'], password: 'password');
      print(user_data);

      return user_data;

    } else {
      throw Exception('Failed to load user data');
    }
  }


  @override
  Future<List<ProductModel>> getProducts() async {
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v2/products';
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');

    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',
    };
    final response = await client.get(Uri.parse(url),headers: head);

    if (response.statusCode == 200) {
      return (jsonDecode(response.body)['data'] as List)
          .map((e) => ProductModel.fromJson(e))
          .toList();
    } else {
      throw Exception('Failed to load data');
    }
  }

  @override
  Future<String> addProduct(ProductEntity product) async {
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v2/products';
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token').toString();
    
   
    final request = http.MultipartRequest('POST', Uri.parse(url));
    request.fields['name']= product.name;
    request.fields['description']= product.description;
    request.fields['price'] = product.price.toString();
    


    
    request.files.add(await http.MultipartFile.fromPath('image', product.imagePath,contentType: MediaType("image","jpg")));

    
    
    request.headers['Authorization'] = 'Bearer $temp2';
    

    

    var response = await request.send();

    print(response.statusCode);

    if (response.statusCode == 201) {

      return "Added Successfully";
    } else {
      throw Exception('Failed to add data');
    }
  //  To be edited with the image picker if needed also the response has to be product entity changed using from json
  }

  Future<String> deleteProduct(String id) async {
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',      
    };
    final response = await client.delete(Uri.parse(Urls.deleteProduct(id)),headers: head);
    print(temp2);
    print(response.statusCode);
    if (response.statusCode == 200) {
      return 'Product deleted successfully';
    } else {
      throw Exception('Failed to delete product');
    }
  }

  @override
  Future<String> updateProduct(ProductModel product) async {
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v2/products/';
    var temp1 = await SharedPreferences.getInstance();
    var temp2 = temp1.getString('access_token');
    
    var temp = jsonEncode(product.toJson());
    

    
    

    print(temp);
    print(url);

    var headers = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',
    };

    var response = await client.put(Uri.parse(url+product.id),headers: headers,body: temp);

    print(response.statusCode);
    


    if (response.statusCode==200){
      return "Product Updated SUccesfully";
    }else{
      throw Exception('Error ' + response.statusCode.toString());
    }
    // the response has to be product entity changed using from json
  }


}