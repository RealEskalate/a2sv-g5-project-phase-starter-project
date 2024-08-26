import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:equatable/equatable.dart';

abstract class ChatState extends Equatable {
  const ChatState();

  @override
  List<Object?> get props => [];
}
class ChatInitateLoading extends ChatState{}
class ChatInitateLoaded extends ChatState{
  final ChatEntity chat;
  const ChatInitateLoaded(this.chat);
  @override
  List<Object?> get props => [chat];


}
class ChatInitateFailure extends ChatState{
  final String error;

  const ChatInitateFailure(this.error);

  @override
  List<Object?> get props => [error];

}

class ChatInitial extends ChatState {}

class ChatLoading extends ChatState {}

class ChatLoaded extends ChatState {
  final List<ChatEntity> messages;

  const ChatLoaded(this.messages);

  @override
  List<Object?> get props => [messages];
}

class ChatError extends ChatState {
  final String error;

  const ChatError(this.error);

  @override
  List<Object?> get props => [error];
}

// import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
// import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
// import 'package:equatable/equatable.dart';

// sealed class ChatsState extends Equatable {
//   final List<ChatEntity> chats;

//   const ChatsState(this.chats);

//   @override
//   List<Object> get props => [];
// }

// class ChatsInitial extends ChatsState {
//   const ChatsInitial(super.chats);
// }

// class ChatsLoadSuccess extends ChatsState {
//   const ChatsLoadSuccess(super.chats);
// }

// class ChatsLoadInProgress extends ChatsState {
//   const ChatsLoadInProgress(super.chats);
// }

// class ChatsInitiateInProgress extends ChatsState {
//   const ChatsInitiateInProgress(super.chats);
// }

// class ChatsInitiateSuccess extends ChatsState {
//   final ChatEntity addedChat;

//   const ChatsInitiateSuccess(this.addedChat, super.chats);
// }

// class ChatsDeleteInProgress extends ChatsState {
//   const ChatsDeleteInProgress(super.chats);
// }

// class ChatsDeleteSuccess extends ChatsState {
//   final String deletedChatName;

//   const ChatsDeleteSuccess(this.deletedChatName, super.chats);
// }

// class ChatsFailure extends ChatsState {
//   final String message;

//   const ChatsFailure(this.message, super.chats);
// }