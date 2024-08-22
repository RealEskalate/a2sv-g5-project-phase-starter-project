import 'package:dartz/dartz.dart';

import '../../../../core/error/exceptions.dart';
import '../../../../core/failure/failure.dart';
import '../../domain/entities/user.dart';
import '../../domain/repositories/user_repository.dart';
import '../data_sources/user_local_data_source.dart';
import '../data_sources/user_remote_data_source.dart';


class UserRepositoryImpl implements UserRepository {
  final UserRemoteDataSource remoteDataSource;
  final UserLocalDataSource localDataSource;

  UserRepositoryImpl({required this.remoteDataSource, required this.localDataSource});

  @override
  Future<Either<Failure, String>> loginUser(String email, String password) async {
    try {
      final access_token = await remoteDataSource.loginUser(email, password);
      await localDataSource.saveAccessToken(access_token);
      return Right(access_token);
    } on ServerException {
      return Left(ServerFailure(message: 'Failed to login'));
    }
  }

  @override
  Future<Either<Failure, User>> registerUser(String email, String password, String name) async {
    try {
      final userModel = await remoteDataSource.registerUser(email, password, name);
      return Right(userModel);
    } on ServerException {
      return Left(ServerFailure(message: 'Failed to register'));
    }
  }
}
