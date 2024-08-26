import '../repository/chat_repository.dart';

class CreateChatRoomUseCase {
  final ChatRepository repository;

  CreateChatRoomUseCase(this.repository);

  Future<void> call(String userId) async {
    await repository.createChatRoom(userId);
  }
}
