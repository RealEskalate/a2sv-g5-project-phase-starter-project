import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import '../../../../core/errors/failure/failures.dart';


abstract class ChatRepository {
  Future<Either<Failure, List<ChatEntity>>> getAllChats();
  Future<Either<Failure, ChatEntity>> getMyChatById(String chatId);
  Future<Either<Failure, String>> initiateChat(String userId);
  Future<Either<Failure, List<ChatModel>>> getChatMessages(String chatId);
  Future<Either<Failure, String>> deleteChat(String chatId);
  
} 