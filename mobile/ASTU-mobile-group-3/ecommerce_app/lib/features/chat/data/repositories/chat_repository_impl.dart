import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';
import '../../domain/repository/chat_repository.dart';
import '../data_resources/remote_chat_data_source.dart';
import '../data_resources/socket_io_sesrvice.dart';

class ChatRepositoryImpl implements ChatRepository {
  final ChatRemoteDataSource chatRemoteDataSource;
  final SocketIOService socketIOService;

  ChatRepositoryImpl({
    required this.chatRemoteDataSource,
    required this.socketIOService,
  });

  @override
  Future<List<ChatEntity>> retrieveChatRooms() async {
    return await chatRemoteDataSource.getChatRooms();
  }

  @override
  Future<List<MessageEntity>> retrieveMessages(String chatId) async {
    return await chatRemoteDataSource.getMessages(chatId);
  }

  @override
  Future<void> createChatRoom(String userId) async {
    await chatRemoteDataSource.createChatRoom(userId);
  }

  @override
  Future<void> acknowledgeMessageDelivery(String messageId) {
    // TODO: implement acknowledgeMessageDelivery
    throw UnimplementedError();
  }

  @override
  Stream<MessageEntity> onMessageReceived() {
    // TODO: implement onMessageReceived
    throw UnimplementedError();
  }

  @override
  Future<void> sendMessage(String chatId, String content, String type) async {
    // TODO: implement sendMessage
    // throw UnimplementedError();
    await socketIOService.emitSendMessage(chatId, content, type);
  }

  // @override
  // Future<void> sendMessage(String chatId, String content, String type) async {
  // await socketIOService.emitSendMessage(chatId, content, type);
  // }

  // @override
  // Future<void> acknowledgeMessageDelivery(String messageId) async {
  //   await socketIOService.emitMessageDelivered(messageId);
  // }

  // @override
  // Stream<MessageEntity> onMessageReceived() {
  //   // return socketIOService.onMessageReceived();
  // }
}
