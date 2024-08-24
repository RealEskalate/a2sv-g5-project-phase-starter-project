
import 'package:dartz/dartz.dart';

import '../../../../../core/error/failure.dart';
import '../../model/chat_model.dart';
import '../../model/message_model.dart';

abstract class RemoteDataSource {

    Future<Either<Failure,String>>initiateChat(String recieverId);
    Future<bool> deleteChat(String chatId);
    Future<ChatModel> getChatById(String receiverId);
    Future<List<ChatModel>> getAllChats();
    Future<void> sendMessage(String chatId, String message, String type);
    Stream<MessageModel> getChatMessages(String chatId);

}
 