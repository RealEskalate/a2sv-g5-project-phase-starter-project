import 'dart:convert';
import 'package:ecommerce/core/constants/constants.dart';
import 'package:ecommerce/core/error/exception.dart';
import 'package:ecommerce/features/chat_feature/chat/data_layer/data_source/Service/socker_service.dart';
import 'package:http/http.dart' as http;
import 'package:dartz/dartz.dart';

import '../../../../../core/error/failure.dart';
import '../model/chat_model.dart';
import '../model/message_model.dart';
import 'remote_abstract.dart';

class RemoteChatSource extends RemoteAbstract {
  final http.Client client;
  late final SocketService socketService;

  RemoteChatSource({required this.client});

  // Initialize SocketService when the user authenticates or provides a token
  void initializeSocket(String token) {
    socketService = SocketService(token);
  }

  @override
  Future<Either<Failure, String>> chatRoom(
      String token, String receiverId) async {
    try {
      final chatRoomUrl = Uri.parse(Urls.getChatRoom);
      final response = await client.post(
        chatRoomUrl,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({
          'userId': receiverId,
        }),
      );
      if (response.statusCode == 201) {
        final jsondata = jsonDecode(response.body);
        return Right(jsondata['data']['_id']);
      } else {
        return Left(ServerFailure('Failed to create chat room'));
      }
    } catch (e) {
      return Left(ServerFailure('An error occurred: ${e.toString()}'));
    }
  }

 @override
  Future<Either<Failure, void>> deleteMessage(String chatId,String token) async {
    try {
      final response = await client.delete(
        Uri.parse("${Urls.deleteChat}/$chatId"),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json', 
        },
      );

      if (response.statusCode == 201) {
        print("delete ");
        return Right(null);
      } else {
        return Left(ServerFailure('Failed to delete message'));
      }
    } catch (e) {
      return Left(ServerFailure('Server error occurred'));
    }
  }

  @override
  Stream<MessageModel> getMessages(String chatId, String token) async* {
    // Ensure the SocketService is initialized
   
    // Fetch existing messages via HTTP GET
    try {
      final headers = {
        'Authorization': 'Bearer $token',
      };

      final response = await client.get(
        Uri.parse('${Urls.getChatHistory}/$chatId/messages'),
        headers: headers,
      );

      if (response.statusCode == 200) {
        final List<dynamic> messages = jsonDecode(response.body)['data'];
        for (var message in messages) {
          // print("message: $message");
          yield MessageModel.fromJson(message);
        }
      } else {
        throw ServerException(
            'Server returned an error: ${response.statusCode}');
      }
    } catch (e) {
      
      throw ServerException('An error occurred while fetching messages: ${e.toString()}');
    }

   
    yield* socketService.messages.where((msg) => msg.chat.chatId == chatId);
  }

  @override
  void sendMessage(String chatId, String message, String type) {
    // Use SocketService to send messages
    socketService.sendMessage(chatId, message, type);
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
      yield Left(ServerFailure('An error has occurred: ${e.toString()}'));
    }
  }

  List<ChatModel> parseChatModelList(String message) {
    final Map<String, dynamic> jsonMap = jsonDecode(message);
    final dataList = jsonMap['data'] as List<dynamic>;
    return dataList.map<ChatModel>((json) => ChatModel.fromJson(json)).toList();
  }
}
