

import 'package:dartz/dartz.dart';

import '../../../../core/Error/failure.dart';
import '../entity/login_entity.dart';

abstract class LoginRepositories {

  Future<Either<Failure,LoginEntity>> login(String email, String password);


  Future<Either<Failure,LoginEntity>>  forgotPassword(String email);

  Future<Either<Failure,bool>>  register(String email, String password, String fullName);



}
