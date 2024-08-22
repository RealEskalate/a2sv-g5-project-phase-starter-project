import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/auth/data/model/user_model.dart';
import 'package:ecommerce/features/auth/domain/entities/user.dart';

import '../../../../core/error/failure.dart';


abstract class UserRepository {
  Future<Either<Failure, UserEntity>> registerUser(String name, String email , String password);
  Future<Either<Failure, UserModel>> loginUser(String email, String password);
}