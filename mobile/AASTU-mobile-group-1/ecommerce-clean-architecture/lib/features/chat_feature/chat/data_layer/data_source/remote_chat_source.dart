import 'dart:convert';
import 'dart:io';

import 'package:ecommerce/core/constants/constants.dart';
import 'package:ecommerce/core/error/exception.dart';
import 'package:http/http.dart' as http;
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

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
  Stream<Either<Failure, List<ChatModel>>> getChatHistory(String token) async* {
    try {
          final url = Uri.parse(Urls.getChatHistory);
      final headers = {
        'Authorization': 'Bearer $token',
      };

      final response = await client.get(url, headers: headers);
     
      if (response.statusCode == 200) {
      
        try {
          
          final chatList = parseChatModelList(response.body);
          yield Right(chatList);
        } catch (e) {
         
          yield Left(ServerFailure('Error parsing data: $e'));
        }
      } else {
       
        yield Left(
            ServerFailure('Server returned an error: ${response.statusCode}'));
      }
    } catch (e) {
    
      yield Left(ServerFailure('An error has occured'));
    }
  }
}

List<ChatModel> parseChatModelList(String message) {
  
  final Map<String, dynamic> jsonMap = jsonDecode(message);

   
    final dataList = jsonMap['data'] as List<dynamic>;

    return dataList.map<ChatModel>((json) => ChatModel.fromJson(json)).toList();

}
