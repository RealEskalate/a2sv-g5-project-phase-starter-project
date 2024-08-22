import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/user.dart';

abstract class UserRepository{
  Future<Either<Failure, String>> loginUser(String email, String password);
  Future<Either<Failure, User>> registerUser(String email, String password, String name);
}