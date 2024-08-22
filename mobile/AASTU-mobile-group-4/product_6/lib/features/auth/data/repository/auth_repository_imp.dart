import 'package:dartz/dartz.dart';

import '../../../../core/connections/network_info.dart';
import '../../../../core/errors/failure.dart';
import '../../domain/entity/auth_entity.dart';
import '../../domain/repository/auth_repository.dart';
import '../data_source/auth_local_datasource.dart';
import '../data_source/auth_remote_datasource.dart';
import '../models/auth_model.dart';

class AuthRepositoryImpl implements AuthRepository {
  final AuthRemoteDataSource remoteDataSource;
  final AuthLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  AuthRepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
    required this.networkInfo,
  });

  @override
  Future<Either<Failure, AuthResponseEntity>> login(
      AuthEntity authEntity) async {
    final isConnected = await networkInfo.isConnected;

    if (isConnected == true) {
      try {
        final remoteResponse = await remoteDataSource.login(authEntity);
        await localDataSource.cacheAccessToken(
            remoteResponse.accessToken); // Cache the access token
        return Right(remoteResponse);
      } catch (exception) {
        return Left(ServerFailure('Server Error'));
      }
    } else {
      return Left(ServerFailure('Network Error'));
    }
  }

  @override
  Future<Either<Failure, UserEntity>> getUserProfile() async {
    final isConnected = await networkInfo.isConnected;
    
    if (isConnected == true) {

      try {
        final token = await localDataSource.getAccessToken();

        final userProfile = await remoteDataSource.getUserProfile(token!);
        final user = UserEntity(id: userProfile.id, name: userProfile.name, email: userProfile.email); 
        return Right(user);
        
      } catch (exception) {
        return Left(ServerFailure('Server error'));
      }
    } else {
      return Left(ServerFailure('Network error'));
    }
  }

  @override
  Future<Either<Failure, void>> register(SentUserEntity userEntity) async {
    final isConnected = await networkInfo.isConnected;

    if (isConnected == true) {
      try {
        await remoteDataSource.register(userEntity);
        return Right(Future.value());
      } catch (exception) {
        return Left(ServerFailure('Server Error'));
      }
    } else {
      return Left(ServerFailure('Network error'));
    }
  }
  
  @override
  Future<Either<Failure, void>> logout() async{

      try {
        await localDataSource.clearAccessToken();
        return Right(null);
      } catch (exception) {
        return Left(ServerFailure('user not logged in'));
      }
   
  }
}
