import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import '../../../../core/errors/failure/failures.dart';
import '../../domain/repositories/product_repository.dart';
import '../data_sources/local_data_source/local_data_source.dart';
import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';

import '../models/product_models.dart';


class ProductRepositoryImpl implements ProductRepository {
  final ProductRemoteDataSource remoteDataSource;
  final ProductLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  ProductRepositoryImpl(this.remoteDataSource, this.localDataSource, this.networkInfo);

  @override
  Future<Either<Failure,List<ProductModel>>> getProducts() async {
     if(await networkInfo.isConnected){
       try{
        final products = await remoteDataSource.getProducts();
        return Right(products);
      } on ServerFailure {
        return Left(ServerFailure("Server Failure"));
      }
     }else {
      try{
        final products = await localDataSource.getProducts();
        return Right(products);
      }on LocalDataSourceFailure {
        return Left(LocalDataSourceFailure('Failed to get products'));
      }
     }
  }
  
  @override
  Future<Either<Failure,UserModel>> getUserInfo() async {
     if(await networkInfo.isConnected){
       try{
        final user = await remoteDataSource.getUserInfo();
        return Right(user);
      } on ServerFailure {
        return Left(ServerFailure("Server Failure"));
      }
     }else {
      return Left(ServerFailure('No User Found'));
     }
  }// Get User info




  @override
  Future<Either<Failure,ProductModel>> getProduct(String productId) async {
    if(await networkInfo.isConnected){
       try{
        final product = await remoteDataSource.getProduct(productId);
        localDataSource.addProduct(product);
        return Right(product);
      } on ServerFailure {
        return Left(ServerFailure('Server: Failed to get product'));
      }
     }else {
      try{
        final product = await localDataSource.getProduct(productId);
        return Right(product);
      }on LocalDataSourceFailure {
        return Left(LocalDataSourceFailure('Failed to get product'));
      }
     }
  }// Get Product impl

  @override
  Future<Either<Failure,String>> addProduct(ProductEntity product) async {
    try{
        final results = await remoteDataSource.addProduct(product);
        return Right(results);
      } catch(e) {
        return Left(ServerFailure(e.toString()));
      } 
  } // Add Product impl




  @override
  Future<Either<Failure,String>> updateProduct(ProductModel product) async {
    try{
        final result = await remoteDataSource.updateProduct(product);
        return Right(result);
      } on Exception {
        return Left(ServerFailure('Server: Failed to update product'));
      } 
      } // Update Product impl

  @override
  Future<Either<Failure,String>> deleteProduct(String productId) async {
    try{
        final products = await remoteDataSource.deleteProduct(
          productId);
        return Right(products);
      } on Exception {
        return Left(ServerFailure('Server: Failed to delete product'));
      } 
      } // Delete Product impl
  
}
