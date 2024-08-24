import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class InitiateChatUseCase extends UseCase<ChatEntity, InitiateChatParams> {
  final ChatRepository chatRepository;

  InitiateChatUseCase(this.chatRepository);

  @override
  Future<Either<Failure, ChatEntity>> call(InitiateChatParams params) async {
    return await chatRepository.initiateChat(params.sellerId);
  }
}

class InitiateChatParams extends Equatable {
  final String sellerId;

  const InitiateChatParams({required this.sellerId});

  @override
  List<Object> get props => [sellerId];
}
