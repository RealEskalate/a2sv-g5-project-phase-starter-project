import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import '../../../../core/errors/failure/failures.dart';
import '../../../../core/usecases/usecases.dart';
import '../repositories/product_repository.dart';
import 'package:equatable/equatable.dart';
import '../entities/product_entity.dart';

class UpdateUsecase implements UseCase<String,ProductModel>{
  final ProductRepository abstractProductRepository;
  UpdateUsecase(this.abstractProductRepository);

  @override
  Future<Either<Failure,String>> call(ProductModel product) async{
     return await abstractProductRepository.updateProduct(product); 
  }

}