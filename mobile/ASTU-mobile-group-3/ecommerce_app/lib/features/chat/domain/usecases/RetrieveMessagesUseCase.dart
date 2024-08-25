import '../entity/message.dart';
import '../repository/chat_repository.dart';

class RetrieveMessagesUseCase {
  final ChatRepository repository;

  RetrieveMessagesUseCase(this.repository);

  Future<List<MessageEntity>> call(String chatId) async {
    return await repository.retrieveMessages(chatId);
  }
}
