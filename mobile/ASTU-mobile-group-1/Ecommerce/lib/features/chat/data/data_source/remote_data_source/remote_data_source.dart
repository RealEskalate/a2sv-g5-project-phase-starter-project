import 'package:dartz/dartz.dart';

import '../../../../../core/error/failure.dart';
import '../../model/chat_model.dart';
import '../../model/message_model.dart';

abstract class RemoteDataSource {
  ///this will take chat recievers id and then initiate the chat and then return chat id
  ///if it already created just it will return chat id
  Future<ChatModel> initiateChat(String recieverId);

  ///this will delete the chat that user has with specific user
  Future<bool> deleteChat(String chatId);

  ///this method will return chat model that user has with specific user with out
  ///showng chat messages
  Future<ChatModel> getChatById(String chatId);

  ///will return all chat that users has
  Future<List<ChatModel>> getAllChats();

  ///used to send message from sender to server and the server will deliver to reciever
  Future<void> sendMessage(String chatId, String message, String type);

  ///this to always listen server to get new message, so it have to stream
  Future<List<MessageModel>> getChatMessages(String chatId);

  Stream<MessageModel> getMessages();
  Future<void> updateAccessToken(String token);
}
