// ignore_for_file: constant_identifier_names

import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/exception/exception.dart';
import '../models/chat_model.dart';

abstract class ChatLocalDataSource {
  Future<Unit> cacheChats(List<ChatModel> chats);
  Future<List<ChatModel>> getChats();
}

const CHATS = 'CHATS';
const MESSAGES = 'MESSAGES';

class ChatLocalDataSourceImpl implements ChatLocalDataSource {
  final SharedPreferences sharedPreferences;

  ChatLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<Unit> cacheChats(List<ChatModel> chats) async {
    try {
      // convert chats to json one by one
      final jsonChats = chats.map((chat) => chat.toJson()).toList();
      // encode the json chats
      final encodedChats = json.encode(jsonChats);

      sharedPreferences.setString(CHATS, encodedChats);
      return Future.value(unit);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<List<ChatModel>> getChats() async {
    try {
      final chats = sharedPreferences.getString(CHATS);
      if (chats != null) {
        final decodedJson = json.decode(chats);
        return ChatModel.fromJsonList(decodedJson);
      } else {
        throw CacheException();
      }
    } catch (e) {
      throw CacheException();
    }
  }
}
