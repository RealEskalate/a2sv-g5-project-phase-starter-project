import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/repository/users_repository.dart';

class RegisterUser extends UseCase<String,UserModel>{
  final UsersRepository usersRepository;
  RegisterUser(this.usersRepository);

  @override
  Future<Either<Failure, String>> call(UserModel user) async {
    return await usersRepository.registerUser(user);
  }

  
}