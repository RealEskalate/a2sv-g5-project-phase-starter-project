import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/exception.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/repository/chat_repository.dart';

import '../../../../auth/domain/entities/user.dart';
import '../entity/message.dart';

class SendMessage {
  final ChatRepository repository;

  SendMessage({required this.repository});

  void call(
      String chatId, String content, String type)  {
    try {
       repository.sendMessage(chatId, content, type);
     
    } catch (e) {
      throw ServerException("Server Error");
    }
  }
}
