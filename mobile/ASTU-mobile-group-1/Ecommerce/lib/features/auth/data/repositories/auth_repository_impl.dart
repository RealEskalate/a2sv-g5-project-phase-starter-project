import 'dart:async';
import 'dart:developer';

import 'package:dartz/dartz.dart';

import '../../../../core/cubit/user_cubit.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/network/network_info.dart';
import '../../../../injection_container.dart';
import '../../domain/entities/sign_in_entity.dart';
import '../../domain/entities/sign_up_entity.dart';
import '../../domain/entities/signed_in_entity.dart';
import '../../domain/entities/user_entity.dart';
import '../../domain/repositories/auth_repository.dart';
import '../data_sources/auth_local_data_source.dart';
import '../data_sources/auth_remote_data_source.dart';
import '../models/sign_in_model.dart';
import '../models/sign_up_model.dart';

class AuthRepositoryImp implements AuthRepository {
  final AuthRemoteDataSource authRemoteDataSource;
  final AuthLocalDataSource authLocalDataSource;
  final NetworkInfo networkInfo;
  AuthRepositoryImp(
      {required this.networkInfo,
      required this.authRemoteDataSource,
      required this.authLocalDataSource});

  @override
  Future<Either<Failure, UserEntity>> getUser() async {
    final isConnected = await networkInfo.isConnected;

    if (isConnected) {
      try {
        final userModel = await authRemoteDataSource.getUser();
        final toke = authLocalDataSource.getToken();

        final userEntity = UserEntity.fromModel(userModel);
        sl<UserCubit>().updateUser(userEntity);

        return Right(userEntity);
      } catch (e) {
        log(e.toString());
        return const Left(UnkownFailure());
      }
    } else {
      log('outside');
      return const Left(ConnectionFailure());
    }
  }

  @override
  Future<Either<Failure, bool>> checkSignedIn() async {
    try {
      final result = await authLocalDataSource.checkSignedIn();
      return Right(result);
    } catch (_) {
      return const Left(CacheFailure('Error while checking status'));
    }
  }

  @override
  Future<Either<Failure, SignedInEntity>> signIn(SignInEntity signIn) async {
    if (await networkInfo.isConnected) {
      try {
        final signInModel = SignInModel(
          email: signIn.email,
          password: signIn.password,
        );
        final signedIn = await authRemoteDataSource.signIn(signInModel);
        return Right(signedIn);
      } on ServerException catch (e) {
        return Left(
          ServerFailure(e.message),
        );
      } catch (_) {
        return const Left(UnkownFailure());
      }
    } else {
      return const Left(ConnectionFailure());
    }
  }

  @override
  Future<Either<Failure, bool>> signUp(SignUpEntity signUp) async {
    if (await networkInfo.isConnected) {
      try {
        final signInModel = SignUpModel(
            email: signUp.email,
            password: signUp.password,
            name: signUp.name,
            repeatedPassword: signUp.repeatedPassword);
        final response = await authRemoteDataSource.signUp(signInModel);
        return Right(response);
      } on ServerException catch (e) {
        return Left(ServerFailure(e.message));
      } catch (_) {
        return const Left(ServerFailure(
            'An unknown error occurred. Please try again later.'));
      }
    } else {
      return const Left(ServerFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, bool>> logOut() async {
    try {
      final response = await authLocalDataSource.logOut();
      return Right(response);
    } catch (e) {
      return const Left(ServerFailure('Something thinge went wrong'));
    }
  }
}
