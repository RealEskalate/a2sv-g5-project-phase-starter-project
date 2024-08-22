import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_access.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';

abstract class UsersRepository{
  Future<Either<Failure,String>> registerUser(UserModel user);
  Future<Either<Failure,String>> loginUser(UserEntity user);
  // Future<Either<Failure,String>> registerUser(UserEntity user);

}