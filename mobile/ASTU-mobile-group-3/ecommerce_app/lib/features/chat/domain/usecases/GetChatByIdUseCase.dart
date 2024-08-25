import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class GetChatByIdUseCase {
  final ChatRepository repository;

  GetChatByIdUseCase(this.repository);

  Future<ChatEntity> call(String chatId) async {
    return await repository.chatById(chatId);
  }
}
