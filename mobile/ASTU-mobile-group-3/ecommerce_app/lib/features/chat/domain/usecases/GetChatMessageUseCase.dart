import '../entity/message.dart';
import '../repository/chat_repository.dart';

class GetChatMessagesUseCase {
  final ChatRepository repository;

  GetChatMessagesUseCase(this.repository);

  Future<List<MessageEntity>> call(String chatId) async {
    return await repository.getChatMessages(chatId);
  }
}
