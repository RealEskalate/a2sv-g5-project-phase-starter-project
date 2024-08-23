
import 'package:dartz/dartz.dart';
import '../../../../core/Error/failure.dart';
import '../entity/ecommerce_entity.dart';

abstract class EcommerceRepositories{


  Future<Either<Failure,EcommerceEntity>> getProductById(String id);

  Future<Either<Failure,List<EcommerceEntity>>> getAllProduct();

  Future<Either<Failure,bool>> addProduct(Map<String,dynamic> product);

  Future<Either<Failure,bool>> editeProduct(String id,Map<String,dynamic> product);
  
  Future<Either<Failure,bool>> deleteProduct(String id);
  Future<Either<Failure, Map<String,dynamic>>> selectImage();


  Future<String> getUserName(String key);
  Future<bool> logoutUser(String key);

}