import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:flutter/material.dart';

import '../../../../core/exception/exception.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/platform/network_info.dart';
import '../../domain/entities/sign_in_user_entitiy.dart';
import '../../domain/entities/sign_up_user_entitiy.dart';
import '../../domain/entities/user_data_entity.dart';
import '../../domain/repositories/auth_repository.dart';
import '../data_source/auth_local_data_source.dart';
import '../data_source/auth_remote_data_source.dart';
import '../models/sign_in_user_model.dart';
import '../models/sign_up_user_model.dart';

class AuthRepositoryImpl extends AuthRepository {
  final AuthRemoteDataSource authRemoteDataSource;
  final AuthLocalDataSource authLocalDataSource;
  final NetworkInfo networkInfo;
  AuthRepositoryImpl(
      {required this.authRemoteDataSource,
      required this.authLocalDataSource,
      required this.networkInfo});

  @override
  Future<Either<Failure, Unit>> signUp(
      SignUpUserEntitiy signUpUserEntitiy) async {
   if (await networkInfo.isConnected) {
      try {
        final result = await authRemoteDataSource
            .signUp(SignUpUserModel.toModel(signUpUserEntitiy));
        return Right(result);
      } on ServerException {
        return const Left(ServerFailure(message: 'Server Error'));
      } on SocketException {
        return const Left(ConnectionFailure(message: 'No Internet Connection'));
      } on UserAlreadyExistsException {
        return const Left(PermissionFailure(message: 'User Already Exists'));
      }
    } else {
      return const Left(ConnectionFailure(message: 'No Internet Connection'));
    }
  }

  @override
  Future<Either<Failure, Unit>> signIn(
      SignInUserEntitiy signInUserEntitiy) async {
    if (await networkInfo.isConnected) {
      try {
        final result =
            await authRemoteDataSource.signIn(SignInUserModel.toModel(signInUserEntitiy));
        try {
          await authLocalDataSource.cacheToken(result.token);
        } on CacheException {
          debugPrint('Caching Token Error');
        }
        return const Right(unit);
      } on UnauthorizedException {
        return const Left(PermissionFailure(message: 'Unauthorized'));
      } on ServerException {
        return const Left(ServerFailure(message: 'Server Error'));
      } on SocketException {
        return const Left(ConnectionFailure(message: 'No Internet Connection'));
      }
    } else {
      return const Left(ConnectionFailure(message: 'No Internet Connection'));
    }
  }

  @override
  Future<Either<Failure, UserDataEntity>> getUser() async {
    if (await networkInfo.isConnected) {
      try {
        final token = await authLocalDataSource.getToken();
        final result = await authRemoteDataSource.getUser(token);
        return Right(result.toEntity());
      } on ServerException {
        return const Left(ServerFailure(message: 'Server Error'));
      } on SocketException {
        return const Left(ConnectionFailure(message: 'No Internet Connection'));
      }
    } else {
      return const Left(ConnectionFailure(message: 'No Internet Connection'));
    }
  }

  @override
  Future<Either<Failure, Unit>> logOut() async {
    try {
      await authLocalDataSource.deleteToken();
      return const Right(unit);
    } on CacheException {
      return const Left(CacheFailure(message: 'Cache Error'));
    }
  }
}
