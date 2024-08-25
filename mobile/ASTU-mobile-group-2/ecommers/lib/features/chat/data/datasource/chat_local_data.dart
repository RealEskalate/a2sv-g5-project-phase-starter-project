import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../model/chat_model.dart';

abstract class ChatLocalData {
  Future<ChatModel?> getCachedChats();
  Future<void> cacheChats(List<String> chats);
  Future<void> clearCachedChats();
}

class ChatLocalDataImpl implements ChatLocalData {
  final SharedPreferences sharedPreferences;
  // ignore: non_constant_identifier_names
  final String CACHED_CHATS = 'CACHED_CHATS';

  ChatLocalDataImpl({required this.sharedPreferences});

  @override
  Future<ChatModel?> getCachedChats() async {
    final jsonString = sharedPreferences.getString(CACHED_CHATS);
    if (jsonString != null) {
      return ChatModel.fromJson(jsonDecode(jsonString));
    } else {
      return null;
    }
  }

  @override
  Future<void> cacheChats(List<String> chats) async {
    // final userJson = jsonEncode(userToCache.toJson());
    await sharedPreferences.setStringList(CACHED_CHATS, chats);
  }

  @override
  Future<void> clearCachedChats() async {
    await sharedPreferences.remove(CACHED_CHATS);
  }
}
