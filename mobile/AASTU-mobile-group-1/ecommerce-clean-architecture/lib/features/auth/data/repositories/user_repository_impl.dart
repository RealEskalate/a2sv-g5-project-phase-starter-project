import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/auth/data/data_sources/local_data_source.dart';
import 'package:ecommerce/features/auth/data/data_sources/remote_data_source.dart';
import 'package:ecommerce/features/auth/domain/entities/user.dart';

import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/networkinfo.dart';
import '../../domain/repository/user_repository.dart';
import '../model/user_model.dart';

class UserRepositoryImpl implements UserRepository {
  final UserRemoteDataSource remoteDataSource;
  final UserLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  UserRepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
    required this.networkInfo,
  });

  @override
  Future<Either<Failure, UserEntity>> registerUser(
      String name, String email, String password) async {
    try {
      if (await networkInfo.isConnected) {
        final remoteUser =
            await remoteDataSource.register(name, email, password);
        localDataSource.cacheUser(UserModel(
            id: remoteUser.id, name: name, email: email, password: password));
        return Right(remoteUser);
      } else {
        throw ServerException('failed to register user');
      }
    } on ServerException {
      return Left(ServerFailure('An error has occured'));
    } on SocketException {
      return Left(ServerFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, UserModel>> loginUser(
      String email, String password) async {
    try {
      if (await networkInfo.isConnected) {
        final remoteUser = await remoteDataSource.login(email, password);
        localDataSource.cacheUser(remoteUser);
        return Right(remoteUser);
      } else {
        throw ServerException('failed to login user');
      }
    } on ServerException {
      return Left(ServerFailure('An error has occured'));
    } on SocketException {
      return Left(ServerFailure('No internet connection'));
    }
  }
}
