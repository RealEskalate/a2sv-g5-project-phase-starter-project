import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class CreateChatRoomUseCase {
  final ChatRepository repository;

  CreateChatRoomUseCase(this.repository);

  Future<void> call(ChatEntity chat) async {
    await repository.createChatRoom(chat);
  }
}
