import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../repositories/user_repository.dart';

class LoginParams {
  late String email;
  late String password;

  LoginParams({required this.email, required this.password});
}


class LoginUser extends UseCase< String, LoginParams> {
  final UserRepository userRepository;

  LoginUser(this.userRepository);

  @override
  Future<Either<Failure, String>> call(LoginParams params) async {
    return userRepository.loginUser(params.email, params.password);
  }
}


/* class AddProductParams {
  final ProductModel product;
  final String imagePath;

  AddProductParams(this.product, this.imagePath);
}

class AddProduct extends UseCase<ProductModel, AddProductParams> {
  final ProductRepository repository;

  AddProduct(this.repository);

  @override
  Future<Either<Failure, ProductModel>> call(AddProductParams params) async {
    return repository.addProduct(params.product, params.imagePath);
  }
}*/