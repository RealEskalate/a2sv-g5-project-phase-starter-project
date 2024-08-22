import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:flutter/widgets.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/platform/network_info.dart';
import '../../domain/entities/login_entity.dart';
import '../../domain/entities/register_entity.dart';
import '../../domain/entities/user_data_entity.dart';
import '../../domain/repositories/auth_repository.dart';
import '../datasources/auth_local_data_source.dart';
import '../datasources/auth_remote_data_source.dart';
import '../models/login_model.dart';
import '../models/register_model.dart';

class AuthRepositoryImpl extends AuthRepository {
  final AuthRemoteDataSource remoteDataSource;
  final AuthLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  AuthRepositoryImpl(
      {required this.remoteDataSource,
      required this.localDataSource,
      required this.networkInfo});
  @override
  Future<Either<Failure, UserDataEntity>> getUser() async {
    if (await networkInfo.isConnected) {
      try {
        final token = await localDataSource.getToken();
        final result = await remoteDataSource.getUser(token);
        return Right(result.toEntity());
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(ConnectionFailure(ErrorMessages.noInternet));
      } on CacheException {
        return const Left(CacheFailure(ErrorMessages.cacheError));
      }
    } else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }

  @override
  Future<Either<Failure, Unit>> login(LoginEntity loginEntity) async {
    if (await networkInfo.isConnected) {
      try {
        final result =
            await remoteDataSource.login(LoginModel.toModel(loginEntity));
        try {
          await localDataSource.cacheToken(result.token);
        } on CacheException {
          debugPrint('Caching Token Error');
        }
        return const Right(unit);
      } on UnauthorizedException {
        return Left(UnauthorizedFailure(ErrorMessages.forbiddenError));
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(ConnectionFailure(ErrorMessages.noInternet));
      }
    } else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }

  @override
  Future<Either<Failure, Unit>> register(
      RegistrationEntity registrationEntity) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteDataSource
            .register(RegisterModel.toModel(registrationEntity));
        return Right(result);
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(ConnectionFailure(ErrorMessages.noInternet));
      } on UserAlreadyExistsException {
        return Left(UserAlreadyExistsFailure(ErrorMessages.userAlreadyExists));
      }
    } else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }

  @override
  Future<Either<Failure, Unit>> logout() async {
    try {
      await localDataSource.logout();
      return const Right(unit);
    } on CacheException {
      return const Left(CacheFailure(ErrorMessages.cacheError));
    }
  }
}
