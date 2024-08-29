part of 'chat_list_bloc.dart';

sealed class ChatListEvent extends Equatable {
  const ChatListEvent();

  @override
  List<Object?> get props => [];
  
}

class LoadAllChatEvent extends ChatListEvent {
  const LoadAllChatEvent();

  @override
  List<Object?> get props => [];
}

class GetChatEvent extends ChatListEvent {
  final String chatId;
  const GetChatEvent({required this.chatId});

  @override
  List<Object?> get props => [chatId];
}

class InitiateChatEvent extends ChatListEvent {
  final String sellerId;
  const InitiateChatEvent({required this.sellerId});

  @override
  List<Object?> get props => [sellerId];
}

class DeleteChatEvent extends ChatListEvent {
  final String chatId;
  const DeleteChatEvent({required this.chatId});

  @override
  List<Object?> get props => [chatId];
}
