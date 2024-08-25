part of 'chat_bloc.dart';

sealed class ChatState extends Equatable {
  const ChatState();
  
  @override
  List<Object> get props => [];
}

final class ChatInitial extends ChatState {}


final class LoadedAllChatState extends ChatState{
  final List<ChatModel> allChats;

  const LoadedAllChatState({required this.allChats});
}

final class MessagingState extends ChatState{
  final ChatModel chatModel;
  final Stream<MessageModel> chatMessages;
  const MessagingState(this.chatMessages, {required this.chatModel});
}

final class ChatFailureState extends ChatState{
  final String message;

  const ChatFailureState({required this.message});
}

final class ChatLoadingState extends ChatState{}