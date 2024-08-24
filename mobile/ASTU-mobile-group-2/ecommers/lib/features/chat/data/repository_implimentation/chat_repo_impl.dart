


import '../../../../core/Error/failure.dart';
import '../../domain/entity/chat_entity.dart';
import 'package:dartz/dartz.dart';

import '../../domain/repository/chat_repo.dart';

class ChatRepoImpl  implements ChatRepositories{
  @override
  Future<Either<Failure, bool>> deleteMessages(String id) {
    // TODO: implement deleteMessages
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, ChatEntity>> getChatById(String chatId) {
    // TODO: implement getChatById
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<ChatEntity>>> getMyChat() {
    // TODO: implement getMyChat
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, bool>> initiateChat(String userId) {
    // TODO: implement initiateChat
    throw UnimplementedError();
  }


}