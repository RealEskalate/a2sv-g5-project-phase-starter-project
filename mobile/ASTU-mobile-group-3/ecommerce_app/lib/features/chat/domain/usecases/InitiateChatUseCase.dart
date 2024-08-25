import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class InitiateChatUseCase {
  final ChatRepository repository;

  InitiateChatUseCase(this.repository);

  Future<ChatEntity> call(String userId) async {
    return await repository.initiateChat(userId);
  }
}
