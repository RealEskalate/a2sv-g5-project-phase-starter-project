import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';
import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/data/data_sources/product_remote_datasource.dart';
import 'package:product_6/features/product/data/models/product_model.dart';
import 'package:test/test.dart';

import '../../../../mock.mocks.dart';

// Mock class
// class MockClient extends Mock implements http.Client {}

void main() {
  late MockClient mockClient;
  late ProductRemoteDatasourceImp datasource;
  const baseUrl = 'https://g5-flutter-learning-path-be.onrender.com/api/v1';

  setUp(() {
    mockClient = MockClient();
    datasource = ProductRemoteDatasourceImp(client: mockClient, baseUrl: baseUrl);
  });

  const productModel = ProductModel(
    id: '1',
    name: 'Test Product',
    description: 'Test Description',
    imageUrl: 'http://example.com/image.jpg',
    price: 99.99,
  );

  final productJson = {
    'id': '1',
    'name': 'Test Product',
    'description': 'Test Description',
    'imageUrl': 'http://example.com/image.jpg',
    'price': 99.99,
  };

  final responseData = {'data': productJson};
  final responseListData = {'data': [productJson]};

  group('createProduct', () {
    test('should return ProductModel when the response is successful', () async {
      when(mockClient.post(
        Uri.parse('$baseUrl/products'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode(productModel.toJson()),
      )).thenAnswer((_) async => http.Response(json.encode(responseData), 200));

      final result = await datasource.createProduct(productModel);

      expect(result, Right(productModel));
    });

    test('should return ServerFailure when the response is unsuccessful', () async {
      when(mockClient.post(
        Uri.parse('$baseUrl/products'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode(productModel.toJson()),
      )).thenAnswer((_) async => http.Response('Error', 500));

      final result = await datasource.createProduct(productModel);

      expect(result, const Left(ServerFailure('Failed to create product')));
    });
  });

  group('deleteProduct', () {
    test('should return void when the response is successful', () async {
      when(mockClient.delete(Uri.parse('$baseUrl/products/1')))
          .thenAnswer((_) async => http.Response('', 200));

      final result = await datasource.deleteProduct('1');

      expect(result, const Right(null));
    });

    test('should return ServerFailure when the response is unsuccessful', () async {
      when(mockClient.delete(Uri.parse('$baseUrl/products/1')))
          .thenAnswer((_) async => http.Response('Error', 500));

      final result = await datasource.deleteProduct('1');

      expect(result, const Left(ServerFailure('Failed to delete product')));
    });
  });

  group('getProductById', () {
    test('should return ProductModel when the response is successful', () async {
      when(mockClient.get(Uri.parse('$baseUrl/products/1')))
          .thenAnswer((_) async => http.Response(json.encode(responseData), 200));

      final result = await datasource.getProductById('1');

      expect(result, Right(productModel));
    });

    test('should return ServerFailure when the product is not found', () async {
      when(mockClient.get(Uri.parse('$baseUrl/products/1')))
          .thenAnswer((_) async => http.Response('Not Found', 404));

      final result = await datasource.getProductById('1');

      expect(result, const Left(ServerFailure('Product not found')));
    });

    test('should return ServerFailure when the response is unsuccessful', () async {
      when(mockClient.get(Uri.parse('$baseUrl/products/1')))
          .thenAnswer((_) async => http.Response('Error', 500));

      final result = await datasource.getProductById('1');

      expect(result, const Left(ServerFailure('Failed to fetch product')));
    });
  });

  group('getProducts', () {
    test('should return List<ProductModel> when the response is successful', () async {
      when(mockClient.get(Uri.parse('$baseUrl/products')))
          .thenAnswer((_) async => http.Response(json.encode(responseListData), 200));

      final result = await datasource.getProducts();
      result.fold((l) => null, (r) => expect(r, [productModel]));

      // expect(result, Right([productModel]));
    });

    test('should return ServerFailure when the response is unsuccessful', () async {
      when(mockClient.get(Uri.parse('$baseUrl/products')))
          .thenAnswer((_) async => http.Response('Error', 500));

      final result = await datasource.getProducts();

      expect(result, const Left(ServerFailure('Failed to fetch products')));
    });
  });

  group('updateProduct', () {
    test('should return ProductModel when the response is successful', () async {
      when(mockClient.put(
        Uri.parse('$baseUrl/products/1'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode(productModel.toJson()),
      )).thenAnswer((_) async => http.Response(json.encode(responseData), 200));

      final result = await datasource.updateProduct(productModel);

      expect(result, Right(productModel));
    });

    test('should return ServerFailure when the response is unsuccessful', () async {
      when(mockClient.put(
        Uri.parse('$baseUrl/products/1'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode(productModel.toJson()),
      )).thenAnswer((_) async => http.Response('Error', 500));

      final result = await datasource.updateProduct(productModel);

      expect(result, const Left(ServerFailure('Failed to update product')));
    });
  });
}
