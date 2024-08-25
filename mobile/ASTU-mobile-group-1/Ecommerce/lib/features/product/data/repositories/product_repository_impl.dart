import 'dart:io';

import 'package:dartz/dartz.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/network/network_info.dart';
import '../../domain/entities/product_entity.dart';

import '../../domain/repositories/product_repository.dart';
import '../data_sources/local_data_source.dart';
import '../data_sources/remote_data_source.dart';

class ProductRepositoryImpl extends ProductRepository {
  final ProductRemoteDataSource remoteDataSource;
  final ProductLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  ProductRepositoryImpl(
      {required this.remoteDataSource,
      required this.localDataSource,
      required this.networkInfo});

  @override
  Future<Either<Failure, bool>> deleteProduct(String id) async {
    try {
      final res = await remoteDataSource.deleteProduct(id);
      return Right(res);
    } on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    } on SocketException {
      return const Left(ConnectionFailure());
    }
  }

  @override
  Future<Either<Failure, List<ProductEntity>>> getAllProducts() async {
    try {
      final products = await remoteDataSource.getAllProducts();
      return Right(products.map((product) => product.toEntity()).toList());
    } on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    } on SocketException {
      return const Left(ConnectionFailure());
    }
  }

  @override
  Future<Either<Failure, ProductEntity>> getProductById(String id) async {
    try {
      final product = await remoteDataSource.getProductById(id);
      return Right(product.toEntity());
    } on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    } on SocketException {
      return const Left(ConnectionFailure());
    }
  }

  @override
  Future<Either<Failure, List<ProductEntity>>> getProductsByCategory(
      String category) {
    // TODO: implement getProductsByCategory
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, bool>> insertProduct(ProductEntity product) async {
    if (await networkInfo.isConnected) {
      try {
        final call = await remoteDataSource.insertProduct(product.toModel());

        return Right(call);
      } on ServerException catch (e) {
        return Left(ServerFailure(e.message));
      } catch (e) {
        return const Left(UnkownFailure());
      }
    } else {
      throw const ConnectionFailure();
    }
  }

  @override
  Future<Either<Failure, bool>> updateProduct(ProductEntity product) async {
    try {
      final call = await remoteDataSource.updateProduct(product.toModel());

      return Right(call);
    } on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    } on SocketException {
      return const Left(ConnectionFailure());
    }
  }
}
