import 'package:dartz/dartz.dart';

import '../../../../core/error/exceptions.dart';
import '../../../../core/error/failures.dart';
import '../../../../core/platform/network_info.dart';
import '../../domain/entities/sign_up_entity.dart';
import '../../domain/repository/auth_repo.dart';
import '../datasource/auth_local_datasource/login_local_datasource.dart';
import '../datasource/remote_datasource/auth_remote_datasource.dart';
import '../models/auth_model.dart';

class AuthRepoImpl implements AuthRepository {
  final AuthRemoteDataSource authRemoteDataSource;
  final NetworkInfo networkInfo;
  final UserLogInLocalDataSource localDataSource;
  AuthRepoImpl(
      this.authRemoteDataSource, this.networkInfo, this.localDataSource);

  @override
  Future<Either<Failure, String>> signUP(SignUpEntity signup) async {
    if (await networkInfo.isConnected) {
      try {
        final user = SignUPModel(
            id: signup.id,
            name: signup.name,
            email: signup.email,
            password: signup.password);

        final res = await authRemoteDataSource.signUp(user);

        return Right(res);
      } on ServerException {
        return Left(ServerFailure('Server Failure'));
      }
    } else {
      return Left(NetworkFailure('You are not connected to the internet'));
    }
  }

  @override
  Future<Either<Failure, String>> login(logIn) async {
    if (await networkInfo.isConnected) {
      try {
        final user = LogInModel(
            id: logIn.id, email: logIn.email, password: logIn.password);
        final token = await authRemoteDataSource.logIn(user);
        await localDataSource.cacheToken(token);
        return Right(token);
      } on ServerException {
        return Left(ServerFailure('Server Failure'));
      }
    } else {
      return Left(NetworkFailure('You are not connected to the internet'));
    }
  }

  @override
  Future<Either<Failure, SignUPModel>> getUser() async {
    if (await networkInfo.isConnected) {
      try {
        final userDataModel = await localDataSource.getUser();
        return Right(userDataModel!);
      } on ServerException {
        return Left(ServerFailure('Server failure'));
      }
    } else {
      return Left(NetworkFailure('you are not connected to the internet'));
    }
  }

  @override
  Future<Either<Failure, String>> logOut() async {
    if (await networkInfo.isConnected) {
      try {
        await localDataSource.deleteUser();

        return const Right('Successfully logged out');
      } on ServerException {
        return Left(ServerFailure('Server failure'));
      } on CacheFailure {
        return Left(CacheFailure('Cache failure'));
      }
    } else {
      return Left(NetworkFailure('You are not connected to the internet.'));
    }
  }
}
