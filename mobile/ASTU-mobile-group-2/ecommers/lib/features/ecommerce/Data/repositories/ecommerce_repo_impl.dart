import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:file_picker/file_picker.dart';

import '../../../../core/Error/failure.dart';
import '../../../../core/network/check_connectivity.dart';
import '../../Domain/entity/ecommerce_entity.dart';
import '../../Domain/repositories/ecommerce_repositories.dart';
import '../data_source/local_data_source.dart';
import '../data_source/remote_data_source.dart';
import '../model/ecommerce_model.dart';

class EcommerceRepoImpl implements EcommerceRepositories  {
  final EcommerceRemoteDataSource remoteDataSource;
  final NetworkInfo networkInfo;
  final LocalDataSource localDataSource;
  const EcommerceRepoImpl({
    required this.remoteDataSource,
    required this.networkInfo,
    required this.localDataSource
  });

  @override
  Future<Either<Failure, bool>> addProduct(Map<String,dynamic> product) async{
    try {
      final connection = await networkInfo.isConnected;
      if(connection == false){
        return const Left(ConnectionFailur(message: 'connection error'));
      }

     
      final result = await remoteDataSource.addProduct(product);
      return Right(result);
    } on ConnectionFailur {
      return const Left(ConnectionFailur(message: 'connection error'));
    }
  }

  @override
  Future<Either<Failure, bool>> deleteProduct(String id) async{
    try {
      final connection = await networkInfo.isConnected;
      if(connection == false){
        return const Left(ConnectionFailur(message: 'connection error'));
      }
      final result = await remoteDataSource.deleteProduct(id);
      return Right(result);
    } on ConnectionFailur {
      return const Left(ConnectionFailur(message: 'connection error'));
    }
  }

  @override
  Future<Either<Failure, bool>> editeProduct(String id,Map<String,dynamic> product) async{
    try {
      final connection = await networkInfo.isConnected;
      if(connection == false){
        return const Left(ConnectionFailur(message: 'connection error'));
      }
      final result = await remoteDataSource.editProduct(id,product);
      return Right(result);
    } on ConnectionFailur {
      return const Left(ConnectionFailur(message: 'connection error'));
    }
  }

  @override
  Future<Either<Failure, List<EcommerceEntity>>> getAllProduct() async {
    try {
      final connection = await networkInfo.isConnected;
      if(connection == false){
        return const Left(ConnectionFailur(message: 'connection error'));
      }
      final result = await remoteDataSource.getAllProducts();
      final entities = EcommerceModel.listToEntity(result);
      return Right(entities);
    } on ServerFailure {
      return const Left(ServerFailure(message: 'server Error'));
    } on ConnectionFailur {
      return const Left(ConnectionFailur(message: 'Connection Error'));
    }
  }

  @override
  Future<Either<Failure, EcommerceEntity>> getProductById(String id) async {
    try {
      final connection = await networkInfo.isConnected;
      if(connection == false){
        return const Left(ConnectionFailur(message: 'connection error'));
      }
      final result = await remoteDataSource.getProduct(id);
      final entities = result.toEntity();
      return Right(entities);
    } on ServerFailure {
      return const Left(ServerFailure(message: 'server Error'));
    } on ConnectionFailur {
      return const Left(ConnectionFailur(message: 'Connection Error'));
    }
  }

  @override
  Future<Either<Failure, Map<String,dynamic>>> selectImage() async{
    try {
  
    final result = await FilePicker.platform.pickFiles(
      type: FileType.image, // Restrict to image files
    );
    if (result == null || result.files.isEmpty) {
      return const Left(ServerFailure(message: 'try again'));
    }
    final filePath = result.files.single.path;

    if (filePath == null) {
      return const Left(ServerFailure(message: 'try again'));
    }
    final Map<String,dynamic> data = {
      'image': filePath,
      'file': File(filePath)
    };
    return Right(data);
  } catch (e) {
    
    return Left(ServerFailure(message: e.toString()));
  }
  }
  
  @override
  Future<String> getUserName(String key) async{
    try {
      final result = await localDataSource.getName(key);
      if (result.isNotEmpty) {
        return result;
      }
      return 'Geust';
    } catch (e) {
      return 'Geust';
    }
  }
  
  @override
  Future<bool> logoutUser(String key) async{
    try {
      final result =  await localDataSource.deleteToken(key);
      return result;
    } catch (e) {
      return false;
    }
    
  }


  
  
  
}
