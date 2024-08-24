import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class GetAllChatsUseCase {
  late ChatRepository chatRepository;
  GetAllChatsUseCase(this.chatRepository);

  Future<Either<Failure, List<Chat>>> execute() async {
    final chats = await chatRepository.getMyChats();
    return Right(chats);
  }
}
