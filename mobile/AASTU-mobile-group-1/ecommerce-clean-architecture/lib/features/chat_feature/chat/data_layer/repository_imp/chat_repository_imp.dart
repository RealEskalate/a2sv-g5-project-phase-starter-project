import '../../../../../core/error/failure.dart';
import '../../domain/entity/message.dart';
import 'package:dartz/dartz.dart';

import '../../domain/repository/chat_repository.dart';

class ChatRepositoryImp extends ChatRepository{
  @override
  Future<Either<Failure, String>> chatRoom(String receiverId) {
    // TODO: implement chatRoom
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, void>> deleteMessage(String chatId) {
    // TODO: implement deleteMessage
    throw UnimplementedError();
  }

  @override
  Stream<Either<Failure, List<Message>>> getMessages(String chatId) {
    // TODO: implement getMessages
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, void>> sendMessage(Message message) {
    // TODO: implement sendMessage
    throw UnimplementedError();
  }
}