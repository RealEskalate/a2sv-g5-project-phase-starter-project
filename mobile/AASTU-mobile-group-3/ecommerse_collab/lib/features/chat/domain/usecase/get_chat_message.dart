import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class GetChatMessageUseCase {
  late ChatRepository chatRepository;
  GetChatMessageUseCase(this.chatRepository);

  Future<Either<Failure,Chat>> execute(String id) async{
    final chat = await chatRepository.getChatMessage(id);
    return Right(chat as Chat);  
  }
}
