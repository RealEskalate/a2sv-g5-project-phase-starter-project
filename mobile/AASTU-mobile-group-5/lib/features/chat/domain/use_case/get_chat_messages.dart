// import 'package:dartz/dartz.dart';

// import '../../../../core/failure/failure.dart';
// import '../entities/message_entity.dart';
// import '../repositories/message_repository.dart';

// class GetChatMessages{
//   final MessageRepository repository;

//   GetChatMessages(this.repository);

//   Future<Either<Failure, List<MessageEntity>>> call(String messageId) async {
//     return await repository.getChatMessages(messageId);
//   }
// }

import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

import '../../../../core/failure/failure.dart';
import '../../data/datasources/websocket_service.dart';
import '../entities/message_entity.dart';
import '../repositories/message_repository.dart';

class GetChatMessages {
  final MessageRepository repository;
  final WebSocketService webSocketService;

  GetChatMessages(this.repository, this.webSocketService);

  Future<Either<Failure, List<MessageEntity>>> call(String chatId) async {
    try {
      // Listen for incoming messages related to the chat
      webSocketService.channel.sink.add(jsonEncode({'action': 'getMessages', 'chatId': chatId}));
      final messages = await repository.getChatMessages(chatId);

      // Handle the messages received through WebSocket
      webSocketService.messages.listen((message) {
        // Parse and handle incoming messages here
        // For example, update the chat UI with new messages
      });

      return messages;
    } catch (e) {
      return Left(ServerFailure(message: 'Failed to get chat messages'));
    }
  }
}
