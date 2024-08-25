import 'dart:convert';
import 'dart:io';

import 'package:dartz/dartz.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/exception/exception.dart';
import '../../../../core/network/custom_client.dart';
import '../models/chat_model.dart';
import '../models/message_model.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatModel>> myChats();
  Future<ChatModel> myChatById(String chatId);
  Future<List<MessageModel>> getChatMessages(String chatId);
  Future<ChatModel> initiateChat(String sellerId);
  Future<Unit> deleteChat(String chatId);
}

//  implement the class
class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final CustomHttpClient client;
  ChatRemoteDataSourceImpl({required this.client});
  @override
  Future<List<ChatModel>> myChats() async {
    try {
      final response = await client.get(Urls.myChats());
      if (response.statusCode == 200) {
        return ChatModel.fromJsonList(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<ChatModel> myChatById(String chatId) async {
    try {
      final response = await client.get(Urls.myChatById(chatId));
      if (response.statusCode == 200) {
        return ChatModel.fromJson(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<List<MessageModel>> getChatMessages(String chatId) async {
    try {
      final response = await client.get(Urls.getChatMessages(chatId));
      if (response.statusCode == 200) {
        return MessageModel.fromJsonList(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<ChatModel> initiateChat(String sellerId) async {
    try {
      final response = await client.post(Urls.myChats(), body: {
        'userId': sellerId,
      });
      if (response.statusCode == 200) {
        return ChatModel.fromJson(json.decode(response.body)['data']);
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<Unit> deleteChat(String chatId) async {
    try {
      final response = await client.delete(Urls.myChatById(chatId));
      if (response.statusCode == 200) {
        return unit;
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }
}
