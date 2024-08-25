import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/user_entity.dart';

abstract class AuthRepository {
  Future<Either<Failure, bool>> logIn(UserEntity user);
  Future<Either<Failure, bool>> logOut();
  Future<Either<Failure, bool>> signUp(UserEntity user);
}
