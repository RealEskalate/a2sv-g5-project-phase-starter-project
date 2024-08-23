import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/entity/message.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/repository/chat_repository.dart';

class SendMessage {
  final ChatRepository repository;

  SendMessage({required this.repository});

  Future<Either<Failure, String>> call(String recieverId) async {
    return await repository.chatRoom(recieverId);
  }
}
