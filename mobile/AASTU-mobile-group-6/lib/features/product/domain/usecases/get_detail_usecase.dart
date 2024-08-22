import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import '../../../../core/errors/failure/failures.dart';
import '../../../../core/usecases/usecases.dart';
import '../repositories/product_repository.dart';
import 'package:equatable/equatable.dart';
import '../entities/product_entity.dart';

class GetDetailUseCase implements UseCase<ProductModel, String> {
  final ProductRepository abstractProductRepository;

  GetDetailUseCase(this.abstractProductRepository);

  @override
  Future<Either<Failure,ProductModel>> call(String id) async {
    return await abstractProductRepository.getProduct(id);
  }
}