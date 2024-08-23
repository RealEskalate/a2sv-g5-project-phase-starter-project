import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';

import '../model/chat_model.dart';
import '../model/message_model.dart';

abstract class RemoteAbstract {
  Future<Either<Failure,void>>sendMessage(MessageModel message);
  Stream<Either<Failure,List<MessageModel>>> getMessages(String chatId);
  Future<Either<Failure,void>> deleteMessage(String chatId);
  Future<Either<Failure,String>> chatRoom(String token,String receiverId);
  Stream<Either<Failure,List<ChatModel>>> getChatHistory(String token);

}