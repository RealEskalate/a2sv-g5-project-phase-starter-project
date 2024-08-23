import 'dart:convert';

import 'package:ecommerce/core/constants/constants.dart';
import 'package:ecommerce/core/error/exception.dart';
import 'package:http/http.dart' as http;

import '../../../../../core/error/failure.dart';
import '../../presentation/pages/data.dart';
import '../model/chat_model.dart';
import '../model/message_model.dart';
import 'package:dartz/dartz.dart';

import 'remote_abstract.dart';

class RemoteChatSource extends RemoteAbstract {
  final http.Client client;
  RemoteChatSource({required this.client});
  @override
  Future<Either<Failure, String>> chatRoom(
      String token, String receiverId) async {
    try {
      final chatRoomUrl = Uri.parse(Urls.getChatRoom);
       final response = await client.post(
        chatRoomUrl,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token'
        },
        body: jsonEncode({
          'userId': receiverId,
        }),
      );
      if (response.statusCode == 201) {
        final jsondata = jsonDecode(response.body);
        return Right(jsondata['data']['_id']);
      } else {
        return Left(ServerFailure('registration failed'));
      }
    } catch (e) {
      return Left(ServerFailure('registration failed'));
    }
  }

  @override
  Future<Either<Failure, void>> deleteMessage(String chatId) {
    // TODO: implement deleteMessage
    throw UnimplementedError();
  }

  @override
  Stream<Either<Failure, List<MessageModel>>> getMessages(String chatId) {
    // TODO: implement getMessages
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, void>> sendMessage(MessageModel message) {
    // TODO: implement sendMessage
    throw UnimplementedError();
  }

  @override
  Stream<Either<Failure, List<ChatModel>>> getChatHistory(String token) {
    // TODO: implement getChatHistory
    throw UnimplementedError();
  }
}
