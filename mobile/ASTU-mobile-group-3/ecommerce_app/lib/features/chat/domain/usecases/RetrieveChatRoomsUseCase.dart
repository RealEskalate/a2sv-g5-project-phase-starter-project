import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class RetrieveChatRoomsUseCase {
  final ChatRepository repository;

  RetrieveChatRoomsUseCase(this.repository);

  Future<List<ChatEntity>> call() async {
    return await repository.retrieveChatRooms();
  }
}
