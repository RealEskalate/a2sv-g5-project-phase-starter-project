class ChatEvent {}


class ChatMessageSentEvent extends ChatEvent {
  final String message;
  ChatMessageSentEvent(this.message);
}

class DeleteChatEvent extends ChatEvent {
  final String chatId;
  DeleteChatEvent(this.chatId);
}

class GetAllChatEvent extends ChatEvent {}

class GetChatByIdEvent extends ChatEvent{
  final String chatId;
  GetChatByIdEvent(this.chatId);
}