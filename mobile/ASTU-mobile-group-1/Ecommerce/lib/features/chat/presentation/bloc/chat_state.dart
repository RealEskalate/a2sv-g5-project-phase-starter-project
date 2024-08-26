part of 'chat_bloc.dart';

sealed class ChatState extends Equatable {
  const ChatState();
  
  @override
  List<Object> get props => [];
}

final class ChatInitial extends ChatState {}


final class LoadedAllChatState extends ChatState{
  final List<ChatEntity> allChats;
  final List<UserEntity> users;
  const LoadedAllChatState(this.users, {required this.allChats});
}

final class IndividualChatingState extends ChatState{
  final List<MessageEntity> chatMessages;
  final ChatEntity chatEntity;
  const IndividualChatingState({required this.chatEntity, required this.chatMessages});
}

final class ChatFailureState extends ChatState{
  final String message;

  const ChatFailureState({required this.message});
}

final class ChatLoadingState extends ChatState{}