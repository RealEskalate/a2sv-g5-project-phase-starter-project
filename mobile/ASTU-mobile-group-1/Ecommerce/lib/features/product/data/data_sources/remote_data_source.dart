import 'dart:convert';
import 'dart:developer';
import 'dart:io';
import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../models/product_model.dart';
import 'local_data_source.dart';

abstract class ProductRemoteDataSource {
  Future<List<ProductModel>> getAllProducts();
  Future<List<ProductModel>> getProductsByCategory(String category);

  Future<ProductModel> getProductById(String id);

  Future<bool> insertProduct(ProductModel product);

  Future<bool> updateProduct(ProductModel product);

  Future<bool> deleteProduct(String id);
}

class ProductRemoteDataSourceImpl extends ProductRemoteDataSource {
  final http.Client client;
  final ProductLocalDataSource productLocalDataSource;

  ProductRemoteDataSourceImpl({
    required this.client,
    required this.productLocalDataSource,
  });

  @override
  Future<List<ProductModel>> getAllProducts() async {
    try {
      final token = await productLocalDataSource.getToken();

      final response = await client.get(
        Uri.parse(
          Urls.getAllProducts(),
        ),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );
      final Map<String, dynamic> jsonResponse = json.decode(response.body);

      if (response.statusCode == 200) {
        final List<dynamic> productData = jsonResponse['data'];
        final List<ProductModel> products =
            productData.map((json) => ProductModel.fromJson(json)).toList();

        await productLocalDataSource.cacheProducts(products);
        print(products);
        return products;
      } else {
        throw ServerException();
      }
    } on CacheException {
      throw CacheException();
    } catch (e) {
      throw ServerException();
    }
  }

  @override
  Future<ProductModel> getProductById(String id) async {
    try {
      final token = await productLocalDataSource.getToken();
      final response = await client.get(
        Uri.parse(
          Urls.getProductById(id),
        ),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        final Map<String, dynamic> jsonResponse = json.decode(response.body);

        final productData = jsonResponse['data'];
        return ProductModel.fromJson(
          productData,
        );
      } else {
        throw ServerException();
      }
    } catch (_) {
      throw ServerException();
    }
  }

  @override
  Future<bool> deleteProduct(String id) async {
    final token = await productLocalDataSource.getToken();

    final response = await client.delete(
      Uri.parse(
        Urls.getProductById(id),
      ),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
    );

    log(response.body.toString());
    if (response.statusCode == 200) {
      return true;
    } else {
      throw ServerException();
    }
  }

  @override
  Future<List<ProductModel>> getProductsByCategory(String category) {
    throw UnimplementedError();
  }

  @override
  Future<bool> insertProduct(ProductModel product) async {
    try {
      final uri = Uri.parse(Urls.getAllProducts());
      final request = http.MultipartRequest('POST', uri);

      final imageFile = File(product.imageUrl);
      if (!imageFile.existsSync()) {
        throw ServerException(message: 'Image file does not exist');
      }

      request.files.add(
        await http.MultipartFile.fromPath(
          'image',
          imageFile.path,
          contentType: MediaType('image', 'png'),
        ),
      );

      request.fields['name'] = product.name;
      request.fields['description'] = product.description;
      request.fields['price'] = product.price.toString();

      final token = await productLocalDataSource.getToken();
      request.headers['Content-Type'] = 'application/json';
      request.headers['Authorization'] = 'Bearer $token';

      final response = await request.send();

      log('Response reason phrase: ${response.reasonPhrase}');

      if (response.statusCode == 201) {
        log('Product uploaded successfully');
        return true;
      } else {
        log('Failed to upload product. Status code: ${response.statusCode}');

        throw Exception();
      }
    } on ServerException catch (e) {
      log(e.message);
      rethrow;
    } catch (e) {
      log(e.toString());
      throw ServerException();
    }
  }

  @override
  Future<bool> updateProduct(ProductModel product) async {
    try {
      final uri = Uri.parse(Urls.getProductById(product.id));
      final body = jsonEncode({
        'name': product.name,
        'description': product.description,
        'price': product.price,
      });

      final token = await productLocalDataSource.getToken();
      final response = await client.put(
        uri,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: body,
      );

      if (response.statusCode == 200) {
        return true;
      } else {
        return false;
      }
    } catch (e) {
      throw ServerException();
    }
  }
}
