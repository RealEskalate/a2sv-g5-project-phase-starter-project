import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/chat.dart';
import '../repository/chat_repository.dart';

abstract class IntiateChatUseCase {
  late ChatRepository chatRepository;
  IntiateChatUseCase(this.chatRepository);

  Future<Either<Failure,Chat>> execute(String userId) async{
    final chat = await chatRepository.intiateChat(userId);
    return Right(chat);
  }
}