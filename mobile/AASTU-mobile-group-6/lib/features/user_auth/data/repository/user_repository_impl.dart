import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_access.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/repository/users_repository.dart';

class UserRepositoryImpl extends UsersRepository{
  final UserRemoteDataSourceImpl userRemoteDataSource;
  UserRepositoryImpl({required this.userRemoteDataSource});



  @override
  Future<Either<Failure,String>> registerUser(UserModel user) async{
    try{
      final result = await userRemoteDataSource.registerUser(user);
      return Right(result);
    }on ServerFailure{
      return Left(ServerFailure('Server Failure'));
    }
  }
  @override
  Future<Either<Failure,String>> loginUser(UserEntity user) async{
      final result = await userRemoteDataSource.loginUser(user);
      return result.fold((l)=> Left(l), (r)=> Right(r));
    
  }
}