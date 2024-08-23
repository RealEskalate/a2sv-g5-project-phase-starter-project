import 'package:equatable/equatable.dart';

import '../../domain/entity/chat_entity.dart';

abstract class ChatState extends Equatable {
  const ChatState();

  @override
  List<Object?> get props => [];
}

class InitialChatState extends ChatState {}

class LoadingChatState extends ChatState {}

class ErrorChatState extends ChatState {
  final String errorMessage;

  const ErrorChatState({required this.errorMessage});

  @override
  List<Object?> get props => [errorMessage];
}

class LoadedChatState extends ChatState {
  final List<ChatEntity> chats; 

  const LoadedChatState({required this.chats});

  @override
  List<Object?> get props => [chats];
}

class SingleChatLoadedState extends ChatState {
  final ChatEntity chat; 

  const SingleChatLoadedState({required this.chat});

  @override
  List<Object?> get props => [chat];
}

class InitiatedChatState extends ChatState {
  final ChatEntity chat;

  const InitiatedChatState({required this.chat});

  @override
  List<Object?> get props => [chat];
}

class SuccessChatState extends ChatState {
  final String message;

  const SuccessChatState({required this.message});

  @override
  List<Object?> get props => [message];
}
