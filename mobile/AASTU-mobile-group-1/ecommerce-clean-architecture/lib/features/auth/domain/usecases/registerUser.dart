import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/auth/domain/entities/user.dart';
import 'package:ecommerce/features/auth/domain/repository/user_repository.dart';

import '../../../../core/error/failure.dart';


class RegisterUserUseCase {
  final UserRepository userRepository;

  RegisterUserUseCase(this.userRepository);

  Future<Either<Failure, UserEntity>> register(String name, String email, String password) {
    return userRepository.registerUser(name, email, password);
  }
}