import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/auth/data/model/user_model.dart';
import 'package:ecommerce/features/auth/domain/repository/user_repository.dart';

import '../../../../core/error/failure.dart';


class LoginUserUsecase {
  final UserRepository userrepository;

  LoginUserUsecase(this.userrepository);

  Future<Either<Failure, UserModel>> login(String email, String password){
      return userrepository.loginUser(email, password);
  }
}