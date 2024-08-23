import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/data/data%20sources/remote%20data/remote_contrats.dart';
import 'package:e_commerce_app/features/chat/data/models/chat_model.dart';

class RemoteDataSourceImp implements RemoteContrats {
  @override
  Future<Either<Failure, ChatModel>> createChatById(String sellerId) {
    // TODO: implement createChatById
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, bool>> deleteChatById(String chatId) {
    // TODO: implement deleteChatById
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<ChatModel>>> getAllChats(String userId) {
    // TODO: implement getAllChats
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, ChatModel>> getChatById(String chatId) {
    // TODO: implement getChatById
    throw UnimplementedError();
  }

}