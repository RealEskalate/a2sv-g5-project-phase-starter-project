import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/exception.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/core/network/network_info.dart';
import 'package:e_commerce_app/features/product/data/data_sources/product_local_data_source.dart';
import 'package:e_commerce_app/features/product/data/data_sources/product_remote_data_source.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/repositories/product_repository.dart';

class ProductRepositoryImplimentation extends ProductRepository {
  final ProductRemoteDataSource productRemoteDataSource;
  final NetworkInfo networkInfo;
  final ProductLocalDatasource productLocalDataSource;
  ProductRepositoryImplimentation(
      {required this.networkInfo,
      required this.productLocalDataSource,
      required this.productRemoteDataSource});

  @override
  Future<Either<Failure, List<ProductEntity>>> getAllProduct() async {
    // TODO: implement getAllProduct
    // throw UnimplementedError();
    bool isConnected = await networkInfo.isConnected;
    if (isConnected) {
      try {
        final result = await productRemoteDataSource.getAllProduct();
        productLocalDataSource.cachecProducts(result);
        List<ProductEntity> resultList =
            result.map((prod) => prod.toEntity()).toList();
        return Right(resultList);
      } on ServerException {
        print("failing");
        return Left(ServerFailure("message"));
      } on SocketException {
        return Left(ConnectionFailure("Failed to connect to internet"));
      }
    } else {
      try {
        final result = await productLocalDataSource.getAllProduct();
        List<ProductEntity> resultList =
            result.map((prod) => prod.toEntity()).toList();
        return Right(resultList);
      } catch (e) {
        print(e);
        return Left(ServerFailure("no cahed product"));
      }
    }
  }

  @override
  Future<Either<Failure, ProductEntity>> getOneProduct(String id) async {
    // // TODO: implement getOneProduct
    // throw UnimplementedError();
    try {
      final result = await productRemoteDataSource.getOneProduct(id);
      return Right(result);
    } on ServerFailure {
      return Left(ServerFailure("Error while getting one produt"));
    } on SocketException {
      return Left(ServerFailure("Noconnection"));
    }
  }

  @override
  Future<Either<Failure, ProductEntity>> insertProduct(
      ProductEntity newProduct) async {
    // // TODO: implement insertProduct
    // throw UnimplementedError();
    print("hello");

    try {
      var newm = newProduct.toModel();
      final result = await productRemoteDataSource.insertProduct(newm);
      return Right(result);
    } on ServerFailure {
      return Left(ServerFailure("Error while getting all produt"));
    } on SocketException {
      return Left(ServerFailure("No connection"));
    }
  }

  @override
  Future<Either<Failure, String>> deleteProduct(String id) async {
    // TODO: implement deleteProduct
    // throw UnimplementedError();
    try {
      final result = await productRemoteDataSource.deleteProduct(id);
      return Right(result);
    } on ServerFailure {
      return Left(ServerFailure("Error while getting all produt"));
    } on SocketException {
      return Left(ServerFailure("No connection"));
    }
  }

  @override
  Future<Either<Failure, ProductEntity>> updateProduct(
      String id, ProductEntity updatedProduct) async {
    // TODO: implement updateProduct
    // throw UnimplementedError();
    try {
      final result =
          await productRemoteDataSource.updateProduct(updatedProduct.toModel());
      return Right(result);
    } on ServerFailure {
      return Left(ServerFailure("Error while getting all produt"));
    } on SocketException {
      return Left(ServerFailure("No connection"));
    }
  }
}
