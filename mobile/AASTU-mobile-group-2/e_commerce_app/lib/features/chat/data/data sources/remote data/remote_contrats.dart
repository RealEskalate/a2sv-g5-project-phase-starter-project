import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/data/models/chat_model.dart';

abstract class RemoteContrats {
  Future<Either<Failure,ChatModel>> getChatById(String chatId);
  Future<Either<Failure,List<ChatModel>>> getAllChats(String userId);
  Future<Either<Failure,ChatModel>> createChatById(String sellerId);
  Future<Either<Failure,bool>> deleteChatById(String chatId);
}