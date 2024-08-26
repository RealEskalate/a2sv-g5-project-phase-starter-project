import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/chat_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure, List<ChatEntity>>> myChats();
  Future<Either<Failure, ChatEntity>> myChatById(String id);
  Future<Either<Failure, void>> deleteChat(String chatId);
  Future<Either<Failure, ChatEntity>> initiateChat(String sellerId);
}
