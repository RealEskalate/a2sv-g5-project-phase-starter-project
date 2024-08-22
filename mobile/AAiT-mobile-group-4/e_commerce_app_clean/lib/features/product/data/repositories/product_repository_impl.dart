import 'package:dartz/dartz.dart';

import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/network/network_info.dart';
import '../../domain/entities/product_entity.dart';
import '../../domain/repository/product_repository.dart';
import '../data_sources/local/local_data_source.dart';
import '../data_sources/remote/remote_data_source.dart';
import '../models/product_mapper.dart';
import '../models/product_model.dart';

class ProductRepositoryImpl implements ProductRepository {
  final ProductRemoteDataSource remoteDataSource;
  final ProductLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  ProductRepositoryImpl(
      {required this.remoteDataSource,
      required this.localDataSource,
      required this.networkInfo});
  @override
  Future<Either<Failure, ProductEntity>> addProduct(
      ProductEntity product) async {
    if (await networkInfo.isConnected) {
      try {
        final result =
            await remoteDataSource.addProduct(product.toProductModel());
        return Right(result.toProductEntity());
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, bool>> deleteProduct(String id) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.deleteProduct(id);
        return Right(result);
      } on ServerException {

        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, ProductEntity>> getProduct(String id) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.getProduct(id);
        return Right(result.toProductEntity());
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, List<ProductEntity>>> getProducts() async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.getProducts();
        try {
          await localDataSource.cacheProducts(result);
        } on CacheException {
          return const Left(CacheFailure('error while caching'));
        }
        return Right(ProductModel.toProductListEntity(result));
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      try {
        final result = await localDataSource.getProducts();
        return Right(ProductModel.toProductListEntity(result));
      } catch (e) {
        return const Left(CacheFailure('Unable to load cache'));
      }
    }
  }

  @override
  Future<Either<Failure, ProductEntity>> updateProduct(
      ProductEntity product) async {
    if (await networkInfo.isConnected) {
      try {
        final result =
            await remoteDataSource.updateProduct(product.toProductModel());
        return Right(result.toProductEntity());
      } on ServerException {
        
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }
}
