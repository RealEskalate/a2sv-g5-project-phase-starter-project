import 'dart:async';

import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../entities/sign_in_entity.dart';
import '../entities/sign_up_entity.dart';
import '../entities/signed_in_entity.dart';
import '../entities/user_entity.dart';

abstract class AuthRepository {
  Future<Either<Failure, SignedInEntity>> signIn(SignInEntity signIn);
  Future<Either<Failure, bool>> signUp(SignUpEntity signUp);

  Future<Either<Failure, bool>> logOut();
  Future<Either<Failure, bool>> checkSignedIn();
  Future<Either<Failure, UserEntity>> getUser();
}
