import 'package:dartz/dartz.dart';


import '../../../../core/errors/failure.dart';

import '../../../../core/usecases/auth/usecases.dart';
import '../entity/auth_entity.dart';

import '../repository/auth_repository.dart';

class Login implements UseCase<AuthResponseEntity, AuthEntity> {
  final AuthRepository repository;

  Login(this.repository);

  @override
  Future<Either<Failure, AuthResponseEntity>> call(AuthEntity authEntity) async {
    return await repository.login(authEntity);
  }
}
