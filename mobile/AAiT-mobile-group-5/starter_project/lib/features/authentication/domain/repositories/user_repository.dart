import 'package:dartz/dartz.dart';
import 'package:starter_project/core/error/failure.dart';

abstract class UserRepository {
  Future<Either<Failure, String>> userLogin();
  Future<Either<Failure, List<String>>> userSignin();
  Future<Either<Failure, List<String>>> getUser();
}
