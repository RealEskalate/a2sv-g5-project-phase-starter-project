part of 'chat_list_bloc.dart';

sealed class ChatListState extends Equatable {
  const ChatListState();

  @override
  List<Object?> get props => [];
}

final class ChatListInitial extends ChatListState {}

class LoadAllChatState extends ChatListState {

  final List<ChatEntity> chats;
  const LoadAllChatState({required this.chats});

  @override
  List<Object?> get props => [chats];
}

class LoadingAllChatState extends ChatListState {
  const LoadingAllChatState();

  @override
  List<Object?> get props => [];
}

class LoadingAllChatError extends ChatListState {
  final String errorMessage;
  const LoadingAllChatError({required this.errorMessage});

  @override
  List<Object?> get props => [errorMessage];
}





class LoadedSingleChatState extends ChatListState {
  final ChatEntity chat;
  const LoadedSingleChatState({required this.chat});

  @override
  List<Object?> get props => [chat];
}

class LoadingSingleChat extends ChatListState {
  const LoadingSingleChat();

  @override
  List<Object?> get props => [];
}

class LoadingSingleChatError extends ChatListState {
  final String errorMessage;
  const LoadingSingleChatError({required this.errorMessage});

  @override
  List<Object?> get props => [errorMessage];
}





class InitiatedChatState extends ChatListState {
  final ChatEntity chat;

  const InitiatedChatState({required this.chat});
}

class InitiatingChatState extends ChatListState {
  const InitiatingChatState();

  @override
  List<Object?> get props => [];
}

class InitiatingChatError extends ChatListState {

  final String errorMessage;
  const InitiatingChatError({required this.errorMessage});

  @override
  List<Object?> get props => [errorMessage];
}





class DeletedChatState extends ChatListState {
  const DeletedChatState();
}

class DeletingChatState extends ChatListState {
  const DeletingChatState();

  @override
  List<Object?> get props => [];
}

class DeletingChatError extends ChatListState {

  final String errorMessage;
  const DeletingChatError({required this.errorMessage});

  @override
  List<Object?> get props => [errorMessage];
}
