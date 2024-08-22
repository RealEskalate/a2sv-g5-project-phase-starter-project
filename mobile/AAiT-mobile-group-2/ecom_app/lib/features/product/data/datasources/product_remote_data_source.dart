import 'dart:convert';
import 'dart:io';

import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/network/custom_client.dart';
import '../models/product_model.dart';

abstract class ProductRemoteDataSource {
  Future<ProductModel> getCurrentProduct(String id);
  Future<List<ProductModel>> getAllProducts();
  Future<ProductModel> createProduct(ProductModel product);
  Future<ProductModel> updateProduct(ProductModel product);
  Future<void> deleteProduct(String id);
}

class ProductRemoteDataSourceImpl extends ProductRemoteDataSource {
  // final http.Client client;
  final CustomHttpClient client;

  ProductRemoteDataSourceImpl({required this.client});

  @override
  Future<ProductModel> createProduct(ProductModel product) async {
    var uri = Uri.parse(Urls.baseUrl2);
    var request = http.MultipartRequest('POST', uri);
    // request.headers['Content-Type'] = 'multipart/form-data';
    request.fields['name'] = product.name;
    request.fields['description'] = product.description;
    request.fields['price'] = product.price.toString();

    if (product.imageUrl.isNotEmpty) {
      var imageFile = File(product.imageUrl);
      debugPrint(
          'created image: ${imageFile.existsSync()}, ${product.imageUrl}');
      if (imageFile.existsSync()) {
        request.files
            .add(await http.MultipartFile.fromPath('image', product.imageUrl, contentType: MediaType('image', 'jpg',)));
      } else {
        throw ImageException();
      }
    }

    try {

      http.StreamedResponse streamedResponse = await client.send(request);

      if (streamedResponse.statusCode == 201) {
        final responseString = await streamedResponse.stream.bytesToString();
        final jsonResponse = json.decode(responseString);
        return ProductModel.fromJson(jsonResponse['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException(ErrorMessages.socketError);
    } catch (e) {
      throw Exception(e.toString());
    }
  }

  @override
  Future<void> deleteProduct(String id) async {
    try {
      final response =
          await client.delete(Urls.currentProductById(id));
      if (response.statusCode == 200) {
        return;
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException(ErrorMessages.socketError);
    }
  }

  @override
  Future<List<ProductModel>> getAllProducts() async {
    try {
      final response = await client.get((Urls.baseUrl2));
      if (response.statusCode == 200) {
        // print(response.body);
        return ProductModel.fromJsonList(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException(ErrorMessages.socketError);
    }
  }

  @override
  Future<ProductModel> getCurrentProduct(String id) async {
    try {
      final response = await client.get((Urls.currentProductById(id)));
      if (response.statusCode == 200) {
        return ProductModel.fromJson(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException(ErrorMessages.socketError);
    }
  }

  @override
  Future<ProductModel> updateProduct(ProductModel product) async {
    final productId = product.id;
    final jsonBody = jsonEncode({
      'name': product.name,
      'description': product.description,
      'price': product.price,
    });
    try {
      final response = await client.put(
          (Urls.currentProductById(productId)),
          body: jsonBody,
          );
      if (response.statusCode == 200) {
        return ProductModel.fromJson(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException(ErrorMessages.socketError);
      // } catch (e) {
      //   print('------------------------------------------exception-- ${e.toString()}---------------------');
      //   throw Exception(e.toString());
    }
  }
}
