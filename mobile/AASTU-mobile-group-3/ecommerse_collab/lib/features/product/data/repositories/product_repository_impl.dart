import 'dart:io';

import 'package:dartz/dartz.dart';

import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/network/network_info.dart';
import '../../domain/entity/product.dart';
import '../../domain/repository/productRepository.dart';
import '../data_sources/local_data_source.dart';
import '../data_sources/remote_data_source.dart';


class ProductRepositoryImpl extends ProductRepository{
  final NetworkInfo networkInfo;
  final LocalDataSource localDataSource;
  final ProductRemoteDataSource productRemoteDataSource;
  
  ProductRepositoryImpl({required this.productRemoteDataSource, required this.localDataSource,required this.networkInfo});

  @override
  Future<Either<Failure, void>> addProduct(Product product) async {
    try{
 
      return Right(await productRemoteDataSource.addProduct(product.toModel()));

    }on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    }on SocketException {
      return const Left(ConnectionFailure('Failed to connect'));
    }
  }

  

  @override
  Future<Either<Failure, List<Product>>> getAllProduct() async {
    List<Product> products = [];
    try{
      if (await networkInfo.isConnected){
          final cacheProducts = await productRemoteDataSource.getAllProduct();
          localDataSource.cacheProducts(cacheProducts);
          products = cacheProducts.map((model) => model.toEntity()).toList();
          return Right(products);         
          
      }else{
       products = localDataSource.getAllProduct().map((model) => model.toEntity()).toList();
        return Right(products);
      }
    }on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    }on SocketException {
      return const Left(ConnectionFailure('Failed to connect'));
    }
  }

  @override
  Future<Either<Failure, Product>> getProduct(String id) async {
    try{
      if (await networkInfo.isConnected){
        final product = await productRemoteDataSource.getProduct(id);
        return Right(product.toEntity());
      }else{
        return Right(localDataSource.getProduct(id).toEntity());
      }
    }on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    }on SocketException {
      return const Left(ConnectionFailure('Failed to connect'));
    }
  }

  @override
  Future<Either<Failure, Product>> updateProduct({required String id, required String name, required double price, required String description}) async {
   try{
    return Right(await productRemoteDataSource.updateProduct(id: id, name: name, price: price, description: description));

   }on ServerException {
      return const Left(ServerFailure('An error has occurred'));
    }on SocketException {
      return const Left(ConnectionFailure('Failed to connect'));
    }
  }
  
  @override
  Future<Either<Failure, void>> deleteProduct(String id) async {
    try{
      return Right(await productRemoteDataSource.deleteProduct(id));      

    } catch(e){
      return const Left(ServerFailure('An error has occurred'));
    }
    
  }
  
}