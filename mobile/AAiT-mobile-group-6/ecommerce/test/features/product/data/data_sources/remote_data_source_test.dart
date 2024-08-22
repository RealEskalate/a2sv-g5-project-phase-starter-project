import 'dart:convert';

import 'package:ecommerce/core/constants/constants.dart';
import 'package:ecommerce/core/error/exceptions.dart';
import 'package:ecommerce/features/product/data/datasources/remote_data_source.dart';
import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../helpers/json_reader.dart';
import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late ProductRemoteDataSourceImpl productRemoteDataSourceImpl;
  setUp(() {
    mockHttpClient = MockHttpClient();
    productRemoteDataSourceImpl =
        ProductRemoteDataSourceImpl(client: mockHttpClient);
  });
  const String testId = '1';
  ProductModel testProduct = const ProductModel(
    id: '1',
    name: 'name',
    description: 'description',
    price: 1.0,
    imageUrl: 'imageUrl',
  );

  group('get current product', () {
    test('should return the product model when the response code is 200',
        () async {
      //arrange
      when(mockHttpClient.get(Uri.parse(Urls.getProductById(testId))))
          .thenAnswer((_) async =>
              http.Response(readJson('dummy_product_response.json'), 200));
      //act
      final result = await productRemoteDataSourceImpl.getProduct(testId);
      //assert
      expect(result, isA<Product>());
    });

    test('should throw server exception error when the response code is 404',
        () async {
      //arrange
      when(mockHttpClient.get(Uri.parse(Urls.getProductById(testId))))
          .thenAnswer((_) async => http.Response('Not found', 404));
      //act
      final result = productRemoteDataSourceImpl.getProduct(testId);
      //assert
      expect(result, throwsA(isA<ServerException>()));
    });
  });

  group('delete current product', () {
    test('should return the product model when the response code is 200',
        () async {
      //arrange
      when(mockHttpClient.delete(Uri.parse(Urls.getProductById(testId))))
          .thenAnswer((_) async =>
              http.Response(readJson('dummy_product_response.json'), 200));
      //act
      final result = await productRemoteDataSourceImpl.deleteProduct(testId);
      //assert
      expect(result, isA<Product>());
    });

    test('should throw server exception error when the response code is 404',
        () async {
      //arrange
      when(mockHttpClient.delete(Uri.parse(Urls.getProductById(testId))))
          .thenAnswer((_) async => http.Response('Not found', 404));
      //act
      final result = productRemoteDataSourceImpl.deleteProduct(testId);
      //assert
      expect(result, throwsA(isA<ServerException>()));
    });
  });

  group('add current product', () {
    test('should return the product model when the response code is 200',
        () async {
      //arrange
      when(mockHttpClient.post(
        Uri.parse(Urls.getAllProducts()),
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async =>
          http.Response(readJson('dummy_product_response.json'), 201));
      //act
      final result =
          await productRemoteDataSourceImpl.insertProduct(testProduct);
      //assert
      expect(result, isA<ProductModel>());
    });

    test('should throw server exception error when the response code is 404',
        () async {
      //arrange
      when(mockHttpClient.post(
        Uri.parse(Urls.getProductById(testId)),
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response('Not found', 404));
      //act
      final result = productRemoteDataSourceImpl.insertProduct(testProduct);
      //assert
      expect(result, throwsA(isA<ServerException>()));
    });
  });

  group('update current product', () {
    test('should return the product model when the response code is 200',
        () async {
      //arrange
      when(mockHttpClient.put(
        Uri.parse(Urls.getProductById(testProduct.id)),
        headers: {'Content-Type': 'application/json'},
        body: json.encode(ProductModel.fromProduct(testProduct).toJson()),
      )).thenAnswer((_) async => http.Response(
          '{"data": ${readJson('dummy_product_response.json')}}', 200));
      //act
      final result =
          await productRemoteDataSourceImpl.updateProduct(testProduct);
      //assert
      expect(result, isA<ProductModel>());
    });

    test('should throw server exception error when the response code is 404',
        () async {
      //arrange
      when(mockHttpClient.put(
        Uri.parse(Urls.getProductById(testProduct.id)),
        headers: {'Content-Type': 'application/json'},
        body: json.encode(testProduct.toJson()),
      )).thenAnswer((_) async => http.Response('Not found', 404));
      //act
      final result = productRemoteDataSourceImpl.updateProduct(testProduct);
      //assert
      expect(result, throwsA(isA<ServerException>()));
    });
  });

  group('get all products', () {
    test('should return the product model when the response code is 200',
        () async {
      //arrange
      when(mockHttpClient.get(
        Uri.parse(Urls.getAllProducts()),
      )).thenAnswer((_) async => http.Response(
          '{"data": ${readJson('dummy_products_response.json')}}', 200));
      //act
      final result = await productRemoteDataSourceImpl.getAllProduct();
      //assert
      expect(result, isA<List<ProductModel>>());
    });

    test('should throw server exception error when the response code is 404',
        () async {
      //arrange
      when(
        mockHttpClient.get(Uri.parse(Urls.getAllProducts())),
      ).thenAnswer((_) async => http.Response('Not found', 404));
      //act
      final result = productRemoteDataSourceImpl.getAllProduct();
      //assert
      expect(result, throwsA(isA<ServerException>()));
    });
  });
}
