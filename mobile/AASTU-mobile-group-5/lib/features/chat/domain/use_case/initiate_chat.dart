import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../data/models/chat_model.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class InitiateChat {
  final ChatRepository chatRepository;

  InitiateChat(this.chatRepository);
  
  Future<Either<Failure, ChatModel>> execute(String sellerId) {
    return chatRepository.initiateChat(sellerId);
  }
}

// import 'dart:convert';

// import 'package:dartz/dartz.dart';
// import 'package:web_socket_channel/web_socket_channel.dart';

// import '../../../../core/failure/failure.dart';
// import '../../data/datasources/websocket_service.dart';
// import '../entities/chat_entity.dart';
// import '../repositories/chat_repository.dart';

// class InitiateChat {
//   final ChatRepository chatRepository;
//   final WebSocketService webSocketService;

//   InitiateChat(this.chatRepository, this.webSocketService);

//   Future<Either<Failure, ChatEntity>> execute(String sellerId) async {
//     try {
//       // Notify server to initiate chat
//       webSocketService.sendMessage(jsonEncode({'action': 'initiateChat', 'sellerId': sellerId}), {});
//       final chat = await chatRepository.initiateChat(sellerId);

//       // Listen for server response about chat initiation
//       webSocketService.messages.listen((message) {
//         // Handle server response here
//       });

//       return chat;
//     } catch (e) {
//       return Left(ServerFailure(message: 'Failed to initiate chat'));
//     }
//   }
// }
