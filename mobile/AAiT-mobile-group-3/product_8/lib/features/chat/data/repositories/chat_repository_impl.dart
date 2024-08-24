import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';
import '../../domain/repositories/chat_repository.dart';
import '../data_source/chat_remote_data_source.dart';

class ChatRepositoryImpl extends ChatRepository {
  final ChatRemoteDataSource chatRemoteDataSource;

  ChatRepositoryImpl({required this.chatRemoteDataSource});

  @override
  Future<Either<Failure, List<ChatEntity>>> myChats() async {
    throw UnimplementedError();
  }

  // add the resto of the methods
  @override
  Future<Either<Failure, ChatEntity>> myChatById(String chatId) async {
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(
      String chatId) async {
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, ChatEntity>> initiateChat(String sellerId) async {
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, Unit>> deleteChat(String chatId) async {
    throw UnimplementedError();
  }
}
