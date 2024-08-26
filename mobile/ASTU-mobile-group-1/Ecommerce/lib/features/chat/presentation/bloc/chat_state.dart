part of 'chat_bloc.dart';

@immutable
sealed class ChatState {}

final class ChatInitial extends ChatState {}

final class ChatLoaded extends ChatState {
  final List<MessageEntity> messages;

  ChatLoaded({required this.messages});
}

final class ChatErrorState extends ChatState {
  final String message;

  ChatErrorState({required this.message});
}

 
class ChatLoadingState extends ChatState {}
class ChatFailureState extends ChatState{
  final String message;
  ChatFailureState({required this.message});
}


final class LoadedAllChatState extends ChatState{
  final List<ChatEntity> allChats;
  // final List<UserEntity> users;
   LoadedAllChatState({required this.allChats});
}


final class LoadedAllMessages extends ChatState{
  final List<MessageEntity> allMessages;

  LoadedAllMessages({required this.allMessages});
 
}