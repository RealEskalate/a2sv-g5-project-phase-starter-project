import 'package:dartz/dartz.dart';

import '../../../../../../../core/failure/failure.dart';
import '../entities/chat_entity.dart';
import '../entities/message_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure, List<ChatEntity>>> myChats();
  Future<Either<Failure, ChatEntity>> myChatById(String chatId);
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(String chatId);
  Future<Either<Failure, ChatEntity>> initiateChat(String sellerId);
  Future<Either<Failure, Unit>> deleteChat(String chatId);
}
