part of 'chat_bloc.dart';

sealed class ChatEvent extends Equatable {
  const ChatEvent();

  @override
  List<Object> get props => [];
}

final class GetAllChatEvent extends ChatEvent{}

final class GetChatMessageEvent extends ChatEvent{
  final  ChatEntity chatEntity;

  const GetChatMessageEvent({required this.chatEntity});
}

final class InitiateChatEvent extends ChatEvent{
  final String recieverId;

  const InitiateChatEvent({required this.recieverId});
}

final class DeleteChatEvent extends ChatEvent{
  final String chatId;

  const DeleteChatEvent({required this.chatId});
}
final class SendMessageEvent extends ChatEvent{
  final String message;
  final String type;
  final ChatEntity chatEntity;

  const SendMessageEvent({required this.message, required this.type, required this.chatEntity});
}

