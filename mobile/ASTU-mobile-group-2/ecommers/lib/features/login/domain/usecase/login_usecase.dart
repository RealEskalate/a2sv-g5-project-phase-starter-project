
import 'package:dartz/dartz.dart';

import '../../../../core/Error/failure.dart';
import '../entity/login_entity.dart';
import '../repositories/login_repositories.dart';

class LoginUseCase {
  final LoginRepositories repository;

  LoginUseCase({required this.repository});



  Future<Either<Failure,LoginEntity>> loginUser(String email,String password) async {

    final result = await repository.login(email,password);
  
    return result;
  }


  Future<Either<Failure,LoginEntity>> forgotPassword(String email) async {
    
    return await repository.forgotPassword(email);
  }

  Future<Either<Failure,bool>> registerUser(String email,String password,String fullName) async {
    return await repository.register(email,password,fullName);
  }


}