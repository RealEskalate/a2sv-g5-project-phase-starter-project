import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:flutter/foundation.dart';

import '../../../../core/exception/exception.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/platform/network_info.dart';
import '../../domain/entities/product_entity.dart';
import '../../domain/repositories/product_repository.dart';
import '../data_source/local_data_source/product_local_data_source.dart';
import '../data_source/remote_data_source/product_remote_data_source.dart';
import '../models/product_model.dart';

class ProductRepositoryImpl implements ProductRepositories {
  late final ProductRemoteDataSource productRemoteDataSource;
  late final ProductLocalDataSource productLocalDataSource;
  late final NetworkInfo networkInfo;
  ProductRepositoryImpl(
      {required this.productRemoteDataSource,
      required this.productLocalDataSource,
      required this.networkInfo});

  @override
  Future<Either<Failure, Product>> createProduct(Product product) async {

    if (await networkInfo.isConnected) {
      try {
        final result = await productRemoteDataSource
            .createProduct(ProductModel.toModel(product));
        return Right(ProductModel.toEntity(result));
      } on ServerException {
        return const Left(ServerFailure(message: 'An error occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure(message: 'Failed to connect to the internet'));
      }
    } else {
      return const Left(
          ConnectionFailure(message: 'Failed to connect to the internet'));
    }
  }

  @override
  Future<Either<Failure, void>> deleteProduct(String id) async {
    
    if (await networkInfo.isConnected) {
      try {
        final result = await productRemoteDataSource.deleteProduct(id);
        return Right(result);
      } on ServerException {
        return const Left(ServerFailure(message: 'An error occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure(message: 'Failed to connect to the internet'));
      }
    } else {
      return const Left(
          ConnectionFailure(message: 'Failed to connect to the internet'));
    }
  }

  @override
  Future<Either<Failure, Product>> getProduct(String id) async {
  
    if (await networkInfo.isConnected) {
      try {
        final result = await productRemoteDataSource.getProductById(id);
        return Right(ProductModel.toEntity(result));
      } on ServerException {
        return const Left(ServerFailure(message: 'An error occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure(message: 'Failed to connect to the internet'));
      }
    } else {
      return const Left(
          ConnectionFailure(message: 'Failed to connect to the internet'));
    }
  }

  @override
  Future<Either<Failure, List<Product>>> getProducts() async {
  
    if (await networkInfo.isConnected) {
      try {
        final result = await productRemoteDataSource.getProducts();
        try {
          await productLocalDataSource.cacheProducts(result);
        } on CacheException {
          debugPrint('Caching All Products error');
        }
        return Right(ProductModel.toEntityList(result));
      } on ServerException {
        return const Left(ServerFailure(message: 'An error occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure(message: 'Failed to connect to the internet'));
      }
    } else {
      try {
        final result = await productLocalDataSource.getProducts();
        return Right(ProductModel.toEntityList(result));
      } on CacheException {
        return const Left(CacheFailure(message: 'An error occurred'));
      }
    }
  }

  @override
  Future<Either<Failure, Product>> updateProduct(Product product) async {
    
    if (await networkInfo.isConnected) {
      try {
        final result = await productRemoteDataSource
            .updateProduct(ProductModel.toModel(product));
        return Right(ProductModel.toEntity(result));
      } on ServerException {
        return const Left(ServerFailure(message: 'An error occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure(message: 'Failed to connect to the internet'));
      }
    } else {
      return const Left(
          ConnectionFailure(message: 'Failed to connect to the internet'));
    }
  }
}
