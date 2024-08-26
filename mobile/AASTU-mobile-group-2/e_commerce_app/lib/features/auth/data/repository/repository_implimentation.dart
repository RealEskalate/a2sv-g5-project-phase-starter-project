import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/auth/data/data_sources/local_data_sources.dart';
import 'package:e_commerce_app/features/auth/data/data_sources/remote_data_sources.dart';
import 'package:e_commerce_app/features/auth/domain/entities/user.dart';
import 'package:e_commerce_app/features/auth/domain/repository/auth_repository.dart';
import 'package:e_commerce_app/features/product/data/data_sources/product_remote_data_source.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/network/network_info.dart';

class AuthRepositoryImplimentation extends AuthRepository {
  AuthRemoteDataSources authRemoteDataSources;
  AuthLocalDataSource authLocalDataSource;
  NetworkInfo networkInfo;
  AuthRepositoryImplimentation(
      {required this.authRemoteDataSources,
      required this.authLocalDataSource,
      required this.networkInfo});
  @override
  Future<Either<Failure, String>> logIn(
      {required String email, required String password}) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await authRemoteDataSources.logIn(email, password);
        authLocalDataSource.cacheToken(result);

        return Right(result);
      } catch (e) {
        print(e);
        return left(Failure("failed to login"));
      }
    } else {
      return left(Failure("network not available"));
    }
  }

  @override
  Future<Either<Failure, void>> signOut() {
    // TODO: implement signOut
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, User>> signUp(
      {required String name,
      required String email,
      required String password}) async {
    if (await networkInfo.isConnected) {
      try {
        final result =
            await authRemoteDataSources.signUp(name, email, password);
        return Right(result.toUser());
      } catch (e) {
        print(e);
        return left(Failure("failed to login"));
      }
    } else {
      return left(Failure("network not available"));
    }
  }

  @override
  Future<Either<Failure, User>> getCurrentUser()  async{
      if (await networkInfo.isConnected) {
      try {
            final result = await authRemoteDataSources.getUser();
        print(result);
        return Right(result.toUser());
      } catch (e) {
        print(e);
        return Left(Failure("failed to get user"));
      }
    } else {
      return Left(Failure("network not available"));
    }
  }
}
