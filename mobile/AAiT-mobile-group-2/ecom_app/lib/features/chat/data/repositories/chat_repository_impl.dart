import 'dart:io';

import 'package:dartz/dartz.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../../../core/platform/network_info.dart';
import '../../../auth/data/datasources/auth_local_data_source.dart';
import '../../domain/entities/message_entity.dart';
import '../../domain/entities/user_chat_entity.dart';
import '../../domain/repositories/chat_repository.dart';
import '../datasources/chat_remote_data_source.dart';
import '../models/message_model.dart';
import '../models/user_chat_model.dart';

class ChatRepositoryImpl extends ChatRepository{
  final ChatRemoteDataSource remoteDataSource;
  final AuthLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  ChatRepositoryImpl({required this.remoteDataSource, required this.localDataSource, required this.networkInfo});


  @override
  Future<Either<Failure, List<UserChatEntity>>> getChats() async{
    if (await networkInfo.isConnected){
      try{
        final email = await localDataSource.getEmail();
        final result = await remoteDataSource.getChats(email);

        return Right(UserChatModel.toEntityList(result));
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
  Future<Either<Failure, List<MessageEntity>>> getMessages(String ChatID) async{
    if (await networkInfo.isConnected){
      try{
        final result = await remoteDataSource.getMessages(ChatID);
        return Right(MessageModel.toEntityList(result));
      } on ServerException {
        return const Left(ServerFailure(ErrorMessages.serverError));
      } on SocketException {
        return const Left(ConnectionFailure(ErrorMessages.noInternet));
      } on CacheException {
        return const Left(CacheFailure(ErrorMessages.cacheError));
      }
    }  else {
      return const Left(ConnectionFailure(ErrorMessages.noInternet));
    }
  }

  @override
  Future<Either<Failure, UserChatEntity>> initiateChat(String UserID) async{
    if (await networkInfo.isConnected){
      try{
        final result = await remoteDataSource.initiateChat(UserID);
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
}