import '../entity/message.dart';
import '../repository/chat_repository.dart';

class SendMessageUseCase {
  final ChatRepository repository;

  SendMessageUseCase(this.repository);

  Future<void> call(String chatId, MessageEntity message) async {
    await repository.sendMessage(chatId, message);
  }
}
