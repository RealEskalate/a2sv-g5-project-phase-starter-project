import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:flutter/foundation.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/platform/network_info.dart';
import '../../domain/entities/product.dart';
import '../../domain/repositories/product_repository.dart';
import '../datasources/product_local_data_source.dart';
import '../datasources/product_remote_data_source.dart';
import '../models/product_model.dart';

class ProductRepositoryImpl extends ProductRepository {
  final ProductRemoteDataSource remoteDataSource;
  final ProductLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  ProductRepositoryImpl(
      {required this.remoteDataSource,
      required this.localDataSource,
      required this.networkInfo});
  @override
  Future<Either<Failure, Product>> createProduct(Product product) async {
    if (await networkInfo.isConnected) {
      try {
        final result =
            await remoteDataSource.createProduct(ProductModel.toModel(product));
        return Right(result.toEntity());
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(
            ConnectionFailure(ErrorMessages.noInternet));
      } catch (e) {
        final errorMessage = e.toString();
        return Left(RandomFailure(errorMessage));
      }
    } else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }

  @override
  Future<Either<Failure, void>> deleteProduct(String id) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.deleteProduct(id);
        return Right(result);
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(
            ConnectionFailure(ErrorMessages.noInternet));
      } catch (e) {
        final errorMessage = e.toString();
        return Left(RandomFailure(errorMessage));
      }
    } else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }

  @override
  Future<Either<Failure, List<Product>>> getAllProducts() async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.getAllProducts();
        try {
          await localDataSource.cacheAllProducts(result);
        } on CacheException {
          debugPrint('Caching All Product Error');
        }

        return Right(ProductModel.toEntityList(result));
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(
            ConnectionFailure(ErrorMessages.noInternet));
      } catch (e) {
        final errorMessage = e.toString();
        return Left(RandomFailure(errorMessage));
      }
    } else {
      try {
        
        final result = await localDataSource.getAllProducts();
        return Right(ProductModel.toEntityList(result));
      } on CacheException {
        return const Left(CacheFailure(ErrorMessages.cacheError));
      } catch (e) {
        final errorMessage = e.toString();
        return Left(RandomFailure(errorMessage));
      }
    }
  }

  @override
  Future<Either<Failure, Product>> getCurrentProduct(String id) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource.getCurrentProduct(id);
        return Right(result.toEntity());
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(
            ConnectionFailure(ErrorMessages.noInternet));
      } catch (e) {
        final errorMessage = e.toString();
        return Left(RandomFailure(errorMessage));
      }
    } else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }

  @override
  Future<Either<Failure, Product>> updateProduct(Product product) async {
    if (await networkInfo.isConnected) {
      try {
        final result =
            await remoteDataSource.updateProduct(ProductModel.toModel(product));
        return Right(result.toEntity());
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(
            ConnectionFailure(ErrorMessages.noInternet));
      } catch (e) {
        final errorMessage = e.toString();
        return Left(RandomFailure(errorMessage));
      }
    } else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }
}
