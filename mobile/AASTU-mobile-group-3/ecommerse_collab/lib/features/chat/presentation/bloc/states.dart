import '../../domain/entity/chat.dart';

class ChatState {}

class LoadingAllChatsState extends ChatState {}
class AllChatsLoadedState extends ChatState {
  final List<Chat> chats;
  AllChatsLoadedState(this.chats);
}

class LoadingChatState extends ChatState {}
class ChatLoadedState extends ChatState {
  final Chat chat;
  ChatLoadedState(this.chat);
}

class DeleteLoadingChatState extends ChatState {}
class ChatDeletedState extends ChatState {}

class ChatErrorState extends ChatState {
  final String message;
  ChatErrorState(this.message);
}

class ChatInitiateState extends ChatState {}
class ChatSentState extends ChatState {}