import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;
import '../../../../core/constants/constants.dart';
import '../../../../core/error/exceptions.dart';
import '../../domain/entitity/product.dart';
import '../model/product_model.dart';
import 'package:http_parser/http_parser.dart';

abstract class ProductRemoteDataSource {
  Future<Product> getProduct(String id);
  Future<List<Product>> getAllProduct();
  Future<Product> insertProduct(Product product);
  Future<Product> updateProduct(Product product);
  Future<Product> deleteProduct(String id);
}

class ProductRemoteDataSourceImpl implements ProductRemoteDataSource {
  final http.Client client;
  ProductRemoteDataSourceImpl({required this.client});
  @override
  Future<Product> getProduct(String id) async {
    final response = await client.get(Uri.parse(Urls.getProductById(id)));
    if (response.statusCode == 200) {
      var check = ProductModel.fromJson(json.decode(response.body));
      return check;
    } else {
      throw ServerException();
    }
  }

  @override
  Future<Product> deleteProduct(String id) async {
    final response = await client.delete(Uri.parse(Urls.getProductById(id)));
    if (response.statusCode == 200) {
      return ProductModel.fromJson(json.decode(response.body));
    } else {
      throw ServerException();
    }
  }

  @override
  Future<List<Product>> getAllProduct() async {
    final response = await client.get(Uri.parse(Urls.getAllProducts()));
    print('Response status: ${response.statusCode}');
    print('Response body: ${response.body}');

    if (response.statusCode == 200) {
      try {
        List<dynamic> list = json.decode(response.body)['data'];
        print('Parsed data: $list'); // Debug output to verify parsing
        return list.map((product) => ProductModel.fromJson(product)).toList();
      } catch (e) {
        print('Error parsing JSON: $e');
        throw ServerException(); // Handle parsing errors
      }
    } else {
      print('Server responded with status code: ${response.statusCode}');
      throw ServerException();
    }
  }

  @override
  Future<Product> insertProduct(Product product) async {
    var request =
        http.MultipartRequest('POST', Uri.parse(Urls.getAllProducts()));

    // Add text fields
    request.fields['name'] = product.name;
    request.fields['price'] = product.price.toString();
    request.fields['description'] = product.description;

    // Add image file if available
    if (product.imageUrl.isNotEmpty && File(product.imageUrl).existsSync()) {
      request.files.add(await http.MultipartFile.fromPath(
        'image',
        product.imageUrl,
        contentType: MediaType('image', 'jpeg'),
      ));
    }

    // Send the request
    var response = await request.send();
    if (response.statusCode == 201) {
      var responseBody = await response.stream.bytesToString();
      return ProductModel.fromJson(json.decode(responseBody)['data']);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<Product> updateProduct(Product product) async {
    final response = await client.put(
      Uri.parse(Urls.getProductById(product.id)),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(ProductModel.fromProduct(product).toJson()),
    );
    if (response.statusCode == 200) {
      return ProductModel.fromJson(json.decode(response.body)['data']);
    } else {
      throw ServerException();
    }
  }
}
