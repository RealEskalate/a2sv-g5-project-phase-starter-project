
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/core/connectivity/network_info.dart';
import 'package:task_9/core/error/exceptions.dart';
import 'package:task_9/core/failure/failure.dart';
import 'package:task_9/features/product/data/data_sources/product_local_data_source.dart';
import 'package:task_9/features/product/data/data_sources/product_remote_data_source.dart';
import 'package:task_9/features/product/data/models/product_model.dart';
import 'package:task_9/features/product/data/repositories/product_repository_impl.dart';

import 'product_repository_impl_test.mocks.dart';

@GenerateMocks([ProductRemoteDataSource, ProductLocalDataSource, NetworkInfo])
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
  const tProductModel = ProductModel(
    id: '6672776eb905525c145fe0bb',
    name: 'Anime website',
    description: 'Explore anime characters.',
    price: 123,
    imageUrl:
        'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777711/images/clmxnecvavxfvrz9by4w.jpg',
  );

  const tImagePath = 'assets/images/boots.jpg';
  group('addProduct', () {
    test('should add product from remote and local data sources when online',
        () async {
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.addProduct(any, any))
          .thenAnswer((_) async => tProductModel);
      when(mockLocalDataSource.addProduct(any))
          .thenAnswer((_) async => tProductModel);
      final result = await repository.addProduct(tProductModel, tImagePath);
      verify(mockRemoteDataSource.addProduct(tProductModel, tImagePath));
      verify(mockLocalDataSource.addProduct(
        tProductModel,
      ));
      expect(result, const Right(tProductModel));
    });
    test('should add product from local data source only when offline',
        () async {
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      when(mockLocalDataSource.addProduct(any))
          .thenAnswer((_) async => tProductModel);
      await repository.addProduct(tProductModel, tImagePath);
      verifyNever(mockRemoteDataSource.addProduct(tProductModel, tImagePath));
      verify(mockLocalDataSource.addProduct(
        tProductModel,

      ));
    });
  });

  group('updateProduct', () {
    test('should update product from remote and local data sources when online',
        () async {
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.updateProduct(any, any))
          .thenAnswer((_) async => tProductModel);
      when(mockLocalDataSource.updateProduct(any))
          .thenAnswer((_) async => tProductModel);
      final result = await repository.updateProduct(tProductModel);
      verify(
          mockRemoteDataSource.updateProduct(tProductModel.id, tProductModel));
      verify(mockLocalDataSource.updateProduct(tProductModel));
      expect(result, const Right(tProductModel));
      
    });
    test('should update product in local data source only when offline',
        () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      when(mockLocalDataSource.updateProduct(any))
          .thenAnswer((_) async => (tProductModel));
      // Act
      await repository.updateProduct(tProductModel);
      // Assert
      verifyNever(
          mockRemoteDataSource.updateProduct(tProductModel.id, tProductModel));
      verify(mockLocalDataSource.updateProduct(tProductModel));
    });
  });
  group('deleteProduct', () {
    test('should delete product from remote and local data sources when online',
        () async {
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.deleteProduct(any)).thenAnswer((_) async => {});
      when(mockLocalDataSource.deleteProduct(any)).thenAnswer((_) async => {});
      await repository.deleteProduct(tProductModel.id);
      verify(mockRemoteDataSource.deleteProduct(tProductModel.id));
      verify(mockLocalDataSource.deleteProduct(tProductModel.id));
    });
    test('should delete product from local data source only when offline',
        () async {
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      when(mockLocalDataSource.deleteProduct(any)).thenAnswer((_) async => {});
      await repository.deleteProduct(tProductModel.id);
      verifyNever(mockRemoteDataSource.deleteProduct(tProductModel.id));
      verify(mockLocalDataSource.deleteProduct(tProductModel.id));
    });
  });
  group('getAllProducts', () {
    final tProductList = [tProductModel];

    test('should return remote data when online', () async {
      // arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.getAllProducts())
          .thenAnswer((_) async => tProductList);
      when(mockLocalDataSource.cacheProducts(any)).thenAnswer((_) async => {});

      // act
      final result = await repository.getAllProducts();

      // assert
      verify(mockRemoteDataSource.getAllProducts());
      verify(mockLocalDataSource.cacheProducts(tProductList));
      expect(result, Right(tProductList));
    });

    test('should return local data when offline', () async {
      // arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      when(mockLocalDataSource.getAllProducts())
          .thenAnswer((_) async => tProductList);

      // act
      final result = await repository.getAllProducts();

      // assert
      verifyNever(mockRemoteDataSource.getAllProducts());
      verify(mockLocalDataSource.getAllProducts());
      expect(result, Right(tProductList));
    });
  });
  group('getProductById', () {
    const tId = '6672776eb905525c145fe0bb';

    test('should return remote data when online', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.getProductById(any))
          .thenAnswer((_) async => tProductModel);
      when(mockLocalDataSource.cacheProducts(any)).thenAnswer((_) async => {});
      // Act
      final result = await repository.getProductById(tId);
      // Assert
      verify(mockRemoteDataSource.getProductById(tId));
      verify(mockLocalDataSource.cacheProducts([tProductModel]));
      expect(result, const Right(tProductModel));
    });

    test('should return local data when offline', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
      when(mockLocalDataSource.getProductById(any))
          .thenAnswer((_) async => tProductModel);
      // Act
      final result = await repository.getProductById(tId);
      // Assert
      verifyNever(mockRemoteDataSource.getProductById(tId));
      verify(mockLocalDataSource.getProductById(tId));
      expect(result, const Right(tProductModel));
    });

    test('should return a ServerFailure when there is an exception', () async {
      // Arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockRemoteDataSource.getProductById(any))
          .thenThrow(ServerException());
      // Act
      final result = await repository.getProductById(tId);
      // Assert
      expect(result,
          equals(Left(ServerFailure(message: 'Failed to fetch product.'))));
    });
  });
}
