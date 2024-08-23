import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/core/exception/exception.dart';
import 'package:product_8/core/failure/failure.dart';
import 'package:product_8/features/product/data/models/product_model.dart';
import 'package:product_8/features/product/data/repositories/product_repository_impl.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockProductRemoteDataSource mockProductRemoteDataSource;
  late MockProductLocalDataSource mockProductLocalDataSource;
  late MockNetworkInfo mockNetworkInfo;

  late ProductRepositoryImpl productRepositoryImpl;
  const testProductModelList = [
    ProductModel(
        id: '1',
        name: 'Nike',
        description: 'Nike is the best',
        price: 345.0,
        imageUrl: 'imageUrl')
  ];
  const testProductEntityList = [
    Product(
        id: '1',
        name: 'Nike',
        description: 'Nike is the best',
        price: 345.0,
        imageUrl: 'imageUrl')
  ];
  const testProductModel = ProductModel(
      id: '1',
      name: 'Nike',
      description: 'Nike is the best',
      price: 345.0,
      imageUrl: 'imageUrl');
  const testProductEntity = Product(
      id: '1',
      name: 'Nike',
      description: 'Nike is the best',
      price: 345.0,
      imageUrl: 'imageUrl');

  setUp(() {
    mockProductRemoteDataSource = MockProductRemoteDataSource();
    mockProductLocalDataSource = MockProductLocalDataSource();
    mockNetworkInfo = MockNetworkInfo();
    productRepositoryImpl = ProductRepositoryImpl(
        productRemoteDataSource: mockProductRemoteDataSource,
        productLocalDataSource: mockProductLocalDataSource,
        networkInfo: mockNetworkInfo);
  });

  void runTestsOnline(Function body) {
    group('device is online', () {
      setUp(() {
        when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      });
      body();
    });
  }

  void runTestsOffline(Function body) {
    group('device is offline', () {
      setUp(() {
        when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      });
      body();
    });
  }

// get product by id
  group('get Product', () {
    runTestsOnline(() {
      test('should return product when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.getProductById('1'))
            .thenAnswer((_) async => testProductModel);
        // act
        final result = await productRepositoryImpl.getProduct('1');
        // assert
        expect(result, equals(const Right(testProductEntity)));
      });
    });
    runTestsOnline(() {
      test(
          'should return server faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.getProductById('1'))
            .thenThrow(ServerException());
        // act
        final result = await productRepositoryImpl.getProduct('1');
        // assert
        expect(result, const Left(ServerFailure(message: 'An error occurred')));
      });
    });
    runTestsOnline(() {
      test(
          'should return connection  faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.getProductById('1')).thenThrow(
            const SocketException('Failed to connect to the internet'));
        // act
        final result = await productRepositoryImpl.getProduct('1');
        // assert
        expect(
            result,
            const Left(ConnectionFailure(
                message: 'Failed to connect to the internet')));
      });
    });
  });

// delete product

  group('delete product', () {
    runTestsOnline(() {
      test('should delete product when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.deleteProduct('1'))
            .thenAnswer((_) async => (unit));
        // act
        final result = await productRepositoryImpl.deleteProduct('1');
        // assert
        expect(result, const Right(unit));
      });
    });

    runTestsOnline(() {
      test(
          'should return server faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.deleteProduct('1'))
            .thenThrow(ServerException());
        // act
        final result = await productRepositoryImpl.deleteProduct('1');
        // assert
        expect(result, const Left(ServerFailure(message: 'An error occurred')));
      });
    });

    runTestsOnline(() {
      test(
          'should return connection  faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.deleteProduct('1')).thenThrow(
            const SocketException('Failed to connect to the internet'));
        // act
        final result = await productRepositoryImpl.deleteProduct('1');
        // assert
        expect(
            result,
            const Left(ConnectionFailure(
                message: 'Failed to connect to the internet')));
      });
    });
  });

  // update product
  group('Update product', () {
    runTestsOnline(() {
      test('should update product when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.updateProduct(testProductModel))
            .thenAnswer((_) async => testProductModel);
        // act
        final result =
            await productRepositoryImpl.updateProduct(testProductModel);
        // assert
        expect(result, equals(const Right(testProductEntity)));
      });
    });

    runTestsOnline(() {
      test(
          'should return server faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.updateProduct(testProductModel))
            .thenThrow(ServerException());
        // act
        final result =
            await productRepositoryImpl.updateProduct(testProductModel);
        // assert
        expect(result, const Left(ServerFailure(message: 'An error occurred')));
      });
    });

    runTestsOnline(() {
      test(
          'should return connection  faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.updateProduct(testProductModel))
            .thenThrow(
                const SocketException('Failed to connect to the internet'));
        // act
        final result =
            await productRepositoryImpl.updateProduct(testProductModel);
        // assert
        expect(
            result,
            const Left(ConnectionFailure(
                message: 'Failed to connect to the internet')));
      });
    });
  });

  // create product
  group('create product', () {
    runTestsOnline(() {
      test('should create product when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.createProduct(testProductModel))
            .thenAnswer((_) async => testProductModel);
        // act
        final result =
            await productRepositoryImpl.createProduct(testProductModel);
        // assert
        expect(result, equals(const Right(testProductEntity)));
      });
    });

    runTestsOnline(() {
      test(
          'should return server faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.updateProduct(testProductModel))
            .thenThrow(ServerException());
        // act
        final result =
            await productRepositoryImpl.updateProduct(testProductModel);
        // assert
        expect(result, const Left(ServerFailure(message: 'An error occurred')));
      });
    });

    runTestsOnline(() {
      test(
          'should return connection  faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.createProduct(testProductModel))
            .thenThrow(
                const SocketException('Failed to connect to the internet'));
        // act
        final result =
            await productRepositoryImpl.createProduct(testProductModel);
        // assert
        expect(
            result,
            const Left(ConnectionFailure(
                message: 'Failed to connect to the internet')));
      });
    });
  });

  //  get products
  group('get all products', () {
    runTestsOnline(() {
      test('should return products when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.getProducts())
            .thenAnswer((_) async => testProductModelList);
        // act
        final result = await productRepositoryImpl.getProducts();
        final unPackedResult =
            result.fold((failure) => null, (productList) => productList);
        // assert
        expect(unPackedResult, equals(testProductEntityList));
      });
    });

    runTestsOnline(() {
      test(
          'should cache products after getting them from the remote data source',
          () async {
        //arrange
        when(mockProductRemoteDataSource.getProducts())
            .thenAnswer((_) async => testProductModelList);
        when(mockProductLocalDataSource.cacheProducts(testProductModelList))
            .thenAnswer((_) async => unit);

        //act
        await productRepositoryImpl.getProducts();

        //assert
        verify(mockProductLocalDataSource.cacheProducts(testProductModelList));
      });
    });
    runTestsOffline(() {
      test('should return cached products when no network is available',
          () async {
        //arrange
        when(mockProductLocalDataSource.getProducts())
            .thenAnswer((_) async => testProductModelList);

        //act
        final result = await productRepositoryImpl.getProducts();
        final unpackedResult =
            result.fold((failure) => null, (productList) => productList);

        //assert
        expect(unpackedResult, equals(testProductEntityList));
      });
    });

    runTestsOnline(() {
      test(
          'should return a server failure when a call to the remote data source is unsuccessful',
          () async {
        //arrange
        when(mockProductRemoteDataSource.getProducts())
            .thenThrow(ServerException());
        //act
        final result = await productRepositoryImpl.getProducts();

        //assert
        expect(
            result, equals(const Left(ServerFailure(message: 'An error occurred'))));
      });
    });
    runTestsOnline(() {
      test(
          'should return connection  faliure when a call to data source is success ',
          () async {
        // arrange
        when(mockProductRemoteDataSource.getProducts()).thenThrow(
            const SocketException('Failed to connect to the internet'));
        // act
        final result = await productRepositoryImpl.getProducts();
        // assert
        expect(
            result,
            const Left(ConnectionFailure(
                message: 'Failed to connect to the internet')));
      });
    });

    runTestsOffline(() {
      test('should return cache failure when failing to get cached products',
          () async {
        //arrange
        when(mockProductLocalDataSource.getProducts())
            .thenThrow(CacheException());

        //act
        final result = await productRepositoryImpl.getProducts();

        //assert
        expect(
            result, equals(const Left(CacheFailure(message: 'An error occurred'))));
      });
    });
  });
}
