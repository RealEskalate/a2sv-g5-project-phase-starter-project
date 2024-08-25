import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

import '../model/chat_model.dart';
import '../model/message_model.dart';

abstract class RemoteAbstract {
   void initializeSocket(String token) {}

  void sendMessage(String chatId, String content, String type);
  Stream<MessageModel> getMessages(String chatId,String token);
  Future<Either<Failure,void>> deleteMessage(String chatId);
  Future<Either<Failure,String>> chatRoom(String token,String receiverId);
  Stream<Either<Failure,List<ChatModel>>> getChatHistory(String token);

}

