import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/sign_up_entity.dart';
import '../repositories/auth_repository.dart';

class SignUpUsecase {
  final AuthRepository repository;

  SignUpUsecase({required this.repository});
  Future<Either<Failure, bool>> call({
    required String name,
    required String email,
    required String password,
    required String repeatedPassword,
  }) async {
    return await repository.signUp(
      SignUpEntity(
          name: name,
          email: email,
          password: password,
          repeatedPassword: repeatedPassword),
    );
  }
}
