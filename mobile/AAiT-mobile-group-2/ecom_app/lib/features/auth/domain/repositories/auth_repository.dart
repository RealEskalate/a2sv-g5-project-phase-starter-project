import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/login_entity.dart';
import '../entities/register_entity.dart';
import '../entities/user_data_entity.dart';

abstract class AuthRepository {
  Future<Either<Failure, Unit>> login(LoginEntity loginEntity);
  Future<Either<Failure, Unit>> register(RegistrationEntity registrationEntity);
  Future<Either<Failure, UserDataEntity>> getUser();
  Future<Either<Failure, Unit>> logout();
}
