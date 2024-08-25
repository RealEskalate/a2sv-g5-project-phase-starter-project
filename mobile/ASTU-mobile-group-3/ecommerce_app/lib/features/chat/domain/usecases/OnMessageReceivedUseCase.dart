import '../entity/message.dart';
import '../repository/chat_repository.dart';

class OnMessageReceivedUseCase {
  final ChatRepository repository;

  OnMessageReceivedUseCase(this.repository);

  Stream<MessageEntity> execute() {
    return repository.onMessageReceived();
  }
}
