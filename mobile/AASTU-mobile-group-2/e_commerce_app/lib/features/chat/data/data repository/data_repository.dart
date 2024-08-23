import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';

class DataRepositoryImp implements ChatRepository {
  @override
  Future<Either<Failure, ChatEntity>> createChatById(String sellerId) {
    
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, bool>> deleteChatById(String chatId) {
    // TODO: implement deleteChatById
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<ChatEntity>>> getAllChats(String userId) {
    // TODO: implement getAllChats
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, ChatEntity>> getChatById(String chatId) {
    
    throw UnimplementedError();
  }
  
}