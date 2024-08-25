import 'package:dartz/dartz.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/errors/exceptions/product_exceptions.dart';
import '../../../../core/errors/failures/failure.dart';
import '../../domain/entities/user_entity.dart';
import '../../domain/repositories/auth_repository.dart';
import '../data_source/auth_local_data_source.dart';
import '../data_source/remote_auth_data_source.dart';

class AuthRepositoryImpl extends AuthRepository {
  final RemoteAuthDataSource remoteAuthDataSource;
  final AuthLocalDataSource authLocalDataSource;

  AuthRepositoryImpl(
      {required this.remoteAuthDataSource, required this.authLocalDataSource});

  @override
  Future<Either<Failure, bool>> logIn(UserEntity user) async {
    try {
      final result = await remoteAuthDataSource.logIn(user);
      await authLocalDataSource.saveToken(result);
      return const Right(true);
    } on CacheException {
      return Left(CacheFailure(AppData.getMessage(AppData.cacheError)));
    } on LoginException {
      return Left(LoginFailure(AppData.getMessage(AppData.loginFailed)));
    } on ServerException {
      return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
    } on Exception {
      return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
    }
  }

  @override
  Future<Either<Failure, bool>> signUp(UserEntity user) async {
    try {
      await remoteAuthDataSource.signUp(user);
      return const Right(true);
    } on UserConflictException {
      return Left(UserExistsFailure(AppData.getMessage(AppData.userExists)));
    } on ServerException {
      return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
    } on Exception {
      return Left(ServerFailure(AppData.getMessage(AppData.serverError)));
    }
  }

  @override
  Future<Either<Failure, bool>> logOut() async {
    try {
      await authLocalDataSource.clearToken();
      return const Right(true);
    } on CacheException {
      return Left(CacheFailure(AppData.getMessage(AppData.logoutError)));
    }
  }
}
