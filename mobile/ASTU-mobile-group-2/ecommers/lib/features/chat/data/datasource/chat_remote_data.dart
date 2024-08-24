// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:async';


import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/Error/failure.dart';
import '../../../../core/const/const.dart';
import '../../domain/entity/chat_entity.dart';

abstract interface class ChatRemoteData {
  Future<List<ChatEntity>> getMychats();
  Future<ChatEntity> getChatById(String chatId);
  Future<bool> deleteChats(String id);
  Future<bool> initiate(String id);
}

class ChatRemoteDataImpl implements ChatRemoteData {
  final http.Client client;
  SharedPreferences sharedPreferences;

  ChatRemoteDataImpl({
    required this.client,
    required this.sharedPreferences,
  });

  @override
  Future<bool> deleteChats(String id) {
    try {
      client.delete(
        Uri.parse(ChatApi.deleteChatApi(id)),
        headers: {
          'Authorization': 'Bearer ${sharedPreferences.getString('key')}',
        },
      );
      return Future(true as FutureOr<bool> Function());
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }

  @override
  Future<ChatEntity> getChatById(String chatId) {
    throw UnimplementedError();
  }

  @override
  Future<List<ChatEntity>> getMychats() {
    throw UnimplementedError();
  }

  @override
  Future<bool> initiate(String id) {
    throw UnimplementedError();
  }
}
