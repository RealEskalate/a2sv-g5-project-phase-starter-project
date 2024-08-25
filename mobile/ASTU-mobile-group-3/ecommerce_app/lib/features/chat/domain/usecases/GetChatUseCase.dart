import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class GetChatsUseCase {
  final ChatRepository repository;

  GetChatsUseCase(this.repository);

  Future<List<ChatEntity>> call() async {
    return await repository.getChats();
  }
}
