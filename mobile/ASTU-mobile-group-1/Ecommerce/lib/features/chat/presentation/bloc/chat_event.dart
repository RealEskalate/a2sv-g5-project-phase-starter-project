part of 'chat_bloc.dart';

sealed class ChatEvent extends Equatable {
  const ChatEvent();

  @override
  List<Object> get props => [];
}


final class GetAllChatEvent extends ChatEvent{}

final class GetGetChatMessageEvent extends ChatEvent{
  final  ChatModel chatModel;

  const GetGetChatMessageEvent({required this.chatModel});
}

