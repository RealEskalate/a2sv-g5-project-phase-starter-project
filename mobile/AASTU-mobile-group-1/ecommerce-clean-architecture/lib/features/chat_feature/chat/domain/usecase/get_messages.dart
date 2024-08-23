import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/entity/message.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/repository/chat_repository.dart';

class GetMessages {
  final ChatRepository repository;

  GetMessages({required this.repository});

  Stream<Either<Failure,List<Message>>> call(String chatId) async* {
    yield* repository.getMessages(chatId);
  }
}