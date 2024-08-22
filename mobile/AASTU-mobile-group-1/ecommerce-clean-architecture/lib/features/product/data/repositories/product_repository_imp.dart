import 'dart:io';

import 'package:dartz/dartz.dart';

import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/networkinfo.dart';
import '../../domain/entities/product.dart';
import '../../domain/repository/product_repository.dart';
import '../data_sources/local_data_source.dart';
import '../data_sources/remote_data_source.dart';
import '../model/product_model.dart';

class ProductRepositoryImpl extends ProductRepository {
  final ProductRemoteDataSource productRemoteDataSource;
  final localDataSource productLocalDataSource;
  final NetworkInfo networkInfo;

  ProductRepositoryImpl({
    required this.productRemoteDataSource,
    required this.productLocalDataSource,
    required this.networkInfo,
  });

  @override
  Future<Either<Failure, List<Productentity>>> getAllProducts() async {
    try {
      if (await networkInfo.isConnected) {
        final products = await productRemoteDataSource.getallproduct();
        productLocalDataSource.cachedproduct(products);
        return Right(products);
      } else {
        final products = productLocalDataSource.getallproduct();
        return Right(products);
      } 
    } on ServerException {
      return Left(ServerFailure('An error has occured'));
    } on SocketException {
      return Left(ServerFailure('No internet connection'));
    }
  }


  Future<Either<Failure, void>> addProduct(Productentity newProduct) async {
    try {
        final result = await productRemoteDataSource.addproduct(ProductModel.fromEntity(newProduct));
        return Right(result);
      } on ServerException {
      return Left(ServerFailure('An error has occured'));
    } on SocketException {
      return Left(ServerFailure('No internet connection'));
    }

  }
  Future<Either<Failure,Productentity>> updateProduct(Productentity newProduct) async {
    try{
      if(await networkInfo.isConnected){
        final result = await productRemoteDataSource.updateproduct(ProductModel.fromEntity(newProduct));
        return Right(result);
    }else{
      final result = await productLocalDataSource.updateproduct(ProductModel.fromEntity(newProduct));
      return Right(result);
    }
  }on ServerException {
      return Left(ServerFailure('An error has occured'));
    } on SocketException {
      return Left(ServerFailure('No internet connection'));
    }
  }

  Future<Either<Failure,bool>> deleteProduct(String id) async {
    try{
      if(await networkInfo.isConnected){
        final result = await productRemoteDataSource.deleteproduct(id);
        return Right(result);
    }
    else{
      final result = await productLocalDataSource.deleteproduct(id);
      return Right(result);
    }
  }on ServerException {
      return Left(ServerFailure('An error has occured'));
    } on SocketException {
      return Left(ServerFailure('No internet connection'));
    }
  }
  Future<Either<Failure,Productentity>> getProduct(String id) async {
    try{
      if(await networkInfo.isConnected){
        final result = await productRemoteDataSource.getproduct(id);
        return Right(result);
    }
    else{
      final result = await productLocalDataSource.getproduct(id);
      return Right(result);
    }}on ServerException {
      return Left(ServerFailure('An error has occured'));
    } on SocketException {
      return Left(ServerFailure('No internet connection'));
    }
  }
  }
  