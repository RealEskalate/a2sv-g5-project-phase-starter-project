import 'dart:io';

import 'package:dartz/dartz.dart';

import '../../../../core/connections/network_info.dart';
import '../../../../core/errors/exception.dart';
import '../../../../core/errors/failure.dart';
import '../../domain/entities/product.dart';
import '../../domain/repository/product_repository.dart';
import '../data_sources/product_local_data_source.dart';
import '../data_sources/product_remote_datasource.dart';
import '../models/product_model.dart';

class ProductRepositoryImp extends ProductRepository {
  final ProductRemoteDatasource productRemoteDatasource;
  final ProductLocalDataSource productLocalDataSource;
  final NetworkInfo networkInfo;

  ProductRepositoryImp({
    required this.productRemoteDatasource,
    required this.networkInfo,
    required this.productLocalDataSource,
  });

  @override
  Future<Either<Failure, Product>> createProduct(Product product) async {
    final isConnected = await networkInfo.isConnected;
    if (isConnected == true) {
      try {
        final productModel = ProductModel(
          id: product.id,
          name: product.name,
          description: product.description,
          imageUrl: product.imageUrl,
          price: product.price,
        );
        final result =
            await productRemoteDatasource.createProduct(productModel);
        return result.fold(
          (failure) => left(failure),
          (model) => right(model.toEntity()),
        );
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure('Failed to connect with the internet.'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection.'));
    }
  }

  @override
  Future<Either<Failure, void>> deleteProduct(String id) async {
    final isConnected = await networkInfo.isConnected;
    if (isConnected == true) {
      try {
        final result = await productRemoteDatasource.deleteProduct(id);
        return result.fold(
          (failure) => left(failure),
          (_) => right(null),
        );
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure('Failed to connect with the internet.'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection.'));
    }
  }

  @override
  Future<Either<Failure, Product>> getProductById(String id) async {
    final isConnected = await networkInfo.isConnected;
    if (isConnected == true) {
      try {
        final result = await productRemoteDatasource.getProductById(id);
        return result.fold(
          (failure) => left(failure),
          (model) => right(model.toEntity()),
        );
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure('Failed to connect with the internet.'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection.'));
    }
  }

  @override
  Future<Either<Failure, List<Product>>> getProducts() async {
    final isConnected = await networkInfo.isConnected;
    if (isConnected == true) {
      try {
        final result = await productRemoteDatasource.getProducts();
        return result.fold(
          (failure) => left(failure),
          (models) => right(models.map((model) => model.toEntity()).toList()),
        );
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure('Failed to connect with the internet.'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection.'));
    }
  }

  @override
  Future<Either<Failure, Product>> updateProduct(Product product, String id) async {
    final isConnected = await networkInfo.isConnected;
    if (isConnected == true) {
      try {
        final productModel = ProductModel(
          id: product.id,
          name: product.name,
          description: product.description,
          imageUrl: product.imageUrl,
          price: product.price,
        );
        final result =
            await productRemoteDatasource.updateProduct(productModel);
        return result.fold(
          (failure) => left(failure),
          (model) => right(model.toEntity()),
        );
      } on ServerException {
        return const Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return const Left(
            ConnectionFailure('Failed to connect with the internet.'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection.'));
    }
  }
}
