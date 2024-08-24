import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/repositories/chat_repository.dart';
import '../datasources/chat_remote_data_source.dart';

class ChatRepositoryImpl implements ChatRepository {
  final ChatRemoteDataSource remoteDataSource;

  ChatRepositoryImpl(this.remoteDataSource);

  @override
  Future<Either<Failure, List<ChatEntity>>> myChats() async {
    try {
      final chats = await remoteDataSource.myChats();
      return Right(chats);
    } catch (e) {
      return Left(ServerFailure(message: 'message'));
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> myChatById() async {
    // try {
    //   final chat = await remoteDataSource.myChatById(chatId);
    //   return Right(chat);
    // } catch (e) {
    //   return Left(ServerFailure(message: 'message'));
    // }
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, void>> deleteChat(String chatId) async {
    try {
      await remoteDataSource.deleteChat(chatId);
      return const Right(null);
    } catch (e) {
      return Left(ServerFailure(message: 'message'));
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> initiateChat(String sellerId) async {
    try {
      final chat = await remoteDataSource.initiateChat(sellerId);
      return Right(chat);
    } catch (e) {
      return Left(ServerFailure(message: 'message'));
    }
  }
}
