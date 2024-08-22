import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/user.dart';
import '../repository/auth_repository.dart';

class SignUp {
  final AuthRepository authrepository;

  SignUp(this.authrepository);
  Future<Either<Failure,User>> execute({required String name, required String email, required String password}){
    return  authrepository.signUp(name: name,email: email,password: password);
  }
}