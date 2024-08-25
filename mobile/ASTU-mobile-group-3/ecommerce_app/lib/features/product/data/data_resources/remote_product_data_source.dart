import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
// ignore: depend_on_referenced_packages
import 'package:http_parser/http_parser.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/errors/exceptions/product_exceptions.dart';
import '../../../auth/data/data_source/auth_local_data_source.dart';
import '../models/product_model.dart';

abstract class RemoteProductDataSource {
  Future<int> deleteProduct(String id);

  Future<List<ProductModel>> getAllProducts();

  Future<ProductModel> getProduct(String id);

  Future<int> insertProduct(ProductModel productModel);

  Future<int> updateProduct(ProductModel productModel);
}

class RemoteProductDataSourceImp implements RemoteProductDataSource {
  final http.Client client;
  final AuthLocalDataSource authLocalDataSource;
  Map<String, String> headerWithToken = Map.from(AppData.jsonHeader);

  RemoteProductDataSourceImp(this.client, this.authLocalDataSource);

  /// Send request to server to delete a given product by the provided [id]
  ///
  /// Throws [ServerException] error if the request failed
  @override
  Future<int> deleteProduct(String id) =>
      executeQuery(AppData.delete, '${AppData.allProductUrl}/$id');

  /// Sends request to server to get list of products from the server
  ///
  /// Throws [ServerException] if the request failed
  @override
  Future<List<ProductModel>> getAllProducts() =>
      getProductRefactor(AppData.allProductUrl);

  /// Sends Request to server to get single product by provided [id]
  ///
  /// Throws [ServerException] error if the request failed
  @override
  Future<ProductModel> getProduct(String id) async {
    List<ProductModel> result =
        await getProductRefactor(AppData.allProductUrl, id);
    return result[0];
  }

  /// Sends request to server to insert a given [productModel]
  ///
  /// Send image through
  ///
  /// Throws [ServerException] if the request failed
  @override
  Future<int> insertProduct(ProductModel productModel) async {
    final headerAuth = await authLocalDataSource.getToken();

    headerWithToken['Authorization'] = 'Bearer ${headerAuth.token}';
    try {
      final uri = Uri.parse(AppData.allProductUrl);
      final request = http.MultipartRequest('POST', uri);
      request.headers.addAll(headerWithToken);
      if (!File(productModel.imageUrl).existsSync()) {
        throw ServerException();
      }
      request.files.add(
        await http.MultipartFile.fromPath(
          'image',
          productModel.imageUrl,
          contentType: MediaType('image', 'png'),
        ),
      );
      request.fields['name'] = productModel.name;
      request.fields['description'] = productModel.description;
      request.fields['price'] = productModel.price.toString();

      final result = await client.send(request);

      if (result.statusCode == 201) {
        return AppData.successInsert;
      } else {
        debugPrint(result.statusCode.toString());

        throw ServerException();
      }
    } on Exception {
      throw ServerException();
    }
  }

  /// Sends request to the server to be updated [id] in the url and [productModel] as json
  ///
  ///
  /// Throws [ServerException] if request is failed
  @override
  Future<int> updateProduct(ProductModel productModel) => executeQuery(
        AppData.put,
        '${AppData.allProductUrl}/${productModel.id}',
        {
          'name': productModel.name,
          'description': productModel.description,
          'price': productModel.price,
        },
      );

  /// Methods used for refactoring that is employed by inser, update, delete methods
  ///
  /// Throws an exceptions[ServerException] which is later also thrown by the calling methods
  Future<int> executeQuery(String requestType, String url,
      [Map<String, dynamic>? data]) async {
    final headerAuth = await authLocalDataSource.getToken();

    headerWithToken['Authorization'] = 'Bearer ${headerAuth.token}';
    Map<String, dynamic> typeMap = {
      AppData.post: client.post,
      AppData.get: client.get,
      AppData.delete: client.delete,
      AppData.put: client.put
    };

    try {
      late http.Response result;
      if (data != null) {
        result = await typeMap[requestType](Uri.parse(url),
            body: json.encode(data), headers: headerWithToken);
      } else {
        result = await typeMap[requestType](Uri.parse(url),
            headers: headerWithToken);
      }
      if (result.statusCode == 200) {
        return AppData.getCorrespondingSuccess(requestType);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw ServerException();
    } on Exception {
      rethrow;
    }
  }

  /// Refactor done for data retrival based in the existance of [id]
  ///
  /// Throws [ServerException] if request is failed
  Future<List<ProductModel>> getProductRefactor(String url,
      [String? id]) async {
    final headerAuth = await authLocalDataSource.getToken();

    headerWithToken['Authorization'] = 'Bearer ${headerAuth.token}';

    try {
      late http.Response result;

      if (id != null) {
        result = await client.get(Uri.parse('${AppData.allProductUrl}/$id'),
            headers: headerWithToken);
      } else {
        result = await client.get(Uri.parse(url), headers: headerWithToken);
      }

      if (result.statusCode == 200) {
        if (id != null) {
          Map<String, dynamic> finalResult = json.decode(result.body);
          return [ProductModel.fromJson(finalResult['data'])];
        } else {
          List<ProductModel> finalResult = <ProductModel>[];
          Map<String, dynamic> jsonModel = json.decode(result.body);

          for (Map<String, dynamic> model in jsonModel['data']) {
            finalResult.add(ProductModel.fromJson(model));
          }

          return finalResult;
        }
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw ServerException();
    } on Exception {
      rethrow;
    }
  }
}
