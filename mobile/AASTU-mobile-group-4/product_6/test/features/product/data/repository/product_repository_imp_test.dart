import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/data/models/product_model.dart';
import 'package:product_6/features/product/data/repository/product_repository_imp.dart';
import 'package:product_6/features/product/domain/entities/product.dart';

import '../../../../mock.mocks.dart'; // Adjust the path to your mocks file

void main() {
  late MockProductRemoteDatasource mockProductRemoteDatasource;
  late MockNetworkInfo mockNetworkInfo;
  late ProductRepositoryImp productRepositoryImp;

  setUp(() {
    mockProductRemoteDatasource = MockProductRemoteDatasource();
    mockNetworkInfo = MockNetworkInfo();
    productRepositoryImp = ProductRepositoryImp(
      productRemoteDatasource: mockProductRemoteDatasource,
      networkInfo: mockNetworkInfo,
    );
  });

  const productModel = ProductModel(
    id: '1',
    name: 'Test Product',
    description: 'Test Description',
    imageUrl: 'http://example.com/image.jpg',
    price: 99.99,
  );

  final Product product = productModel.toEntity();
  group('createProduct', () {
    test(
        'should return Product when the call to remote data source is successful and device is online',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.createProduct(any))
          .thenAnswer((_) async => const Right(productModel));

      // Act
      final result = await productRepositoryImp.createProduct(product);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.createProduct(productModel));
      expect(result, equals(Right(product)));
    });

    test(
        'should return ServerFailure when the call to remote data source is unsuccessful',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.createProduct(any)).thenAnswer(
          (_) async => const Left(ServerFailure('An error has occurred')));

      // Act
      final result = await productRepositoryImp.createProduct(product);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.createProduct(productModel));
      expect(
          result, equals(const Left(ServerFailure('An error has occurred'))));
    });

    test('should return ConnectionFailure when there is no internet connection',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      // Act
      final result = await productRepositoryImp.createProduct(product);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verifyNever(mockProductRemoteDatasource.createProduct(any));
      expect(result,
          equals(const Left(ConnectionFailure('No internet connection.'))));
    });
  });
  group('deleteProduct', () {
    const String tId = '1';

    test(
        'should return void when the call to remote data source is successful and device is online',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.deleteProduct(tId))
          .thenAnswer((_) async => const Right(null));

      // Act
      final result = await productRepositoryImp.deleteProduct(tId);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.deleteProduct(tId));
      expect(result, equals(const Right(null)));
    });

    test(
        'should return ServerFailure when the call to remote data source is unsuccessful',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.deleteProduct(tId)).thenAnswer(
          (_) async => const Left(ServerFailure('An error has occurred')));

      // Act
      final result = await productRepositoryImp.deleteProduct(tId);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.deleteProduct(tId));
      expect(
          result, equals(const Left(ServerFailure('An error has occurred'))));
    });

    test('should return ConnectionFailure when there is no internet connection',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      // Act
      final result = await productRepositoryImp.deleteProduct(tId);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verifyNever(mockProductRemoteDatasource.deleteProduct(any));
      expect(result,
          equals(const Left(ConnectionFailure('No internet connection.'))));
    });
  });
  group('getProductById', () {
    const String tId = '1';

    test(
        'should return Product when the call to remote data source is successful and device is online',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.getProductById(tId))
          .thenAnswer((_) async => const Right(productModel));

      // Act
      final result = await productRepositoryImp.getProductById(tId);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.getProductById(tId));
      expect(result, equals(Right(product)));
    });

    test(
        'should return ServerFailure when the call to remote data source is unsuccessful',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.getProductById(tId)).thenAnswer(
          (_) async => const Left(ServerFailure('An error has occurred')));

      // Act
      final result = await productRepositoryImp.getProductById(tId);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.getProductById(tId));
      expect(
          result, equals(const Left(ServerFailure('An error has occurred'))));
    });

    test('should return ConnectionFailure when there is no internet connection',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      // Act
      final result = await productRepositoryImp.getProductById(tId);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verifyNever(mockProductRemoteDatasource.getProductById(any));
      expect(result,
          equals(const Left(ConnectionFailure('No internet connection.'))));
    });
  });

  group('getProducts', () {
    final List<ProductModel> productModelList = [productModel];
    final List<Product> productList =
        productModelList.map((model) => model.toEntity()).toList();

    test(
        'should return List<Product> when the call to remote data source is successful and device is online',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.getProducts())
          .thenAnswer((_) async => Right(productModelList));

      // Act
      final result = await productRepositoryImp.getProducts();

      // Assert
      result.fold((l) => const ServerFailure('message'),
          (r) => expect(r, productList));
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.getProducts());
      // expect(result, equals(Right(productList)));
    });

    test(
        'should return ServerFailure when the call to remote data source is unsuccessful',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.getProducts()).thenAnswer(
          (_) async => const Left(ServerFailure('An error has occurred')));

      // Act
      final result = await productRepositoryImp.getProducts();

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.getProducts());
      expect(
          result, equals(const Left(ServerFailure('An error has occurred'))));
    });

    test('should return ConnectionFailure when there is no internet connection',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      // Act
      final result = await productRepositoryImp.getProducts();

      // Assert
      verify(mockNetworkInfo.isConnected);
      verifyNever(mockProductRemoteDatasource.getProducts());
      expect(result,
          equals(const Left(ConnectionFailure('No internet connection.'))));
    });
  });

  group('updateProduct', () {
    test(
        'should return Product when the call to remote data source is successful and device is online',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.updateProduct(any))
          .thenAnswer((_) async => const Right(productModel));

      // Act
      final result = await productRepositoryImp.updateProduct(product);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.updateProduct(productModel));
      expect(result, equals(Right(product)));
    });

    test(
        'should return ServerFailure when the call to remote data source is unsuccessful',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockProductRemoteDatasource.updateProduct(any)).thenAnswer(
          (_) async => const Left(ServerFailure('An error has occurred')));

      // Act
      final result = await productRepositoryImp.updateProduct(product);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verify(mockProductRemoteDatasource.updateProduct(productModel));
      expect(
          result, equals(const Left(ServerFailure('An error has occurred'))));
    });

    test('should return ConnectionFailure when there is no internet connection',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      // Act
      final result = await productRepositoryImp.updateProduct(product);

      // Assert
      verify(mockNetworkInfo.isConnected);
      verifyNever(mockProductRemoteDatasource.updateProduct(any));
      expect(result,
          equals(const Left(ConnectionFailure('No internet connection.'))));
    });
  });
}
