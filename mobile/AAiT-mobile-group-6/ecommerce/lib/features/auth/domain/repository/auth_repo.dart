import 'package:dartz/dartz.dart';

import '../../../../core/error/failures.dart';
import '../../data/models/auth_model.dart';
import '../entities/log_in_entity.dart';
import '../entities/sign_up_entity.dart';


abstract interface class AuthRepository {
  Future<Either<Failure, String>> signUP(SignUpEntity signup);
  Future<Either<Failure, String>> login(LogInEntity logIn);
  Future<Either<Failure, SignUPModel>> getUser();
  Future<Either<Failure,String>> logOut();
  
}
