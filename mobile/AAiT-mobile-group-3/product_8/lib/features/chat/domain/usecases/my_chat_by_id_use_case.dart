import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class MyChatByIdUseCase extends UseCase<ChatEntity, MyChatByIdParams> {
  final ChatRepository chatRepository;

  MyChatByIdUseCase(this.chatRepository);

  @override
  Future<Either<Failure, ChatEntity>> call(MyChatByIdParams params) async {
    return await chatRepository.myChatById(params.chatId);
  }
}

class MyChatByIdParams extends Equatable {
  final String chatId;

  const MyChatByIdParams({required this.chatId});

  @override
  List<Object> get props => [chatId];
}
