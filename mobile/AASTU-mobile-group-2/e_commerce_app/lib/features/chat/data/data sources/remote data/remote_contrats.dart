import 'package:e_commerce_app/features/chat/data/models/chat_model.dart';

abstract class RemoteContrats {
  Future<List<ChatModel>> getChats();
  Future<ChatModel> getChat(String id);
  Future<ChatModel> createChat(ChatModel chat);
  Future<ChatModel> deleteChat(String id);
}