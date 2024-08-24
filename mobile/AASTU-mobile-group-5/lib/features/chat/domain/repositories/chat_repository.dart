import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../data/models/chat_model.dart';
import '../entities/chat_entity.dart';

abstract class ChatRepository {
  Future<Either<List, ChatModel>> myChats();
  Future<Either<Failure, ChatEntity>> myChatById();
  Future<Either<Failure, void>> deleteChat(String chatId);
  Future<Either<Failure, ChatModel>> initiateChat(String sellerId);
}
