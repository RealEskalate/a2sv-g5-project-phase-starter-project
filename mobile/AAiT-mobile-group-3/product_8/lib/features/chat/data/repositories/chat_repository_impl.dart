import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/platform/network_info.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';
import '../../domain/repositories/chat_repository.dart';
import '../data_source/chat_local_data_source.dart';
import '../data_source/chat_remote_data_source.dart';

class ChatRepositoryImpl extends ChatRepository {
  final ChatRemoteDataSource chatRemoteDataSource;
  final ChatLocalDataSource chatLocalDataSource;
  final NetworkInfo networkInfo;

  ChatRepositoryImpl(
      {required this.chatRemoteDataSource,
      required this.chatLocalDataSource,
      required this.networkInfo});

  @override
  Future<Either<Failure, List<ChatEntity>>> myChats() async {
    if (await networkInfo.isConnected) {
      try {
        final result = await chatRemoteDataSource.myChats();
        await chatLocalDataSource.cacheChats(result);
        return Right(result);
      } on Exception {
        return const Left(ServerFailure(message: 'Server Failure'));
      }
    } else {
      try {
        final result = await chatLocalDataSource.getChats();
        return Right(result);
      } on Exception {
        return const Left(CacheFailure(message: 'Cache Failure'));
      }
    }
  }

  // add the resto of the methods
  @override
  Future<Either<Failure, ChatEntity>> myChatById(String chatId) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await chatRemoteDataSource.myChatById(chatId);
        return Right(result);
      } on Exception {
        return const Left(ServerFailure(message: 'Server Failure'));
      }
    } else {
      return const Left(ConnectionFailure(message: 'No Internet Connection'));
    }
  }

  @override
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(
      String chatId) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await chatRemoteDataSource.getChatMessages(chatId);
        return Right(result);
      } on Exception {
        return const Left(ServerFailure(message: 'Server Failure'));
      }
    } else {
      return const Left(ConnectionFailure(message: 'No Internet Connection'));
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> initiateChat(String sellerId) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await chatRemoteDataSource.initiateChat(sellerId);
        return Right(result);
      } on Exception {
        return const Left(ServerFailure(message: 'Server Failure'));
      }
    } else {
      return const Left(ConnectionFailure(message: 'No Internet Connection'));
    }
  }

  @override
  Future<Either<Failure, Unit>> deleteChat(String chatId) async {
    if (await networkInfo.isConnected) {
      try {
        await chatRemoteDataSource.deleteChat(chatId);
        return const Right(unit);
      } on Exception {
        return const Left(ServerFailure(message: 'Server Failure'));
      }
    } else {
      return const Left(ConnectionFailure(message: 'No Internet Connection'));
    }
  }
}
