// ignore_for_file: unused_element

import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/exceptions.dart';
import 'package:ecommerce/core/error/failures.dart';
import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:ecommerce/features/product/data/repositories/product_repository_impl.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late ProductRepositoryImpl repository;
  late MockProductRemoteDataSource mockRemoteDataSource;
  late MockProductLocalDataSource mockLocalDataSource;
  late MockNetworkInfo mockNetworkInfo;

  setUp(() {
    mockRemoteDataSource = MockProductRemoteDataSource();
    mockLocalDataSource = MockProductLocalDataSource();
    mockNetworkInfo = MockNetworkInfo();
    repository = ProductRepositoryImpl(
      remoteDataSource: mockRemoteDataSource,
      localDataSource: mockLocalDataSource,
      networkInfo: mockNetworkInfo,
    );
  });

  final testId = '1';
  const testProduct = ProductModel(
    id: '1',
    name: 'name',
    description: 'description',
    price: 1.0,
    imageUrl: 'imageUrl',
  
  );

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

  group('getProduct', () {
    test('should check if the device is online', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.getProduct(any))
          .thenAnswer((_) async => testProduct);

      // Act
      await repository.getProduct(testId);

      // Assert
      verify(mockNetworkInfo.isConnected);
    });

    void runTestsOnline() {
      test(
        'should return remote data when the call to remote data source is successful',
        () async {
          // Arrange
          when(mockRemoteDataSource.getProduct(testId))
              .thenAnswer((_) async => testProduct);

          // Act
          final result = await repository.getProduct(testId);

          // Assert
          verify(mockRemoteDataSource.getProduct(testId));
          expect(result, equals(const Right(testProduct)));
        },
      );

      test(
        'should cache the data locally when the call to remote data source is successful',
        () async {
          // Arrange
          when(mockRemoteDataSource.getProduct(testId))
              .thenAnswer((_) async => testProduct);

          // Act
          await repository.getProduct(testId);

          // Assert
          verify(mockRemoteDataSource.getProduct(testId));
          verify(mockLocalDataSource.cacheProduct(testProduct));
        },
      );

      test(
        'should return server failure when the call to remote data source is unsuccessful',
        () async {
          // Arrange
          when(mockRemoteDataSource.getProduct(testId))
              .thenThrow(ServerException());

          // Act
          final result = await repository.getProduct(testId);

          // Assert
          verify(mockRemoteDataSource.getProduct(testId));
          verifyZeroInteractions(mockLocalDataSource);
          expect(result, equals(Left(ServerFailure('Server failure'))));
        },
      );
    }

    runTestsOffline(() {
      test(
        'should return last locally cached data when the cached data is present',
        () async {
          // Arrange
          when(mockLocalDataSource.getLastProduct())
              .thenAnswer((_) async => testProduct);
          when(mockRemoteDataSource.getProduct(any))
              .thenAnswer((_) async => testProduct);
          // Act
          final result = await repository.getProduct(testId);

          // Assert
          verifyZeroInteractions(mockRemoteDataSource);
          verify(mockLocalDataSource.getLastProduct());
          expect(result, equals(const Right(testProduct)));
        },
      );

      test(
        'should return CacheFailure when there is no cached data present',
        () async {
          // Arrange
          when(mockLocalDataSource.getLastProduct())
              .thenThrow(CacheException());
          when(mockRemoteDataSource.getProduct(any))
              .thenAnswer((_) async => testProduct);
          // Act
          final result = await repository.getProduct(testId);

          // Assert
          verifyZeroInteractions(mockRemoteDataSource);
          verify(mockLocalDataSource.getLastProduct());
          expect(result, equals(Left(CacheFailure('Cache failure'))));
        },
      );
    });
  });

  group('insertProduct', () {
    test('should check if the device is online', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.insertProduct(testProduct))
          .thenAnswer((_) async => testProduct);

      // Act
      await repository.insertProduct(testProduct);

      // Assert
      verify(mockNetworkInfo.isConnected);
    });
    void runTestsOnline() {
      test('should call the remote data source to insert the product',
          () async {
        // Arrange
        when(mockRemoteDataSource.insertProduct(testProduct))
            .thenAnswer((_) async => testProduct);

        // Act
        await repository.insertProduct(testProduct);

        // Assert
        verify(mockRemoteDataSource.insertProduct(testProduct));
      });

      test(
          'should cache the product locally when the call to remote data source is successful',
          () async {
        // Arrange
        when(mockRemoteDataSource.insertProduct(testProduct))
            .thenAnswer((_) async => testProduct);

        // Act
        await repository.insertProduct(testProduct);

        // Assert
        verify(mockRemoteDataSource.insertProduct(testProduct));
        verify(mockLocalDataSource.cacheProduct(testProduct));
      });

      test(
          'should return server failure when the call to remote data source is unsuccessful',
          () async {
        // Arrange
        when(mockRemoteDataSource.insertProduct(testProduct))
            .thenThrow(ServerException());

        // Act
        final result = await repository.insertProduct(testProduct);

        // Assert
        verify(mockRemoteDataSource.insertProduct(testProduct));
        verifyZeroInteractions(mockLocalDataSource);
        expect(result, equals(Left(ServerFailure('Server failure'))));
      });
    }

    void runTestsOffline() {
      test('should return network failure when insert product is called',
          () async {
        // Act
        final result = await repository.insertProduct(testProduct);

        // Assert
        verifyZeroInteractions(mockRemoteDataSource);
        verifyZeroInteractions(mockLocalDataSource);
        expect(result, equals(Left(NetworkFailure('Network failure'))));
      });
    }
  });

  group('deleteProduct', () {
    test('should check if the device is online', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.deleteProduct(testId))
          .thenAnswer((_) async => testProduct);

      // Act
      await repository.deleteProduct(testId);

      // Assert
      verify(mockNetworkInfo.isConnected);
    });

    void runTestsOnline() {
      test(
        'should return remote data when the call to remote data source is successful',
        () async {
          // Arrange
          when(mockRemoteDataSource.deleteProduct(testId))
              .thenAnswer((_) async => testProduct);

          // Act
          final result = await repository.deleteProduct(testId);

          // Assert
          verify(mockRemoteDataSource.deleteProduct(testId));
          expect(result, equals(const Right(testProduct)));
        },
      );

      test(
        'should return server failure when the call to remote data source is unsuccessful',
        () async {
          // Arrange
          when(mockRemoteDataSource.deleteProduct(testId))
              .thenThrow(ServerException());

          // Act
          final result = await repository.deleteProduct(testId);

          // Assert
          verify(mockRemoteDataSource.deleteProduct(testId));
          expect(result, equals(Left(ServerFailure('Server failure'))));
        },
      );
    }

    runTestsOffline(() {
      test(
        'should return network failure when delete product is called',
        () async {
          // Arrange
          when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

          // Act
          final result = await repository.deleteProduct(testId);

          // Assert
          verifyZeroInteractions(mockRemoteDataSource);
          verifyZeroInteractions(mockLocalDataSource);
          expect(result, equals(Left(NetworkFailure('Network failure'))));
        },
      );
    });
  });

  group('updateProduct', () {
    test('should check if the device is online', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.updateProduct(testProduct))
          .thenAnswer((_) async => testProduct);

      // Act
      await repository.updateProduct(testProduct);

      // Assert
      verify(mockNetworkInfo.isConnected);
    });

    void runTestsOnline() {
      test(
        'should return remote data when the call to remote data source is successful',
        () async {
          // Arrange
          when(mockRemoteDataSource.updateProduct(testProduct))
              .thenAnswer((_) async => testProduct);

          // Act
          final result = await repository.updateProduct(testProduct);

          // Assert
          verify(mockRemoteDataSource.updateProduct(testProduct));
          expect(result, equals(const Right(testProduct)));
        },
      );

      test(
        'should return server failure when the call to remote data source is unsuccessful',
        () async {
          // Arrange
          when(mockRemoteDataSource.updateProduct(testProduct))
              .thenThrow(ServerException());

          // Act
          final result = await repository.updateProduct(testProduct);

          // Assert
          verify(mockRemoteDataSource.updateProduct(testProduct));
          expect(result, equals(Left(ServerFailure('Server failure'))));
        },
      );
    }

    runTestsOffline(() {
      test(
        'should return network failure when update product is called',
        () async {
          // Arrange
          when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

          // Act
          final result = await repository.updateProduct(testProduct);

          // Assert
          verifyZeroInteractions(mockRemoteDataSource);
          verifyZeroInteractions(mockLocalDataSource);
          expect(result, equals(Left(NetworkFailure('Network failure'))));
        },
      );
    });
  });

  group('get all products', () {
    test('should check if the device is online', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.getAllProduct())
          .thenAnswer((_) async => [testProduct]);

      // Act
      await repository.getAllProduct();

      // Assert
      verify(mockNetworkInfo.isConnected);
    });
    runTestsOnline() {
      test(
        'should return remote data when the call to remote data source is successful',
        () async {
          // Arrange
          when(mockRemoteDataSource.getAllProduct())
              .thenAnswer((_) async => [testProduct]);

          // Act
          final result = await repository.getAllProduct();

          // Assert
          verify(mockRemoteDataSource.getAllProduct());
          expect(result, equals(const Right([testProduct])));
        },
      );

      test(
        'should return server failure when the call to remote data source is unsuccessful',
        () async {
          // Arrange
          when(mockRemoteDataSource.getAllProduct())
              .thenThrow(ServerException());

          // Act
          final result = await repository.getAllProduct();

          // Assert
          verify(mockRemoteDataSource.getAllProduct());
          expect(result, equals(Left(ServerFailure('Server failure'))));
        },
      );
    }

    runTestsOffline(() {
      test(
        'should return network failure when get all product is called',
        () async {
          // Arrange
          when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

          // Act
          final result = await repository.getAllProduct();

          // Assert
          verifyZeroInteractions(mockRemoteDataSource);
          verifyZeroInteractions(mockLocalDataSource);
          expect(result, equals(Left(NetworkFailure('Network failure'))));
        },
      );
    });
  });
}
