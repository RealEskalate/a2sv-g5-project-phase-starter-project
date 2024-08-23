import 'package:dartz/dartz.dart';

import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../domain/entities/log_in.dart';
import '../../domain/entities/sign_up.dart';
import '../../domain/entities/user_data.dart';
import '../../domain/repositories/auth_repo.dart';
import '../data_sources/remote/auth_remote_data_source.dart';
import '../model/mapper.dart';

class AuthRepositoryImpl implements AuthRepository {
  final AuthRemoteDataSource authRemoteDataSource;

  AuthRepositoryImpl({required this.authRemoteDataSource});
  @override
  Future<Either<Failure, UserEntity>> getCurrentUser() async {
    try {
      final user = await authRemoteDataSource.getCurrentUser();
      return Right(user.toUserEntity());
    } on Exception {
      return const Left(ServerFailure('An error has occurred'));
    }
  }

  @override
  Future<Either<Failure, void>> logIn(LogInEntity logInEntity) async {
    try {
      await authRemoteDataSource.logIn(logInEntity.toProductModel());

      // ignore: void_checks
      return const Right(unit);
    } catch (e) {
      return const Left(ServerFailure('cannot login'));
    }
  }

  @override
  Future<Either<Failure, void>> logOut() async {
    try {
      await authRemoteDataSource.logOut();
      // ignore: void_checks
      return const Right(unit);
    } catch (e) {
      return const Left(ServerFailure('can not logout'));
    }
  }

  @override
  Future<Either<Failure, void>> signUp(SignUpEntity signUpEntity) async {
    try {
      await authRemoteDataSource.signUp(signUpEntity.toSignUpModel());
      // ignore: void_checks
      return const Right(unit);
    } on ServerException catch (e) {
      return Left(ServerFailure(e.message ?? 'Unknown error occurred'));
    }
  }
}
