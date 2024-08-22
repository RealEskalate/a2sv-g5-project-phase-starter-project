import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import '../../../../core/errors/failure/failures.dart';
import '../../../../core/usecases/usecases.dart';
import '../repositories/product_repository.dart';
import 'package:equatable/equatable.dart';
import '../entities/product_entity.dart';

class GetUserInfo implements UseCase<UserModel, NoParams> {
  final ProductRepository abstractProductRepository;

  GetUserInfo(this.abstractProductRepository);

  @override
  Future<Either<Failure,UserModel>> call(NoParams params) async {
    return await abstractProductRepository.getUserInfo();
  }
}