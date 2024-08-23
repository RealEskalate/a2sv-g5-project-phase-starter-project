

import 'package:dartz/dartz.dart';

import '../../../../core/Error/failure.dart';
import '../../domain/entity/login_entity.dart';
import '../../domain/repositories/login_repositories.dart';
import '../datasource/remote_datasource.dart';


class LoginRepoImpl extends LoginRepositories {

  final RemoteDatasource remoteDatasourceImpl;

  LoginRepoImpl({required this.remoteDatasourceImpl,});
  @override

  
   Future<Either<Failure,LoginEntity>>  forgotPassword(String email) {
    throw UnimplementedError();
  }

  @override
   Future<Either<Failure,LoginEntity>> login(String email, String password) async{
 
    final result = await remoteDatasourceImpl.login(email, password);
    
    return result;
    
  }
  
  @override
  Future<Either<Failure, bool>> register(String email, String password, String fullName) async{
    final result = await remoteDatasourceImpl.register(email, password, fullName);
    return result;
  }
  
 
}