import '../../../../../core/error/failure.dart';
import '../model/message_model.dart';
import 'package:dartz/dartz.dart';

import 'remote_abstract.dart';

class RemoteChatSource extends RemoteAbstract{
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
  Stream<Either<Failure, List<MessageModel>>> getMessages(String chatId) {
    // TODO: implement getMessages
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, void>> sendMessage(MessageModel message) {
    // TODO: implement sendMessage
    throw UnimplementedError();
  }

}