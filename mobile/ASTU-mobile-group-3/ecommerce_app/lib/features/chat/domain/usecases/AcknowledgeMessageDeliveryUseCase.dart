import '../repository/chat_repository.dart';

class AcknowledgeMessageDeliveryUseCase {
  final ChatRepository repository;

  AcknowledgeMessageDeliveryUseCase(this.repository);

  Future<void> call(String messageId) async {
    await repository.acknowledgeMessageDelivery(messageId);
  }
}
