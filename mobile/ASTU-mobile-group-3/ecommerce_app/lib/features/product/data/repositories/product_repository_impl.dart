import 'package:dartz/dartz.dart';
import '../../../../core/constants/constants.dart';
import '../../../../core/errors/exceptions/product_exceptions.dart';
import '../../../../core/errors/failures/failure.dart';
import '../../../../core/network/network_info.dart';
import '../../domain/entities/product.dart';

import '../../domain/repositories/product_repository.dart';
import '../data_resources/local_product_data_source.dart';
import '../data_resources/remote_product_data_source.dart';
import '../models/product_model.dart';

class ProductRepositoryImpl implements ProductRepository {
  final RemoteProductDataSource remoteProductDataSource;
  final LocalProductDataSource localProductDataSource;
  final NetworkInfo networkInfo;

  ProductRepositoryImpl({
    required this.remoteProductDataSource,
    required this.localProductDataSource,
    required this.networkInfo,
  });

  /// Deletes a product from list base on the [id]
  /// Returns [ServerFailure] if the request failed
  ///
  /// Return [ConnectionFailure] if the device is offline
  @override
  Future<Either<Failure, int>> deleteProduct(String id) async {
    final network = await networkInfo.isConnected;
    if (network) {
      try {
        final result = await remoteProductDataSource.deleteProduct(id);
        await localProductDataSource.removeProduct(id);
        return Right(result);
      } on ServerException {
        return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
      } on CacheException {
        return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
      }
    } else {
      return Left(
          ConnectionFailure(AppData.getMessage(AppData.connectionError)));
    }
  }

  /// Retreive data from if the device is online
  /// Returns [ServerFailure] if failed
  ///
  /// Retreive data from local repository if device is offline
  /// Rrturns [CacheFailure] if there is no data on local repository
  @override
  Future<Either<Failure, List<ProductEntity>>> getAllProducts() async {
    final network = await networkInfo.isConnected;
    if (network) {
      try {
        final result = await remoteProductDataSource.getAllProducts();

        await localProductDataSource.addListOfProduct(result);
        return Right(ProductModel.allToEntity(result));
      } on ServerException {
        return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
      } on CacheException {
        return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
      }
    } else {
      try {
        final result = await localProductDataSource.getAllProducts();
        return Right(ProductModel.allToEntity(result));
      } on CacheException {
        return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
      }
    }
  }

  /// Retrive single product based on id from remote if the device is onlin
  /// Returns [ServerFailure] if failed
  ///
  /// Retrive the product with that i if the data is on local repository
  /// Return [CacheFailure] if the data is not on local repo
  @override
  Future<Either<Failure, ProductEntity>> getProduct(String id) async {
    final network = await networkInfo.isConnected;
    if (network) {
      try {
        final result = await remoteProductDataSource.getProduct(id);
        await localProductDataSource.addProduct(result);
        return Right(result.toEntity());
      } on ServerException {
        return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
      } on CacheException {
        return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
      }
    } else {
      try {
        final result = await localProductDataSource.getProduct(id);
        return Right(result.toEntity());
      } on CacheException {
        return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
      }
    }
  }

  /// Insert provided product into a database
  /// Returns [ServerFailure] if the request is failed
  ///
  /// Returns [ConnectionFailure] if the device is offline
  @override
  Future<Either<Failure, int>> insertProduct(ProductEntity product) async {
    final network = await networkInfo.isConnected;
    if (network) {
      try {
        final result = await remoteProductDataSource
            .insertProduct(ProductModel.fromEntity(product));
        await localProductDataSource
            .addProduct(ProductModel.fromEntity(product));
        return Right(result);
      } on ServerException {
        return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
      } on CacheException {
        return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
      }
    } else {
      return Left(
          ConnectionFailure(AppData.getMessage(AppData.connectionError)));
    }
  }

  /// Update the product based on the id inside the provided product
  /// Returns [ServerFailure] if the request is failed on the way
  ///
  /// Return [ConnectionFailure] if the device is not connected to internet
  @override
  Future<Either<Failure, int>> updateProduct(ProductEntity product) async {
    final network = await networkInfo.isConnected;
    if (network) {
      try {
        final result = await remoteProductDataSource
            .updateProduct(ProductModel.fromEntity(product));
        await localProductDataSource
            .updateProduct(ProductModel.fromEntity(product));
        return Right(result);
      } on ServerException {
        return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
      } on CacheFailure {
        return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
      }
    } else {
      return Left(
          ConnectionFailure(AppData.getMessage(AppData.connectionError)));
    }
  }
}
