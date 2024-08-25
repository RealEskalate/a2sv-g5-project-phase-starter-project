import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/Error/failure.dart';
import '../entity/chat_entity.dart';
import '../repository/chat_repo.dart';

class ChatUsecase extends Equatable {
  final ChatRepositories repositories;
  ChatUsecase({required this.repositories});

  // feach all data

  Future<Either<Failure, List<ChatEntity>>> getMychats() {
    return repositories.getMyChat();
  }

  // get data by id
  Future<Either<Failure, ChatEntity>> getChatById(String chatId) {
    return repositories.getChatById(chatId);
  }

  // edit the product
  Future<Either<Failure, bool>> deleteChats(String id) {
    return repositories.deleteMessages(id);
  }

  @override
  List<Object?> get props => throw UnimplementedError();
}
