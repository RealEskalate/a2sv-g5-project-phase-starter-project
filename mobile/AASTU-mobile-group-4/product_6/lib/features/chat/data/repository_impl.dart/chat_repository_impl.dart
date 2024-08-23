import 'package:dartz/dartz.dart';
import 'package:product_6/features/chat/data/datasources/chat_remote_datasource.dart';

import '../../../../core/connections/network_info.dart';
import '../../../../core/errors/exception.dart';
import '../../../../core/errors/failure.dart';
import '../../../auth/data/data_source/auth_local_datasource.dart';
import '../../domain/entity/chat_entity.dart';
import '../../domain/repository/chat_repository.dart';



class ChatRepositoryImpl implements ChatRepository {
  final ChatRemoteDataSource remoteDataSource;
  final AuthLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  ChatRepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
    required this.networkInfo,
  });

  @override
  Future<Either<Failure, List<ChatEntity>>> getAllChats() async {
    final connected = await networkInfo.isConnected;
    if (connected == true) {
      try {
        final token = await localDataSource.getAccessToken();
        if (token == null) {
          return Left(AuthenticationFailure("user not authenticated"));
        }
        final chats = await remoteDataSource.getAllChats(token);
        return Right(chats);
      } on ServerException {
        return Left(ServerFailure('Server Error'));
      }
    } else {
      return Left(NetworkFailure('no internet'));
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> getChatById(String chatId) async {
        final connected = await networkInfo.isConnected;

    if (connected == true) {
      try {
        final token = await localDataSource.getAccessToken();
        if (token == null) {
          return Left(AuthenticationFailure('user not authenticated'));
        }
        final chat = await remoteDataSource.getChatById(chatId, token);
        return Right(chat);
      } on ServerException {
        return Left(ServerFailure('Server Error'));
      }
    } else {
      return Left(NetworkFailure('no internet'));
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> initiateChat(String userId) async {
    final connected = await networkInfo.isConnected;

    if (connected == true) {
      try {
        final token = await localDataSource.getAccessToken();
        if (token == null) {
          return Left(AuthenticationFailure('user not authenticated'));
        }
        final chat = await remoteDataSource.initiateChat(userId, token);
        return Right(chat);
      } on ServerException {
        return Left(ServerFailure('Server Error'));
      }
    } else {
      return Left(NetworkFailure('no internet'));
    }
  }

  @override
  Future<Either<Failure, void>> deleteChat(String chatId) async {
        final connected = await networkInfo.isConnected;

    if (connected == true) {
      try {
        final token = await localDataSource.getAccessToken();
        if (token == null) {
          return Left(AuthenticationFailure('user not authenticated'));
        }
        await remoteDataSource.deleteChat(chatId, token);
        return Right(null);
      } on ServerException {
        return Left(ServerFailure('Server Error'));
      }
    } else {
      return Left(NetworkFailure('no internet'));
    }
  }
}
