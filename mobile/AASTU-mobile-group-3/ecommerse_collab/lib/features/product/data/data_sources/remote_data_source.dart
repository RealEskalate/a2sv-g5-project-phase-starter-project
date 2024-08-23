import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';
import 'package:shared_preferences/shared_preferences.dart';


import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../models/product_model.dart';

abstract class ProductRemoteDataSource {
  Future<void> addProduct(ProductModel productModel);
  Future<List<ProductModel>> getAllProduct();
  Future<ProductModel> getProduct(String id);
  Future<void> deleteProduct(String id);
  Future<ProductModel> updateProduct(
      {required String id,
      required String name,
      required int price,
      required String description});
}

class ProductRemoteDataSourceImpl extends ProductRemoteDataSource {
  final http.Client client;
  ProductRemoteDataSourceImpl({required this.client});
  

  @override
  Future<List<ProductModel>> getAllProduct() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
      var headers = {
        'Authorization': "Bearer ${prefs.getString('token')}",
        'Content-Type': 'application/json',
         };
    
    final response = await client.get(Uri.parse(Urls.baseUrl), headers: headers);
    print("from remote data source ${response.body}");
    if (response.statusCode == 200) {
      final jsonData = json.decode(response.body) as Map<String, dynamic>;
      
      List<ProductModel> products = [];
      for (var data in jsonData['data']) {
        products.add(ProductModel.fromJson(data));
      }
      
print(response.statusCode);
      return products;
    } else {
      throw ServerException();
    }
  }

  @override
  Future<ProductModel> getProduct(String id) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
      var headers = {
        'Authorization': "Bearer ${prefs.getString('token')}",
        'Content-Type': 'application/json',
         };
    final response = await client.get(Uri.parse('${Urls.baseUrl}/$id'), headers: headers);

    if (response.statusCode == 200) {
      final jsonData = json.decode(response.body) as Map<String, dynamic>;

      ProductModel productModel = ProductModel.fromJson(jsonData['data']);

      return productModel;
    } else {
      throw ServerException();
    }
  }

  @override
  Future<void> deleteProduct(String id) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
      var headers = {
        'Authorization': "Bearer ${prefs.getString('token')}",
        'Content-Type': 'application/json',
         };
    final response = await client.delete(Uri.parse('${Urls.baseUrl}/$id'),  headers: headers);

    if (response.statusCode != 200) {
      throw ServerException();
    }
  }

  @override
  Future<ProductModel> updateProduct(
      {required String id,
      required String name,
      required int price,
      required String description}) async {
        SharedPreferences prefs = await SharedPreferences.getInstance();
      var headers = {
        'Authorization': "Bearer ${prefs.getString('token')}",
        'Content-Type': 'application/json',
         };
    final response = await client.put(Uri.parse('${Urls.baseUrl}/$id'),
        headers: headers,
        body: json.encode(
            {'name': name, 'description': description, 'price': price}));


  if (response.statusCode == 200) {
    final jsonData = json.decode(response.body) as Map<String, dynamic>;
    final productData = jsonData['data'];
    return ProductModel.fromJson(productData);
  } else if (response.statusCode == 404) {
    throw Exception('Product with ID $id not found');
  } else {
    throw ServerException();
  }
  }

  @override
  Future<void> addProduct(ProductModel productModel) async {
    // ByteData byteData = await rootBundle.load(productModel.image);
    // List<int> imageBytes = byteData.buffer.asUint8List();
SharedPreferences prefs = await SharedPreferences.getInstance();
      var headers = {
        'Authorization': "Bearer ${prefs.getString('token')}",
        'Content-Type': 'application/json',
         };

    var request = http.MultipartRequest('POST', Uri.parse(Urls.baseUrl), );
    request.headers.addAll(headers);
    request.fields.addAll({
        'name' : productModel.name,
        'description' : productModel.description,
        'price' : productModel.price.toString(),
    }
    );
    
    request.files.add(await http.MultipartFile.fromPath('image', productModel.image, contentType: MediaType('image', 'jpeg')),  );
    http.StreamedResponse response = await request.send();

    if (response.statusCode != 201){
      
        throw ServerException();
    }
    // print("mihret $response");
  }
}
