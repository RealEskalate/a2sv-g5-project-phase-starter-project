import 'dart:convert';
import 'package:dartz/dartz.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/failure/failure.dart';
import '../models/chat_model.dart';


abstract class ChatLocalDataSource {
  Future<Either<Failure, ChatModel>> getChatById(String id);
  Future<void> cacheChat(ChatModel chatToCache);
}

class ChatLocalDataSourceImpl implements ChatLocalDataSource {
  final SharedPreferences sharedPreferences;

  ChatLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<Either<Failure, ChatModel>> getChatById(String id) async {
    try {
      final jsonString = sharedPreferences.getString(id);
      if (jsonString != null) {
        final chatModel = ChatModel.fromJson(json.decode(jsonString));
        return Right(chatModel);
      } else {
        return Left(CacheFailure(message: 'cache failure'));
      }
    } catch (e) {
      return Left(CacheFailure(message: 'cache failure'));
    }
  }

  @override
  Future<void> cacheChat(ChatModel chatToCache) async {
    final jsonString = json.encode(chatToCache.toJson());
    await sharedPreferences.setString(chatToCache.chat_id, jsonString);
  }
}
