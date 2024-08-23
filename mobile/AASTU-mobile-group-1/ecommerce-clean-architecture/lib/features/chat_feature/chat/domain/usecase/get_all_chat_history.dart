import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/repository/chat_repository.dart';

import '../../../../auth/domain/entities/user.dart';
import '../entity/chat.dart';
import '../entity/message.dart';

class GetAllChatHistory {
  final ChatRepository repository;

  GetAllChatHistory({required this.repository});

  Stream<Either<Failure, List<ChatEntity>>> call() async* {
    yield* repository.getChatHistory();
  }
}