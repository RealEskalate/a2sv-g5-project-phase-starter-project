import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';

import '../../../../../core/constants/constants.dart';
import '../../../../../core/exception/exception.dart';
import '../../../../../core/network/custom_client.dart';
import '../../models/product_model.dart';

abstract class ProductRemoteDataSource {
  Future<List<ProductModel>> getProducts();
  Future<ProductModel> getProductById(String id);
  Future<ProductModel> createProduct(ProductModel product);
  Future<ProductModel> updateProduct(ProductModel product);
  Future<void> deleteProduct(String id);
}

class ProductRemoteDataSourceImpl extends ProductRemoteDataSource {
  // final http.Client client;
  final CustomHttpClient client;

  ProductRemoteDataSourceImpl({required this.client});

  @override
  Future<List<ProductModel>> getProducts() async {
    try {
      final response = await client.get(Urls.baseUrl);
      if (response.statusCode == 200) {
        return ProductModel.fromJsonList(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<ProductModel> getProductById(String id) async {
    try {
      final response = await client.get(Urls.getProdutbyId(id));
      if (response.statusCode == 200) {
        return ProductModel.fromJson(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<ProductModel> createProduct(ProductModel product) async {
    var request = http.MultipartRequest(
      'POST',
      Uri.parse(Urls.baseUrl),
    );
    request.fields.addAll({
      'name': product.name,
      'description': product.description,
      'price': product.price.toString(),
    });

    if (product.imageUrl.isNotEmpty) {
      var imageFile = File(product.imageUrl);
      if (imageFile.existsSync()) {
        request.files.add(
          await http.MultipartFile.fromPath('image', product.imageUrl,
              contentType: MediaType('image', 'jpeg')),
        );
      } else {
        throw ImageException();
      }
    }

    try {
      http.StreamedResponse response = await client.send(request);

      if (response.statusCode == 201) {
        final responseJson = await response.stream.bytesToString();
        return ProductModel.fromJson(json.decode(responseJson)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    } catch (e) {
      throw Exception(e.toString());
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
          Urls.getProdutbyId(productId),
          body: jsonBody,
        );
      if (response.statusCode == 200) {
        return ProductModel.fromJson(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<void> deleteProduct(String id) async {
    try {
      final response = await client.delete(Urls.getProdutbyId(id));
      if (response.statusCode == 200) {
        return;
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }
}
