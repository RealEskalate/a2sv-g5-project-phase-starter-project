import 'dart:convert';
import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';

import '../../../../core/errors/failure.dart';
import '../../../auth/data/data_source/auth_local_datasource.dart';
import '../models/product_model.dart';

abstract class ProductRemoteDatasource {
  Future<Either<Failure, List<ProductModel>>> getProducts();
  Future<Either<Failure, ProductModel>> getProductById(String id);
  Future<Either<Failure, ProductModel>> createProduct(ProductModel product,
      {File? file});
  Future<Either<Failure, ProductModel>> updateProduct(ProductModel product);
  Future<Either<Failure, void>> deleteProduct(String id);
}

class ProductRemoteDatasourceImp implements ProductRemoteDatasource {
  final AuthLocalDataSource localDataSource;
  final http.Client client;
  final String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2/products';

  ProductRemoteDatasourceImp(
      {required this.localDataSource, required this.client});

  @override
  Future<Either<Failure, ProductModel>> createProduct(ProductModel product,
      {File? file}) async {
    try {
      // Load the asset as bytes (you can replace this with actual image picking logic)
      File imageFile = File(product.imageUrl);
      List<int> imageBytes = await imageFile.readAsBytes();

      var request = http.MultipartRequest(
        'POST',
        Uri.parse(baseUrl), // Assuming this is the correct endpoint
      );
      final authHeaders = await _getAuthHeader();
      request.headers.addAll(authHeaders);

      request.fields.addAll({
        'name': product.name,
        'description': product.description,
        'price': product.price.toString(), // Ensure price is sent as a string
      });

      // Convert the bytes to MultipartFile and add to the request
      request.files.add(http.MultipartFile.fromBytes(
        'image',
        imageBytes,
        filename: 'fg.jpg',
        contentType: MediaType('image', 'jpeg'),
      ));

      http.StreamedResponse response = await request.send();

      if (response.statusCode == 201) {
        final responseBody = await response.stream.bytesToString();
        final Map<String, dynamic> data = json.decode(responseBody)['data'];

        final ProductModel newProduct = ProductModel.fromJson(data);
        return Right(newProduct);
      } else {
        return Left(
            ServerFailure('Failed to add product: ${response.reasonPhrase}'));
      }
    } catch (e) {
      return Left(ServerFailure('Error during addProduct: $e'));
    }
  }

  @override
  Future<Either<Failure, void>> deleteProduct(String id) async {
    try {
      final url = Uri.parse('$baseUrl/$id');
      final response =
          await client.delete(url, headers: await _getAuthHeader());
      print('suuuuuuuuu ${response.statusCode} and ${id}');
      if (response.statusCode == 204 || response.statusCode == 200) {
        // 204 No Content status for successful delete
        return const Right(null);
      } else {
        return Left(ServerFailure(
            'Failed to delete product. Status code: ${response.statusCode}'));
      }
    } catch (e) {
      return Left(ServerFailure('An error occurred: ${e.toString()}'));
    }
  }

  @override
  Future<Either<Failure, ProductModel>> getProductById(String name) async {
    final Either<Failure, List<ProductModel>> products = await getProducts();
      String productId = '';
      for(var prod in products.getOrElse(() => [])){
          if(prod.name == name){
          productId = prod.id;
          break;
        }
      }
      if (productId == '')
      {
        print("prduct not found");
      }

    try {
      final url = Uri.parse('$baseUrl/$productId');
      final response = await client.get(url,
      headers: await _getAuthHeader()
      );

      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        return Right(ProductModel.fromJson(responseData['data']));
      } else if (response.statusCode == 404) {
        return const Left(ServerFailure('Product not found'));
      } else {
        return Left(ServerFailure(
            'Failed to fetch product. Status code: ${response.statusCode}'));
      }
    } catch (e) {
      return Left(ServerFailure('An error occurred: ${e.toString()}'));
    }
  }

  @override
  Future<Either<Failure, List<ProductModel>>> getProducts() async {
    try {
      final url = Uri.parse(baseUrl);
      final response = await client.get(url,
          headers: await _getAuthHeader()
      );
      print('${response.statusCode} suuuuuuuuu');
      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        final List<ProductModel> products = (responseData['data'] as List)
            .map((json) => ProductModel.fromJson(json))
            .toList();
        return Right(products);
      } else {
        return Left(ServerFailure(
            'Failed to fetch products. Status code: ${response.statusCode}'));
      }
    } catch (e) {
      return Left(ServerFailure('An error occurred: ${e.toString()}'));
    }
  }

  @override
  Future<Either<Failure, ProductModel>> updateProduct(
      ProductModel product) async {
    try {
      final url = Uri.parse('$baseUrl/${product.id}');
      final response = await client.put(
        url,
        headers: {'Content-Type': 'application/json', ...await _getAuthHeader()},
        body: json.encode(product.toJson()),
      );

      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        return Right(ProductModel.fromJson(responseData['data']));
      } else if (response.statusCode == 404) {
        return Left(ServerFailure('Product not found'));
      } else {
        return Left(ServerFailure(
            'Failed to update product. Status code: ${response.statusCode}'));
      }
    } catch (e) {
      return Left(ServerFailure('An error occurred: ${e.toString()}'));
    }
  }

  Future<Map<String, String>> _getAuthHeader() async {
    final token = await localDataSource.getAccessToken();
    final headers = <String, String>{};
    if (token != null) {
      headers['Authorization'] = 'Bearer $token';
    }
    return headers;
  }
}
