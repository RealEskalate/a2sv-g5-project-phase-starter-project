import '../repository/chat_repository.dart';

class SendMessageUseCase {
  final ChatRepository repository;

  SendMessageUseCase(this.repository);

  Future<void> call(String chatId, String content, String type) async {
    await repository.sendMessage(chatId, content, type);
  }
}
