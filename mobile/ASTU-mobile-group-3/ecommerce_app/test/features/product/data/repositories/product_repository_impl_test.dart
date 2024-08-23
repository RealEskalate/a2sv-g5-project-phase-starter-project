import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/exceptions/product_exceptions.dart';
import 'package:ecommerce_app/core/errors/failures/failure.dart';

import 'package:ecommerce_app/features/product/data/repositories/product_repository_impl.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  late MockNetworkInfo mockNetworkInfo;
  late MockLocalProductDataSource mockLocalProductDataSource;
  late MockRemoteProductDataSource mockRemoteProductDataSource;
  late ProductRepositoryImpl productRepositoryImpl;

  setUp(() {
    mockNetworkInfo = MockNetworkInfo();
    mockLocalProductDataSource = MockLocalProductDataSource();
    mockRemoteProductDataSource = MockRemoteProductDataSource();

    productRepositoryImpl = ProductRepositoryImpl(
      remoteProductDataSource: mockRemoteProductDataSource,
      localProductDataSource: mockLocalProductDataSource,
      networkInfo: mockNetworkInfo,
    );
  });

  group('Testing the Implemented Repository by checking thee network info', () {
    /// Test for online calls
    group('Offline actions', () {
      setUp(() {
        when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      });

      /// This test is for getAllProduct
      test('getAllProduct Should return valid list of ProductEntity', () async {
        /// arrange
        when(mockRemoteProductDataSource.getAllProducts())
            .thenAnswer((_) async => TestingDatas.productModelList);

        /// action
        ///
        final result = await productRepositoryImpl.getAllProducts();

        /// assert
        // will verify if isConnected method form network info is called
        verify(mockNetworkInfo.isConnected);
        verify(mockLocalProductDataSource
            .addListOfProduct(TestingDatas.productModelList));
        result.fold((failure) {}, (data) {
          expect(data, TestingDatas.productEntityList);
        });
      });

      /// This test is for getProduct
      test('Should call the network method and return appropriate data',
          () async {
        /// arrange
        when(mockRemoteProductDataSource.getProduct(TestingDatas.id))
            .thenAnswer((_) async => TestingDatas.testDataModel);

        /// action

        final result = await productRepositoryImpl.getProduct(TestingDatas.id);

        /// assertion

        verify(mockNetworkInfo.isConnected);
        verify(
            mockLocalProductDataSource.addProduct(TestingDatas.testDataModel));
        expect(result, const Right(TestingDatas.testDataEntity));
      });

      /// This test for updateProduct
      test('Should return number of row affected when succes when updated',
          () async {
        /// arrange
        when(mockRemoteProductDataSource
                .updateProduct(TestingDatas.testDataModel))
            .thenAnswer((_) async => AppData.successUpdate);

        ///action

        final result = await productRepositoryImpl
            .updateProduct(TestingDatas.testDataEntity);

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(mockLocalProductDataSource
            .updateProduct(TestingDatas.testDataModel));

        expect(result, const Right(AppData.successUpdate));
      });

      /// this test is for deleteProduct
      test('Should return number of row affected when succes when deleted',
          () async {
        /// arrange
        when(mockRemoteProductDataSource.deleteProduct(TestingDatas.id))
            .thenAnswer((_) async => AppData.successDelete);

        ///action

        final result =
            await productRepositoryImpl.deleteProduct(TestingDatas.id);

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(mockLocalProductDataSource.removeProduct(TestingDatas.id));
        expect(result, const Right(AppData.successDelete));
      });

      /// This test is for insertProduct
      test(
          'Should return number of row affected when succes when new data is inserted',
          () async {
        /// arrange
        when(mockRemoteProductDataSource
                .insertProduct(TestingDatas.testDataModel))
            .thenAnswer((_) async => AppData.successInsert);

        ///action

        final result = await productRepositoryImpl
            .insertProduct(TestingDatas.testDataEntity);

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(
            mockLocalProductDataSource.addProduct(TestingDatas.testDataModel));
        expect(result, const Right(AppData.successInsert));
      });
    });

    /// Test for offline calls
    group('offline actions', () {
      setUp(() {
        when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      });

      /// Similar to the online test but has to return when datamanipulation

      /// This test is for getAllProduct offline
      test(
          'getAllProduct Should return valid list of ProductEntity from a cache',
          () async {
        /// arrange
        when(mockLocalProductDataSource.getAllProducts())
            .thenAnswer((_) async => TestingDatas.productModelList);

        /// action
        ///
        final result = await productRepositoryImpl.getAllProducts();

        /// assert
        // will verify if isConnected method form network info is called
        verify(mockNetworkInfo.isConnected);
        verify(mockLocalProductDataSource.getAllProducts());
        result.fold((failure) {}, (data) {
          expect(data, TestingDatas.productEntityList);
        });
      });

      /// This test is for getProduct offline
      test('Should get data from local repo when offline', () async {
        /// arrange
        when(mockLocalProductDataSource.getProduct(TestingDatas.id))
            .thenAnswer((_) async => TestingDatas.testDataModel);

        /// action

        final result = await productRepositoryImpl.getProduct(TestingDatas.id);

        /// assertion

        verify(mockNetworkInfo.isConnected);
        verify(mockLocalProductDataSource.getProduct(TestingDatas.id));
        expect(result, const Right(TestingDatas.testDataEntity));
      });

      /// This test for updateProduct offline
      test('Should return failure when update fails', () async {
        ///action

        final result = await productRepositoryImpl
            .updateProduct(TestingDatas.testDataEntity);

        /// assert

        expect(
            result,
            Left(ConnectionFailure(
                AppData.getMessage(AppData.connectionError))));
      });

      /// this test is for deleteProduct offline
      test('Should return failure when deletion fails', () async {
        ///action

        final result =
            await productRepositoryImpl.deleteProduct(TestingDatas.id);

        /// assert

        expect(
            result,
            Left(ConnectionFailure(
                AppData.getMessage(AppData.connectionError))));
      });

      /// This test is for insertProduct when offline
      test('Should return failure  insertion fails', () async {
        ///action

        final result = await productRepositoryImpl
            .insertProduct(TestingDatas.testDataEntity);

        /// assert

        expect(
            result,
            Left(ConnectionFailure(
                AppData.getMessage(AppData.connectionError))));
      });
    });

    group('See How the methods react to server exception or cache exception',
        () {
      setUp(() {
        when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      });

      /// This test for when ServerException is thrown
      test('Should return ServerFailure when internet fails', () async {
        /// arrange
        when(mockRemoteProductDataSource.getAllProducts())
            .thenThrow(ServerException());

        /// action
        final result = await productRepositoryImpl.getAllProducts();

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(mockRemoteProductDataSource.getAllProducts());

        expect(result,
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
      });

      test(
          'Should return Server failure when getting product  is returned server exception',
          () async {
        when(mockRemoteProductDataSource.getProduct(TestingDatas.id))
            .thenThrow(ServerException());

        /// action
        final result = await productRepositoryImpl.getProduct(TestingDatas.id);

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(mockRemoteProductDataSource.getProduct(TestingDatas.id));

        expect(result,
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
      });

      test(
          'Should return server failure when the deleting user returned server excepion',
          () async {
        when(mockRemoteProductDataSource.deleteProduct(TestingDatas.id))
            .thenThrow(ServerException());

        /// action
        final result =
            await productRepositoryImpl.deleteProduct(TestingDatas.id);

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(mockRemoteProductDataSource.deleteProduct(TestingDatas.id));

        expect(result,
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
      });

      test(
          'Should return server failure when inserting returns server exception ',
          () async {
        when(mockRemoteProductDataSource
                .insertProduct(TestingDatas.testDataModel))
            .thenThrow(ServerException());

        /// action
        final result = await productRepositoryImpl
            .insertProduct(TestingDatas.testDataEntity);

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(mockRemoteProductDataSource
            .insertProduct(TestingDatas.testDataModel));

        expect(result,
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
      });

      test(
          'Should return server failure when updating returns server exception ',
          () async {
        when(mockRemoteProductDataSource
                .updateProduct(TestingDatas.testDataModel))
            .thenThrow(ServerException());

        /// action
        final result = await productRepositoryImpl
            .updateProduct(TestingDatas.testDataEntity);

        /// assert
        verify(mockNetworkInfo.isConnected);
        verify(mockRemoteProductDataSource
            .updateProduct(TestingDatas.testDataModel));

        expect(result,
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
      });
    });

    group('When cache exception of getProduct and getAllProduct is thrown', () {
      setUp(() {
        when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      });
      test('Should return cache failure whe  cache exception is thrown',
          () async {
        /// arrange
        when(mockLocalProductDataSource.getProduct(TestingDatas.id))
            .thenThrow(CacheException());

        /// action
        final result = await productRepositoryImpl.getProduct(TestingDatas.id);

        /// assert

        verify(mockNetworkInfo.isConnected);
        verify(mockLocalProductDataSource.getProduct(TestingDatas.id));
        expect(
            result, Left(CacheFailure(AppData.getMessage(AppData.cacheError))));
      });
      test('Get all should return cache exception when the local data is empty',
          () async {
        /// arrange
        when(mockLocalProductDataSource.getAllProducts())
            .thenThrow(CacheException());

        /// action
        final result = await productRepositoryImpl.getAllProducts();
        verify(mockNetworkInfo.isConnected);
        verify(mockLocalProductDataSource.getAllProducts());
        expect(
            result, Left(CacheFailure(AppData.getMessage(AppData.cacheError))));
      });
    });
  });
}
