import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_access.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/repository/users_repository.dart';

class LoginUsecase extends UseCase<String,UserEntity>{
  final UsersRepository usersRepository;
  LoginUsecase(this.usersRepository);

  @override
  Future<Either<Failure, String>> call(UserEntity user) async {
    return await usersRepository.loginUser(user);
  }
  
}