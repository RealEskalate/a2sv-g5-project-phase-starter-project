import 'package:dartz/dartz.dart';

import '../../../../core/error/exceptions.dart';
import '../../../../core/error/failures.dart';
import '../../../../core/platform/network_info.dart';
import '../../domain/entitity/product.dart';
import '../../domain/repositories/product_repository.dart';
import '../datasources/local_data_resource.dart';
import '../datasources/remote_data_source.dart';
import '../model/product_model.dart';

class ProductRepositoryImpl implements ProductRepository {
  final ProductRemoteDataSource remoteDataSource;
  final ProductLocalDataSource localDataSource;
  final NetworkInfo networkInfo;
  ProductRepositoryImpl(
      {required this.remoteDataSource,
      required this.localDataSource,
      required this.networkInfo});

  @override
  Future<Either<Failure, Product>> getProduct(String id) async {
    if (await networkInfo.isConnected) {
      try {
        final remoteProduct = await remoteDataSource.getProduct(id);
        localDataSource.cacheProduct(ProductModel.fromProduct(remoteProduct));
        return Right(remoteProduct);
      } on ServerException {
        return Left(ServerFailure('Server failure'));
      }
    } else {
      try {
        final localProduct = await localDataSource.getLastProduct();
        return Right(localProduct);
      } on CacheException {
        return Left(CacheFailure('Cache failure'));
      }
    }
  }

  @override
  Future<Either<Failure, Product>> insertProduct(Product product) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource
            .insertProduct(ProductModel.fromProduct(product));
        return Right(result);
      } on ServerException {
        return Left(ServerFailure('Server failure'));
      }
    } else {
      return Left(NetworkFailure('Network failure'));
    }
  }

  @override
  Future<Either<Failure, Product>> deleteProduct(String id) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.deleteProduct(id);
        return Right(result);
      } on ServerException {
        return Left(ServerFailure('Server failure'));
      }
    } else {
      return Left(NetworkFailure('Network failure'));
    }
  }

  @override
  Future<Either<Failure, Product>> updateProduct(Product product) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource
            .updateProduct(ProductModel.fromProduct(product));
        return Right(result);
      } on ServerException {
        return Left(ServerFailure('Server failure'));
      }
    } else {
      return Left(NetworkFailure('Network failure'));
    }
  }

  @override
  Future<Either<Failure, List<Product>>> getAllProduct() async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.getAllProduct();
        return Right(result);
      } on ServerException {
        return Left(ServerFailure('Server failure'));
      }
    } else {
      return Left(NetworkFailure('Network failure'));
    }
  }
}
