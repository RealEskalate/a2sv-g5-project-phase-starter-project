import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/sign_in_entity.dart';
import '../entities/signed_in_entity.dart';
import '../repositories/auth_repository.dart';

class SignInUsecase {
  final AuthRepository repository;

  SignInUsecase({required this.repository});

  Future<Either<Failure, SignedInEntity>> call(
      {required email, required password}) async {
    return await repository
        .signIn(SignInEntity(email: email, password: password));
  }
}
