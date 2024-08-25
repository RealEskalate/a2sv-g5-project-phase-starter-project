import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/entity/message.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/repository/chat_repository.dart';

class GetMessages {
  final ChatRepository repository;

  GetMessages({required this.repository});

  Stream<Message> call(String chatId) async* {
    // print("get");
    // yield* repository.getMessages(chatId);
    yield* repository.getMessages(chatId).map((message) {
      // Log each message received
      print('Received message: ${message}');
      return message;
    });
    ;
  }
}
