import 'package:dartz/dartz.dart'; // Import dartz for Either
import '../../../../core/errors/failure.dart';
import '../entity/chat_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure, List<ChatEntity>>> getAllChats();
  Future<Either<Failure, ChatEntity>> getChatById(String chatId);
  Future<Either<Failure, ChatEntity>> initiateChat(String userId);
  Future<Either<Failure, void>> deleteChat(String chatId);
}
