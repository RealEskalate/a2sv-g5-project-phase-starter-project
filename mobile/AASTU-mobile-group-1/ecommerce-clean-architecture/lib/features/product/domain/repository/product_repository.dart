import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/product/domain/entities/product.dart';

import '../../../../core/error/failure.dart';

abstract class ProductRepository{

 Future<Either<Failure,List<Productentity>>> getAllProducts() ;
 Future<Either<Failure,void>> addProduct(Productentity newproduct) ;
 Future<Either<Failure,Productentity>> updateProduct(Productentity newproduct) ;
 Future<Either<Failure,bool>> deleteProduct(String id) ;
 Future<Either<Failure,Productentity>> getProduct(String id) ;
}