import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/log_in.dart';
import '../entities/sign_up.dart';
import '../entities/user_data.dart';

abstract class AuthRepository {
  Future<Either<Failure, void>> signUp(SignUpEntity signUpEntity);
  Future<Either<Failure, void>> logIn(LogInEntity logInEntity);
  Future<Either<Failure, void>> logOut();
  Future<Either<Failure, UserEntity>> getCurrentUser();
}