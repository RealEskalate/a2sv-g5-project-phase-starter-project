import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/chat_entity.dart';
import '../entities/message_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure, ChatEntity>> initiateChat(String userId);
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(String chatId);
  Future<Either<Failure, ChatEntity>> myChatbyId(String chatId);
  Future<Either<Failure, List<ChatEntity>>> myChat();
  Future<Either<Failure, bool>> deleteChat(String chatId);

 Future<Either<Failure,void>> sendMessage(String chatId, String message, String type) ;
}
