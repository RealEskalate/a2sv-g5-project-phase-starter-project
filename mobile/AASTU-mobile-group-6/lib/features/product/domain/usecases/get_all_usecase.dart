import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import '../../../../core/errors/failure/failures.dart';
import '../../../../core/usecases/usecases.dart';
import '../repositories/product_repository.dart';
import 'package:equatable/equatable.dart';
import '../entities/product_entity.dart';


class GetAllUsecase implements UseCase<List<ProductModel>, NoParams> {
  final ProductRepository abstractProductRepository;

  GetAllUsecase(this.abstractProductRepository);

  @override
  Future<Either<Failure, List<ProductModel>> > call(NoParams params) async {
    return await abstractProductRepository.getProducts();
  }
}