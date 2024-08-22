import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/user.dart';
import '../repository/authentication_repository.dart';

class LoginUseCase {
  final AuthenticationRepository repository;

  LoginUseCase(this.repository);

  Future<Either<Failure, User>> call(String email, String password) async {
    try {
      final result = await repository.logIn(email: email, password: password);
      return Right(result);
    } catch (e) {
      return const Left(ServerFailure('Server Failure'));
    }
  }
}
