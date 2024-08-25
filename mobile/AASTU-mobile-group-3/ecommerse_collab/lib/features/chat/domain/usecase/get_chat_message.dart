import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/message.dart';
import '../repository/chat_repository.dart';

class GetChatMessageUseCase {
  late ChatRepository chatRepository;
  GetChatMessageUseCase(this.chatRepository);

  Future<Either<Failure,List<Message>>> execute(String chatId) async{
    final chats = await chatRepository.getChatMessage(chatId);
    return Right(chats);  
  }
}
