import 'dart:convert';

import 'package:equatable/equatable.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../../core/error/exceptions.dart';
import '../../../../product/data/model/user_model.dart';
import '../../../domain/entities/chat.dart';

class MessageModel extends Equatable {
  final String id;
  final UserModel sender;
  final Chat chat;
  final String content;
  final String type;

  const MessageModel(this.id, this.sender, this.chat, this.content, this.type);

  factory MessageModel.fromjson(json) {
    return MessageModel(json['_id'], UserModel.fromJson(json['sender']),
        ChatModel.fromJson(json['chat']), json['content'], json['type']);
  }

  @override
  List<Object?> get props => [id, sender, chat, content, type];
}

class ChatModel extends Chat {
  const ChatModel(
      {required super.id, required super.sender, required super.receiver});

  factory ChatModel.fromJson(json) {
    return ChatModel(
        id: json['_id'], sender: json['sender'], receiver: json['receiver']);
  }
}

class RemoteChatDataSourceImpl {
  final SharedPreferences sharedPreferences;

  final baseUrl = 'https://g5-flutter-learning-path-be.onrender.com/api/v3';
  final tTokenKey = 'tokenKey';

  RemoteChatDataSourceImpl(this.sharedPreferences);
  Future<List<MessageModel>> getChatMessages(String id) async {
    try {
      var url = Uri.parse('$baseUrl/chat/$id/messages');

      final tToken = sharedPreferences.getString(tTokenKey);
      if (tToken == null) {
        throw Exception('token not found');
      }
      final result = await http.get(url, headers: {
        'Authorization': 'Bearer $tToken',
        'Content-Type': 'application/json'
      });

      if (result.statusCode == 200) {
        final listMessages = jsonDecode(result.body)['data'];
        return listMessages.map((json) => MessageModel.fromjson(json));
      } else {
        throw ServerException(message: 'server failure');
      }
    } catch (e) {
      throw ServerException(message: e.toString());
    }
  }
}
