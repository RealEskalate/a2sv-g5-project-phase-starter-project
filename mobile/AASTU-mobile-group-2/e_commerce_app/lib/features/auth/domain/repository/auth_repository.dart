import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/user.dart';

abstract class AuthRepository {
  Future<Either<Failure, String>> logIn(
      {required String email, required String password});

  Future<Either<Failure, User>> signUp(
      {required String name,required String email, required String password});

  Future<Either<Failure, void>> signOut();
  Future<Either<Failure, User>> getCurrentUser();
}