part of 'chat_bloc.dart';

@immutable
sealed class ChatEvent {}

// final class LoadSignedUpUsers extends ChatEvent{}

final class LoadCurrentChats extends ChatEvent {}
final class LoadChatWithUser extends ChatEvent {
  final String userId;
  LoadChatWithUser(this.userId);
}
final class CreateChat extends ChatEvent {
  final String sellerId;
  CreateChat(this.sellerId);
}

// final class SendMessage extends ChatEvent {
//   // final Message message;
//   // SendMessage(this.message);
// }

// final class Typing extends ChatEvent {
//   final bool isTyping;
//   Typing(this.isTyping);
// }

final class DeleteChat extends ChatEvent {
  final String chatId;
  DeleteChat(this.chatId);
}

// final class ViewUserProfile extends ChatEvent {
//   final String userId;
//   ViewUserProfile(this.userId);
// }