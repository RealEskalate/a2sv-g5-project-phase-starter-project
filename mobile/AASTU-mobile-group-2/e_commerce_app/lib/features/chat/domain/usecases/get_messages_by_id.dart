import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../data/data repository/data_repository.dart';

class GetMessagesByIdUsecase {
  final ChatRepositoryImpl _chatRepository;

  const GetMessagesByIdUsecase(this._chatRepository);

  Future<Either<Failure, List<MessageEntity>>> execute(String chatId) async {
    return await _chatRepository.getMessagesById(chatId);
  }
}

class GetChatMessagesParams extends Equatable {
  final String id;

  const GetChatMessagesParams(this.id);

  @override
  List<Object?> get props => [id];
}
